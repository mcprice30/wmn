package main

import (
	"fmt"
	"os"

	"github.com/mcprice30/wmn/chief"
	"github.com/mcprice30/wmn/config"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Use:", os.Args[0], "<hostname> <display port>")
		os.Exit(127)
	}

	config.LoadConfig("config_test.txt", os.Args[1])
	config.ListenForErrorChanges("error_test.txt")
	chief.RunListener(os.Args[2])
}
