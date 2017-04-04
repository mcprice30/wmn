package main

import (
	"fmt"
	"os"

	"github.com/mcprice30/wmn/config"
	"github.com/mcprice30/wmn/network"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Use:", os.Args[0], "<config file> <hostname>")
		os.Exit(127)
	}

	config.LoadConfig(os.Args[1], os.Args[2])
	conn := network.CreateManet()
	for {
		fmt.Println(conn.Receive())
	}
}
