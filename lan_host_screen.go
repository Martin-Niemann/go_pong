package main

import (
	"bytes"
	"fmt"
	"net"
	"net/netip"
	"strconv"
	"strings"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var mutex = 1

const (
	bad_message      = "BAD"
	send_player_info = "PLAYER_INFO"
	ready            = "READY"
	start_countdown  = "COUNTDOWN"
)

type LanHostScreen struct {
	network          Network
	addr_chan        chan net.Addr
	address          net.Addr
	buffer_chan      chan []byte
	top_text         string
	bottom_text      string
	connection_state int
	gameplay         GameplayMultiplayer
	start_time       time.Time
	update_func      func()
	draw_func        func()
}

func (l *LanHostScreen) initialize_server(addrport netip.AddrPort) {
	l.update_func = l.update_pre
	l.draw_func = l.draw_pre
	roboto = rl.LoadFontEx("./RobotoMono-Medium.ttf", 64, []rune("a b c d e f g h i j k l m n o p q r s t u v w x y z A B C D E F G H I J K L M N O P Q R S T U V W X Y Z 1 2 3 4 5 6 7 8 9 0 . \" : !"))

	l.connection_state = no_connection
	l.top_text = fmt.Sprintf("Hosting on %v", addrport.String())
	l.bottom_text = "Waiting for an opponent to join..."
	l.network = Network{}
	// make the first channel buffered, so we can send the second channel response without being locked
	l.addr_chan = make(chan net.Addr, 1)
	l.buffer_chan = make(chan []byte)
	l.network.initialize_host(addrport)
	go l.network.spawn_host_listener(l.addr_chan, l.buffer_chan)
}

func (l *LanHostScreen) update_pre() {
	l.send_start()
	l.await()
}

func (l *LanHostScreen) update_game() {
	l.await()
	l.gameplay.host_update(&l.gameplay.p1)
	l.send_data()
}

func (l *LanHostScreen) await() {
	if mutex == 1 {
		return
	} else if mutex == 0 {
		address := <-l.addr_chan
		buffer := <-l.buffer_chan

		message := l.buffer_to_string(buffer)
		l.act_on_message(message, address)

		// set mutex back to 1 so we can await the next buffer
		mutex = 1

		go l.network.spawn_host_listener(l.addr_chan, l.buffer_chan)
	}
}

func (l *LanHostScreen) draw_pre() {
	rl.BeginDrawing()
	rl.DrawFPS(10, 10)
	rl.ClearBackground(rl.LightGray)

	enter_text := "Press enter to start the game!"
	enter_text_size := rl.MeasureTextEx(roboto, enter_text, 48, 1)
	top_text_size := rl.MeasureTextEx(roboto, l.top_text, 64, 1)
	bottom_text_size := rl.MeasureTextEx(roboto, l.bottom_text, 32, 1)

	rl.DrawTextEx(roboto, l.top_text, rl.NewVector2((screen_width/2)-(top_text_size.X/2), (screen_height/2)-(top_text_size.Y/2)-80), 64, 1, rl.Black)
	rl.DrawTextEx(roboto, l.bottom_text, rl.NewVector2((screen_width/2)-(bottom_text_size.X/2), (screen_height/2)-(bottom_text_size.Y/2)), 32, 1, rl.Black)

	if l.connection_state == sent_player_info_ack {
		rl.DrawTextEx(roboto, enter_text, rl.NewVector2((screen_width/2)-(enter_text_size.X/2), (screen_height/2)-(enter_text_size.Y/2)+80), 48, 1, rl.Black)
	}

	rl.EndDrawing()
}

func (l *LanHostScreen) buffer_to_string(buffer []byte) string {
	buffer_no_zeroes := bytes.Trim(buffer, "\u0000")
	string_no_whitespace := strings.TrimSpace(string(buffer_no_zeroes))
	return string_no_whitespace
}

func (l *LanHostScreen) act_on_message(message string, address net.Addr) {
	var split_message []string

	if l.connection_state == no_connection {
		split_message = strings.SplitN(message, " ", 3)
		if len(split_message) == 3 && split_message[0] == "JOIN" && split_message[1] == "UDPONG" && split_message[2] == "V0.1" {
			go l.network.respond(address, []byte(send_player_info))
			l.connection_state = correct_join_message
		}
	} else if l.connection_state == correct_join_message {
		split_message = strings.SplitN(message, "\"", 3)
		if len(split_message) != 3 {
			go l.network.respond(address, []byte(bad_message))
		} else {
			player_name := split_message[0]
			player_quote := split_message[1]
			go l.network.respond(address, []byte("gamer21849\"I am the pong champ\""))
			l.connection_state = recieved_player_info
			l.bottom_text = fmt.Sprintf("%v \"%v\" has joined", player_name, player_quote)
		}
	} else if l.connection_state == recieved_player_info {
		if message == "ACK" {
			go l.network.respond(address, []byte(ready))
			l.address = address
			l.connection_state = sent_player_info_ack
		} else {
			go l.network.respond(address, []byte(bad_message))
		}
	} else if l.connection_state == requested_countdown {
		paddle_y, error := strconv.ParseFloat(message, 32)
		//print(paddle_y)
		if error != nil {
			go l.network.respond(address, []byte(bad_message))
			return
		}
		l.gameplay.p2.y = float32(paddle_y)
	}
}

func (l *LanHostScreen) send_start() {
	if l.connection_state == sent_player_info_ack {
		if rl.IsKeyDown(rl.KeyEnter) {
			go l.network.respond(l.address, []byte(start_countdown))
			l.connection_state = requested_countdown
			l.initialize_gameplay()
		}
	}
}

func (l *LanHostScreen) send_data() {
	l.network.respond(l.address, []byte(fmt.Sprintf("%g %g %g", l.gameplay.ball.x, l.gameplay.ball.y, l.gameplay.p1.y)))
}

func (l *LanHostScreen) initialize_gameplay() {
	l.gameplay = GameplayMultiplayer{}
	l.gameplay.initialize_gameobjects()
	l.update_func = l.update_game
	l.draw_func = l.gameplay.draw
}
