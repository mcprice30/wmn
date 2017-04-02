package main

import (
	"fmt"
	"os"

	"github.com/mcprice30/wmn/config"
	"github.com/mcprice30/wmn/network"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Use:", os.Args[0], "<hostname>")
		os.Exit(127)
	}

	config.LoadConfig("config_test.txt", os.Args[1])
	conn := network.CreateManet()
	for {
		fmt.Println(conn.Receive())
	}
}
