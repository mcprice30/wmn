package config

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/mcprice30/wmn/data"
	"github.com/mcprice30/wmn/network"
)

var lastModified = time.Now()

// LoadConfig will load a configuration file with the given file name.
// Each line of the config file should contain:
//
// <manet addr> <name> <physical location:port> <x> <y> [links...]
func LoadConfig(fn, myHostname string) {
	file, err := os.Open(fn)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 1

	network.SetMyHostname(myHostname)

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())

		manetAddr := parseManetAddr(line[0], fn, lineNumber)
		hostname := line[1]
		actualLocation := line[2]
		neighbors := []data.ManetAddr{}
		for i := 5; i < len(line); i++ {
			neighbors = append(neighbors, parseManetAddr(line[i], fn, lineNumber))
		}

		network.SetAddress(manetAddr, actualLocation)
		network.SetHostname(hostname, manetAddr)

		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	go listenForLocationChanges(fn, myHostname)
}

func listenForLocationChanges(fn, hostname string) {
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
		if !info.ModTime().Equal(lastModified) {
			lastModified = info.ModTime()
			updateNodeInfo(fn, hostname)
		}
	}
}

func updateNodeInfo(fn, hostname string) {
	file, err := os.Open(fn)
	if err != nil {
		fmt.Println("Cannot read config file:", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)
	locations := map[data.ManetAddr]*data.Point{}
	neighbors := []data.ManetAddr{}
	myAddr := data.ManetAddr(0)
	lineNumber := 1
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		manetAddr := parseManetAddr(line[0], fn, lineNumber)
		x, err := strconv.ParseInt(line[3], 0, 32)
		if err != nil {
			fmt.Printf("%s [%d]: Invalid x '%s'\n", fn, lineNumber, line[3])
			os.Exit(1)
		}

		y, err := strconv.ParseInt(line[4], 0, 32)
		if err != nil {
			fmt.Printf("%s [%d]: Invalid y '%s'\n", fn, lineNumber, line[4])
			os.Exit(1)
		}
		locations[manetAddr] = &data.Point{
			X: int(x),
			Y: int(y),
		}

		if line[1] == hostname {
			myAddr = manetAddr
			for i := 5; i < len(line); i++ {
				neighbors = append(neighbors, parseManetAddr(line[i], fn, lineNumber))
			}
		}
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	nMap := map[data.ManetAddr]float64{}

	for _, neighbor := range neighbors {
		if dist := locations[myAddr].Dist(locations[neighbor]); dist < 100.0 {
			nMap[neighbor] = dist
		}
	}

	network.SetNeighbors(nMap)

}

func parseManetAddr(str, fn string, lineNumber int) data.ManetAddr {
	addrInt, err := strconv.ParseInt(str, 0, 16)
	if err != nil {
		fmt.Printf("%s [%d]: Invalid address '%s'\n", fn, lineNumber, str)
		os.Exit(1)
	}
	return data.ManetAddr(addrInt)
}
