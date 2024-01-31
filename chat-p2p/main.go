package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {

	isServer := flag.Bool("s", false, "run as server")
	flag.Parse()
	addr := flag.Arg(0)
	if addr == "" {
		usage()
	}

	if isServer != nil && *isServer {
		startServer(addr)
	} else {
		startClient(addr)
	}
}

func usage() {
	fmt.Println("usage: chat-p2p [-s] <address:port>")
	os.Exit(1)
}

func startServer(addr string) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	fmt.Println("Listening on:", l.Addr().String())
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Println("Client connected:", conn.RemoteAddr().String())
		handleConn(conn)
	}
}

func startClient(addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to server:", conn.RemoteAddr().String())
	handleConn(conn)
}

func handleConn(conn net.Conn) {
	inputDoneChan := make(chan bool, 1)
	messagesDoneChan := make(chan bool, 1)
	go handleUserInput(conn, inputDoneChan)
	go handleMessages(conn, messagesDoneChan)

	<-inputDoneChan
	<-messagesDoneChan

	fmt.Println("Connection closed:", conn.RemoteAddr())
}

func handleUserInput(conn net.Conn, done chan bool) {
	for {
		fmt.Printf("> ")
		r := bufio.NewReader(os.Stdin)
		message, err := r.ReadString('\n')
		if err != nil {
			done <- true
			break
		}
		// ignore empty messages
		if message == "\n" {
			continue
		}
		w := bufio.NewWriter(conn)
		_, err = w.WriteString(fmt.Sprintf("%s\n", message))
		if err != nil {
			done <- true
			break
		}
		err = w.Flush()
		if err != nil {
			done <- true
			break
		}
	}
}

func handleMessages(conn net.Conn, done chan bool) {
	for {
		r := bufio.NewReader(conn)
		msg, err := r.ReadString('\n')
		if err != nil {
			done <- true
			break
		}
		fmt.Printf("\n< %s", msg)
	}
}
