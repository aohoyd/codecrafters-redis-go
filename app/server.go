package main

import (
	"fmt"
	"io"
	"net"
	"os"
	// Uncomment this block to pass the first stage
	// "net"
	// "os"
)

func main() {
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println("Error listening: ", err.Error())
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting: ", err.Error())
		os.Exit(1)
	}

	buf := make([]byte, 1024)
	for {
		switch _, err := conn.Read(buf); {
		case err == io.EOF:
			break
		case err != nil:
			fmt.Println("Error reading: ", err.Error())
			os.Exit(1)
		}

		if _, err = io.WriteString(conn, "+PONG\r\n"); err != nil {
			fmt.Println("Error writing: ", err.Error())
			os.Exit(1)
		}
	}
}
