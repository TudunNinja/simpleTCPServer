package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const RECV_BUFFER_SIZE = 2048

func main() {
	port_num := os.Args[1]

	ln, _ := net.Listen("tcp", "127.0.0.1:"+port_num)

	writer := bufio.NewWriter(os.Stdout)

	for {
		conn, _ := ln.Accept()

		message := make([]byte, RECV_BUFFER_SIZE)

		n, _ := conn.Read(message)
		fmt.Println(n)

		writer.Write(message)
		writer.Flush()

		conn.Close()
	}
}
