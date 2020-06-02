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
		os.Exit(0)
	}
}

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr, msg []byte) {
	_, err := conn.WriteToUDP(msg, addr)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func main() {
	port := "10001"

	fmt.Println("Listening on", port)

	/* Hostname */
	hostname, err := os.Hostname()
	fmt.Println("Hostname=" + hostname)
	checkError(err)

	/* Lets prepare a address at any address at port 10001*/
	serverAddr, err := net.ResolveUDPAddr("udp", ":"+port)
	checkError(err)

	/* Now listen at selected port */
	conn, err := net.ListenUDP("udp", serverAddr)
	checkError(err)
	defer conn.Close()

	count := 0
	buf := make([]byte, 2048)

	for {
		count = count + 1

		n, addr, err := conn.ReadFromUDP(buf)
		fmt.Println("Received [", string(buf[0:n]), "] from ", addr, "[", count, "]")
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		msg := []byte(hostname + " " + string(buf[0:n]) + " " + addr.String() + " " + strconv.Itoa(count))

		go sendResponse(conn, addr, msg)
	}
}
