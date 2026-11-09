package main

import (
	"fmt"
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	_, err := conn.Write([]byte("OK\n"))
	if err != nil {
		fmt.Println("Error writing to connection:", err)
		return
	}
}

func run() error {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}
	defer listener.Close()
	fmt.Println("Server is listening on port 8080...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		handleConnection(conn)
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
