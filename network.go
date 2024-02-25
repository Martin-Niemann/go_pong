package main

import (
	"fmt"
	"net"
	"net/netip"
	"os"
)

type Network struct {
	udp_server net.PacketConn
	udp_addr   *net.UDPAddr
	udp_conn   *net.UDPConn
	error      error
}

func (n *Network) initialize_host(addrport netip.AddrPort) {
	mutex = 1
	n.udp_server, n.error = net.ListenPacket("udp", fmt.Sprintf("192.168.0.101:%v", addrport.Port()))

	if n.error != nil {
		fmt.Fprintln(os.Stderr, n.error)
		os.Exit(1)
	}

	//defer n.udp_server.Close()
}

func (n *Network) initialize_client(addrport netip.AddrPort) {
	mutex = 1
	n.udp_addr, n.error = net.ResolveUDPAddr("udp", addrport.String())
	if n.error != nil {
		fmt.Fprintln(os.Stderr, n.error)
		os.Exit(1)
	}

	n.udp_conn, n.error = net.ListenUDP("udp", &net.UDPAddr{IP: net.IP{192, 168, 0, 101}, Port: 0})
	if n.error != nil {
		fmt.Fprintln(os.Stderr, n.error)
		os.Exit(1)
	}

	//defer n.udp_conn.Close()

	println(n.udp_conn.LocalAddr().String())
	println(n.udp_addr.String())
	// send connect message
	n.send([]byte(connect_message))
}

func (n *Network) spawn_host_listener(address_chan chan net.Addr, message_chan chan []byte) {
	buffer := make([]byte, 128)

	// this (hopefully) blocks
	_, address, error := n.udp_server.ReadFrom(buffer)

	if error != nil {
		fmt.Fprintln(os.Stderr, error)
	}

	//log.Println(address, string(buffer))
	//fmt.Println(address, string(buffer))

	mutex = 0
	// this returns(?) for some reason?
	address_chan <- address
	message_chan <- buffer
}

func (n *Network) spawn_client_listener(address_chan chan net.Addr, message_chan chan []byte) {
	buffer := make([]byte, 128)

	// this (hopefully) blocks
	_, address, error := n.udp_conn.ReadFrom(buffer)

	if error != nil {
		fmt.Fprintln(os.Stderr, error)
	}

	//log.Println(address, string(buffer))
	//fmt.Println(address, string(buffer))

	mutex = 0
	// this returns(?) for some reason?
	address_chan <- address
	message_chan <- buffer
}

func (n *Network) respond(address net.Addr, response []byte) {
	n.udp_server.WriteTo(response, address)
}

func (n *Network) send(message []byte) {
	_, error := n.udp_conn.WriteToUDP(message, n.udp_addr)
	if error != nil {
		fmt.Fprintln(os.Stderr, error)
		os.Exit(1)
	}
}
