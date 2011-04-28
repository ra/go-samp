package main

import (
	"fmt"
	"net"
)

func getData(tcpConn *net.TCPConn, done chan bool) {
	for {
		var buffer = make([]byte, 256)
		bytesRead, err := tcpConn.Read(buffer)
		if err != nil {
			fmt.Printf("Error = %v\n", err)
			break
		}
		//
		fmt.Println("Bytes Read", bytesRead)
		bufData := buffer[0:bytesRead]
		fmt.Println("Buffer", bufData)
		var data = string(bufData)
		fmt.Println("Data Read", data)
	}
	//
	err := tcpConn.Close()
	if err != nil {
		fmt.Printf("Error = %v\n", err)
		panic("wtf05")
	}
	//
	done <- true
}

func main() {
	fmt.Println("Start .....")
	//
	// Get a TCP Address
	//
	tcpAddress, err := net.ResolveTCPAddr("", "localhost:45678")
	if err != nil {
		fmt.Printf("Error = %v\n", err)
		panic("wtf01")
	}
	fmt.Printf("tcpAddress = %v\n", tcpAddress)
	//
	// Get a TCP Listener
	//
	listener, err := net.ListenTCP("tcp", tcpAddress)
	if err != nil {
		fmt.Printf("Error = %v\n", err)
		panic("wtf02")
	}
	fmt.Printf("listener = %v\n", listener)
	//
	// Accept a connection.
	//
	tcpConn, err := listener.AcceptTCP()
	if err != nil {
		fmt.Printf("Error = %v\n", err)
		panic("wtf03")
	}
	fmt.Printf("connection = %v\n", tcpConn)
	waitFor := make(chan bool)
	go getData(tcpConn, waitFor)
	<-waitFor
	//
	err = listener.Close()
	if err != nil {
		fmt.Printf("Error = %v\n", err)
		panic("wtf06")
	}
	//
	fmt.Println("End .....")
}
