package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func main() {
	// SERVER_ADDRESS comes from environment variable 
	Server := os.Getenv("SERVER_ADDRESS")
	fmt.Println("SERVER_ADDRESS=" + Server)

	// Open new UDP "connection" on each send so that source port number changes
	i := 0
	buf := make([]byte, 1024)
	for {
		serverAddr, err := net.ResolveUDPAddr("udp", Server+":10001")
		checkError(err)

		conn, err := net.DialUDP("udp", nil, serverAddr)
		checkError(err)

		msg := strconv.Itoa(i)
		fmt.Fprintf(conn, msg)
		i++

		n, addr, err := conn.ReadFromUDP(buf)
		fmt.Println("Received [", string(buf[0:n]), "] from ", addr)

		if err != nil {
			fmt.Println("Error: ", err)
		}

		// Close each "connection" so that client does not run out of ephemeral ports
		conn.Close()

		//time.Sleep(time.Second * 1)
	}
}
