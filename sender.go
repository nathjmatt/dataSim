package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
)

func NewSender(ipAddress string, portNumber uint) Sender {
	c := Sender{}
	i := net.ParseIP(ipAddress)
	if i == nil {
		panic("the ip address passed is not valid")
	}
	c.ipAddress = ipAddress

	if portNumber > 65535 {
		panic("not a valid port number")
	}
	c.portNumber = strconv.FormatUint(uint64(portNumber), 10)

	return c
}

type Sender struct {
	ipAddress  string
	portNumber string
}

func (c Sender) Start(inputChan <-chan []byte, wg *sync.WaitGroup) {
	defer wg.Done()

	conn, err := c.getUDPConn()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Sending packets to %s\n", conn.RemoteAddr().String())

	packetsSent := 0
	// Grab byte slices off channel
	for data := range inputChan {
		// Send the byte slice over the connection
		bytesWritten, err := conn.Write(data)
		if err != nil {
			fmt.Println(err)
		}
		packetsSent += 1
		_ = bytesWritten
		// fmt.Printf("Sent %d bytes to %s:%s\n", bytesWritten, c.ipAddress, c.portNumber)
	}
	fmt.Printf("Sent %d packets to %s:%s\n", packetsSent, c.ipAddress, c.portNumber)
}

// Get a UDP connection to send data to
func (c *Sender) getUDPConn() (*net.UDPConn, error) {
	// Define the IP address and port of the remote server
	addr, err := net.ResolveUDPAddr("udp", c.ipAddress+":"+c.portNumber)
	if err != nil {
		return nil, err
	}
	// Create a connection to the remote server
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return nil, err
	}

	return conn, err
}
