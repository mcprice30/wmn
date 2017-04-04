package main

import (
	"fmt"
	"os"

	"github.com/mcprice30/wmn/chief"
	"github.com/mcprice30/wmn/config"
)

func main() {

	if len(os.Args) != 5 {
		fmt.Println("Use:", os.Args[0], "<config file> <error file> <hostname> <display port>")
		os.Exit(127)
	}

	config.LoadConfig(os.Args[1], os.Args[3])
	config.ListenForErrorChanges(os.Args[2])
	chief.RunListener(os.Args[4])
}
