package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

/*What's wrong with client being a daemon ?? */
func main() {
	androidConnection, err := connectToAndroidHub()
	if err != nil {
		fmt.Println("Couldn't connect to android")
		goto returnHadler
	}

returnHandler:
	return
}

func connectToAndroidHub() (conn net.Conn, err error) {
	conn, err = net.Dial("tcp", "192.168.1.4:9999")
	if err != nil {
		fmt.Println("Could not connect")
	}
	return
	/*
		stdinReader := bufio.NewReader(os.Stdin)
		socketReader := bufio.NewReader(conn)
		receiveString := ""
		for {
			fmt.Print("What do you want to send -->")
			txt, _ := stdinReader.ReadString('\n')
			conn.Write([]byte(txt))
			receiveString, _ = socketReader.ReadString('\n')
			fmt.Print(receiveString)
		}
	*/

}
