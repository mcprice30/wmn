package config

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/mcprice30/wmn/network"
)

var errorLastModified = time.Now()

func ListenForErrorChanges(fn string) {
	go listenForErrorChanges(fn)
}

func listenForErrorChanges(fn string) {
	fmt.Println("Start listen!")
	interval, _ := time.ParseDuration("3s")
	ticker := time.NewTicker(interval)
	for {
		<-ticker.C
		info, err := os.Stat(fn)
		if err != nil {
			fmt.Println("Cannot get info on config file:", err)
			os.Exit(1)
		}
		if !info.ModTime().Equal(errorLastModified) {
			errorLastModified = info.ModTime()
			updateErrorInfo(fn)
		}
	}
}

func updateErrorInfo(fn string) {
	file, err := os.Open(fn)
	if err != nil {
		fmt.Println("Cannot read config file:", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}
	if errorPct, err := strconv.ParseFloat(line, 64); err != nil {
		fmt.Println("Cannot get error percentage: ", err)
	} else {
		fmt.Printf("New error rate: %f%%\n", errorPct*100.0)
		network.SetDropChance(errorPct)
	}
	scanner.Scan()
	line = scanner.Text()
	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}
	if errorPct, err := strconv.ParseFloat(line, 64); err != nil {
		fmt.Println("Cannot get manet error percentage: ", err)
	} else {
		fmt.Printf("New manet error rate: %f%%\n", errorPct*100.0)
		network.SetManetDropChance(errorPct)
	}
}
