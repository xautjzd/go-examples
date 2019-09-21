package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8085")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Text to send: ")
	text, _ := reader.ReadString('\n')
	// conn.Write([]byte(text + "\n"))
	fmt.Fprintf(conn, text+"\n")
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Printf("Message received from server: %s", message)
}
