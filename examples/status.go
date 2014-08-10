package main

import (
	"fmt"
	"github.com/hangy/gooob"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]

	udpAddr, err := net.ResolveUDPAddr("up", service)
	checkError(err)

	clt, err := gooob.New(udpAddr, "dummy")
	checkError(err)

	response, err := clt.Rcon("status")
	checkError(err)

	fmt.Println(response)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}
