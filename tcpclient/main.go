package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("Please provide host:port to connect to")
		os.Exit(1)
	}

	// Resolve the string address to a TCP address
	tcpAddr, err := net.ResolveTCPAddr("tcp", os.Args[1])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Connect to the address with tcp
	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Send a message to the server
	_, err = conn.Write([]byte("Hello TCP Server"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Read from the connection untill a new line is send
	var buf [512]byte
	_, err = conn.Read(buf[0:])
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the data read from the connection to the terminal
	fmt.Print("> ", string(buf[0:]))
}
