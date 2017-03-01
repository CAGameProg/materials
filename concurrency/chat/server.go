package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

var clients []Client

type Client struct {
	conn net.Conn
	name string
}

func RemoveClient(toRemove Client) {
	for i, c := range clients {
		if c.conn == toRemove.conn {
			clients = append(clients[:i], clients[i+1:]...)
		}
	}
}

func SendMessageFrom(client Client, message string) {
	for _, c := range clients {
		if c.conn != client.conn {
			c.conn.Write([]byte(message))
		}
	}
}

func HandleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	client := Client{
		conn: conn,
		name: name,
	}

	clients = append(clients, client)

	connectMsg := name + " has connected\n"
	fmt.Print(connectMsg)
	SendMessageFrom(client, connectMsg)
	for {
		// will listen for message to process ending in newline (\n)
		inMessage, err := reader.ReadString('\n')
		if err != nil {
			// User disconnected
			disconnectMsg := name + " has disconnected\n"
			fmt.Print(disconnectMsg)
			SendMessageFrom(client, disconnectMsg)
			RemoveClient(client)
			return
		}
		outMessage := name + ": " + string(inMessage)
		fmt.Print(outMessage)
		SendMessageFrom(client, outMessage)
	}
}

func main() {
	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// run loop forever (or until ctrl-c)
	for {
		// accept connection on port
		conn, _ := ln.Accept()

		go HandleConnection(conn)
	}
}
