package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.1.4:9999")

	if err != nil {
		fmt.Println("Could not connect")
		return
	}
	stdinReader := bufio.NewReader(os.Stdin)
	socketReader := bufio.NewReader(conn)
	receiveString := ""
	for {
		fmt.Print("What do you want to send -->")
		/* '\n' incluede here from stdin */
		txt, _ := stdinReader.ReadString('\n')
		conn.Write([]byte(txt))
		receiveString, _ = socketReader.ReadString('\n')
		fmt.Print(receiveString)
	}
}
