package main

import (
	"fmt"

	"github.com/dilshodforever/5-oyimtixon/kafkaconnect"
	"github.com/dilshodforever/5-oyimtixon/server"
)

func main() {
	server.Connection()
	kafkaconnect.ConnectKAfka()
	fmt.Println(11111111111)

}
