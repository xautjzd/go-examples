package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	fmt.Println("Lauching server...")
	ln, err := net.Listen("tcp", ":8085")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGPIPE, syscall.SIGUSR1)
	go func() {
		fmt.Printf("received signal: %s\n", <-sigCh)
		os.Exit(0)
	}()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("failed accepting a connection request: %v", err)
			continue
		}
		go handler(conn)
	}
}

func handler(conn net.Conn) {
	defer conn.Close()
	message, _ := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn)).ReadString('\n')
	fmt.Printf("Message recevied: %s", message)
	newMsg := strings.ToUpper(message)
	conn.Write([]byte(newMsg))
}
