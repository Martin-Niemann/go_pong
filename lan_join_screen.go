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

const (
	connect_message = "JOIN UDPONG V0.1"
)

type LanJoinScreen struct {
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

func (l *LanJoinScreen) initialize_connection(addrport netip.AddrPort) {
	l.connection_state = no_connection
	l.update_func = l.update_pre
	l.draw_func = l.draw_pre

	roboto = rl.LoadFontEx("./RobotoMono-Medium.ttf", 64, []rune("a b c d e f g h i j k l m n o p q r s t u v w x y z A B C D E F G H I J K L M N O P Q R S T U V W X Y Z 1 2 3 4 5 6 7 8 9 0 . \" :"))
	l.top_text = fmt.Sprintf("Trying to connect to %v", addrport.String())
	l.bottom_text = "Resolving..."

	l.network = Network{}
	// make the first channel buffered, so we can send the second channel response without being locked
	l.addr_chan = make(chan net.Addr, 1)
	l.buffer_chan = make(chan []byte)
	l.network.initialize_client(addrport)
	go l.network.spawn_client_listener(l.addr_chan, l.buffer_chan)
}

func (l *LanJoinScreen) update_pre() {
	l.await()
}

func (l *LanJoinScreen) update_game() {
	l.await()
	l.gameplay.client_update(&l.gameplay.p2)
	l.send_data(&l.gameplay.p2)
}

func (l *LanJoinScreen) await() {
	if mutex == 1 {
		return
	} else if mutex == 0 {
		address := <-l.addr_chan
		buffer := <-l.buffer_chan

		message := l.buffer_to_string(buffer)
		l.act_on_message(message, address)

		// set mutex back to 1 so we can await the next buffer
		mutex = 1

		go l.network.spawn_client_listener(l.addr_chan, l.buffer_chan)
	}
}

func (l *LanJoinScreen) buffer_to_string(buffer []byte) string {
	buffer_no_zeroes := bytes.Trim(buffer, "\u0000")
	string_no_whitespace := strings.TrimSpace(string(buffer_no_zeroes))
	return string_no_whitespace
}

func (l *LanJoinScreen) draw_pre() {
	rl.BeginDrawing()
	rl.DrawFPS(10, 10)
	rl.ClearBackground(rl.LightGray)

	waiting_text := "Waiting for host to start the game."
	waiting_text_size := rl.MeasureTextEx(roboto, waiting_text, 48, 1)
	top_text_size := rl.MeasureTextEx(roboto, l.top_text, 64, 1)
	bottom_text_size := rl.MeasureTextEx(roboto, l.bottom_text, 32, 1)

	rl.DrawTextEx(roboto, l.top_text, rl.NewVector2((screen_width/2)-(top_text_size.X/2), (screen_height/2)-(top_text_size.Y/2)-80), 64, 1, rl.Black)
	rl.DrawTextEx(roboto, l.bottom_text, rl.NewVector2((screen_width/2)-(bottom_text_size.X/2), (screen_height/2)-(bottom_text_size.Y/2)), 32, 1, rl.Black)

	if l.connection_state == awaiting_game_start {
		rl.DrawTextEx(roboto, waiting_text, rl.NewVector2((screen_width/2)-(waiting_text_size.X/2), (screen_height/2)-(waiting_text_size.Y/2)+80), 48, 1, rl.Black)
	}

	rl.EndDrawing()
}

func (l *LanJoinScreen) act_on_message(message string, address net.Addr) {
	if l.connection_state == no_connection {
		if message == send_player_info {
			l.top_text = fmt.Sprintf("Connected to %v", address.String())
			l.bottom_text = "Exchanging information..."
			l.network.send([]byte("john the gamer\"I am better than everyone\""))
			l.connection_state = awaiting_player_info
		}
	}
	if l.connection_state == awaiting_player_info {
		split_message := strings.SplitN(message, "\"", 3)
		if len(split_message) != 3 {
			l.network.send([]byte(bad_message))
		} else {
			player_name := split_message[0]
			player_quote := split_message[1]
			l.bottom_text = fmt.Sprintf("Your opponent is %v \"%v\"", player_name, player_quote)
			l.network.send([]byte("ACK"))
			l.connection_state = acknowledged_player_info
		}
	}
	if l.connection_state == acknowledged_player_info {
		if message != ready {
			l.network.send([]byte(bad_message))
		} else {
			l.connection_state = awaiting_game_start
		}
	}
	if l.connection_state == awaiting_game_start {
		if message != start_countdown {
			l.network.send([]byte(bad_message))
		} else {
			l.connection_state = started_game
			l.initialize_gameplay()
		}
	}
	if l.connection_state == started_game {
		//println(message)
		split_message := strings.SplitN(message, " ", 3)
		ball_x, error := strconv.ParseFloat(split_message[0], 32)
		if error != nil {
			l.network.send([]byte(bad_message))
			return
		}
		ball_y, error := strconv.ParseFloat(split_message[1], 32)
		if error != nil {
			l.network.send([]byte(bad_message))
			return
		}
		paddle_y, error := strconv.ParseFloat(split_message[2], 32)
		if error != nil {
			l.network.send([]byte(bad_message))
			return
		}

		l.gameplay.ball.x = float32(ball_x)
		l.gameplay.ball.y = float32(ball_y)
		l.gameplay.p1.y = float32(paddle_y)
	}
}

func (l *LanJoinScreen) initialize_gameplay() {
	l.gameplay = GameplayMultiplayer{}
	l.gameplay.initialize_gameobjects()
	l.update_func = l.update_game
	l.draw_func = l.gameplay.draw
}

func (l *LanJoinScreen) send_data(player_paddle *Paddle) {
	l.network.send([]byte(strconv.FormatFloat(float64(player_paddle.y), 'f', -1, 32)))
}
