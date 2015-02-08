package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type Client struct {
	conn net.Conn
	ch   chan<- string
}

func main() {
	ln, err := net.Listen("tcp", ":6000")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	msgchan := make(chan string)
	addchan := make(chan Client)
	rmchan := make(chan net.Conn)
	/* One 'handleMessages()' go routine */
	go handleMessages(msgchan, addchan, rmchan)
	/* a 'handleConnection()' go routine per each successful connection */
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConnection(conn, msgchan, addchan, rmchan)
	}
}

func handleConnection(conn net.Conn, msgchan chan<- string, addchan chan<- Client, rmchan chan<- net.Conn) {
	/* New string channel to this client. Client is 'conn & ch' */
	ch := make(chan string)
	/* messages received by client will be sent to this 'msgs' channel */
	msgs := make(chan string)

	/* Since client is successfully connected, add this client to the addchan channel. handleMessage() will add this client to the map */
	addchan <- Client{conn, ch}

	/* go routine to read messages recived by this client and send it to 'msgs' channel */
	go func() {
		defer close(msgs)
		buffcon := bufio.NewReader(conn)
		conn.Write([]byte("\033[1;30;41mWelcome to the fancy demo chat!\033[0m\r\nWhat is your nick? "))
		nick, _, err := buffcon.ReadLine()
		if err != nil {
			return
		}
		nickname := string(nick)
		conn.Write([]byte("Welcome, " + nickname + "!\r\n\r\n"))
		msgs <- "New user " + nickname + " has joined the chat room."
		/* Read the messages received by this client connection 'c'  and send it to 'msgs' channel */
		for {
			line, _, err := buffcon.ReadLine()
			if err != nil {
				break
			}
			msgs <- nickname + ": " + string(line)
		}
		msgs <- "User " + nickname + " left the chat room."
	}()

	/* 1. Send from 'msgs' channel to 'msgchan' channel. 2. If there is any data in 'ch' channel, write it to client connection 'c' */
LOOP:
	for {
		select {
		case msg, ok := <-msgs:
			if !ok {
				break LOOP
			}
			msgchan <- msg
		/* Any string in 'ch' channel, will be written back to client */
		case msg := <-ch:
			_, err := conn.Write([]byte(msg))
			if err != nil {
				break LOOP
			}
		}
	}
	conn.Close()
	fmt.Printf("Connection from %v closed.\n", conn.RemoteAddr())
	rmchan <- conn
}

func handleMessages(msgchan <-chan string, addchan <-chan Client, rmchan <-chan net.Conn) {
	clients := make(map[net.Conn]chan<- string)
	for {
		select {
		/* Any string in 'msgchan' channel, send to all the registered clients */
		case msg := <-msgchan:
			fmt.Printf("new message: %s\n", msg)
			for _, ch := range clients {
				go func(mch chan<- string) { mch <- "\033[1;33;40m" + msg + "\033[m\r\n" }(ch)
			}
		/* When new client is added by handleConnection(), add it to map */
		case client := <-addchan:
			fmt.Printf("New client: %v\n", client.conn)
			clients[client.conn] = client.ch
		/* When connection is closed by handleConnection(), remove it from the map */
		case conn := <-rmchan:
			fmt.Printf("Client disconnects: %v\n", conn)
			delete(clients, conn)
		}
	}
}
