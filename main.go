package main

import (
	"fmt"
	"net"

	"github.com/acidlemon/rocket-sample/sample"
	"github.com/lestrrat/go-server-starter-listener"
)

func main() {
	//	pwd, _ := os.Getwd()

	fmt.Println("Launch succeeded!")

	listener, _ := ss.NewListener()
	if listener == nil {
		// Fallback if not running under Server::Starter
		var err error
		listener, err = net.Listen("tcp", ":8080")
		if err != nil {
			panic("Failed to listen to port 8080")
		}
	}

	sample.Start(listener)
}
