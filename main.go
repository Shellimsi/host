package main

import (
	"io"
	"log"
	"net"
)

var agents = make(map[string]*Agent)

type Agent struct {
	Connection net.Conn
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	log.Println("Server is running...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	log.Printf("New connection: %s", c.RemoteAddr().String())

	for _, agent := range agents {
		// Copy stdin to the pty and the pty to stdout.
		go func() { _, _ = io.Copy(agent.Connection, c) }()
		_, _ = io.Copy(c, agent.Connection)
	}

	agents[c.RemoteAddr().String()] = &Agent{Connection: c}
}
