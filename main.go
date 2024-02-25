package main

import (
	"flag"
	"fmt"
	"net/netip"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const screen_width = 1280
const screen_height = 800

var roboto rl.Font
var game_state = title_screen

const (
	title_screen = iota
	lan_host_initialize
	lan_host
	lan_join
	lan_join_initialize
	initialize_gameplay
	gameplay
)

const (
	no_connection = iota
	correct_join_message
	awaiting_player_info
	acknowledged_player_info
	awaiting_game_start
	started_game
	recieved_player_info
	sent_player_info_ack
	requested_countdown
)

func main() {
	host_or_client_ptr := flag.String("lan", "", "Set this argument to 'host' or 'join'.")
	ip_address_ptr := flag.String("ip", "192.168.0.101", "Host IP address to connect to.")
	port_number_ptr := flag.String("port", "34788", "Host port to connect to or to host from.")
	flag.Parse()

	//var prof_machine string
	//if *host_or_client_ptr == "host" {
	//	prof_machine = "host"
	//}
	//if *host_or_client_ptr == "join" {
	//	prof_machine = "join"
	//}

	//f, err := os.OpenFile(fmt.Sprintf("%v_network_log", prof_machine), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	//if err != nil {
	//	log.Fatalf("error opening file: %v", err)
	//}
	//defer f.Close()
	//log.SetOutput(f)

	//prof, err := os.Create(fmt.Sprintf("%v_trace.out", prof_machine))
	//if err != nil {
	//	fmt.Fprintln(os.Stderr, err)
	//	os.Exit(1)
	//}

	//trace.Start(prof)
	//defer trace.Stop()
	//pprof.StartCPUProfile(prof)
	//defer pprof.StopCPUProfile()

	addr_port, parse_error := netip.ParseAddrPort(fmt.Sprintf("%v:%v", *ip_address_ptr, *port_number_ptr))
	if parse_error != nil {
		fmt.Fprintln(os.Stderr, parse_error)
		os.Exit(1)
	}

	var window_title string
	host_or_client := *host_or_client_ptr
	if host_or_client == "host" {
		game_state = lan_host_initialize
		window_title = "UDPong Server"
	}
	if host_or_client == "join" {
		game_state = lan_join_initialize
		window_title = "UDPong Client"
	}

	rl.InitWindow(screen_width, screen_height, window_title)
	//rl.SetTargetFPS(200)
	//rl.SetWindowState(rl.FlagVsyncHint)
	defer rl.CloseWindow()

	var gameplay_singleplayer GameplaySingeplayer
	var lan_host_screen LanHostScreen
	var lan_join_screen LanJoinScreen

	//m := &runtime.MemStats{}

	for !rl.WindowShouldClose() {
		delta = rl.GetFrameTime()
		//runtime.ReadMemStats(m)
		//fmt.Printf("Num GC: %d Heap Allocated: %d Heap System: %d Heap Objects: %d Heap Released: %d;\n", m.NumGC, m.HeapAlloc, m.HeapSys, m.HeapObjects, m.HeapReleased)

		switch game_state {
		case title_screen:
			var title_screen = new_singeplayer_game()
			title_screen.update()
			title_screen.draw()
		case lan_host_initialize:
			lan_host_screen = LanHostScreen{}
			lan_host_screen.initialize_server(addr_port)
			game_state = lan_host
		case lan_host:
			go lan_host_screen.update_func()
			lan_host_screen.draw_func()
		case lan_join_initialize:
			lan_join_screen = LanJoinScreen{}
			lan_join_screen.initialize_connection(addr_port)
			game_state = lan_join
		case lan_join:
			go lan_join_screen.update_func()
			lan_join_screen.draw_func()
		case initialize_gameplay:
			gameplay_singleplayer = GameplaySingeplayer{}
			gameplay_singleplayer.initialize_gameobjects()
			game_state = gameplay
		case gameplay:
			gameplay_singleplayer.update()
			gameplay_singleplayer.draw()
		}
	}
}
