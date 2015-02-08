package main

import (
	"fmt"
	"github.com/kteza1/home_nextgen/limitlessled"
)

func main() {
	bridge, err := limitlessled.Dial("192.168.0.100:8899")
	fmt.Println(bridge)
	if err != nil {
		fmt.Println("Something wrong")
		return
	}
	bridge.SendCommand(limitlessled.ALL_OFF)

}
