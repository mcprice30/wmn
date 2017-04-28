package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Line struct {
	addr      string
	name      string
	realAddr  string
	x         int
	y         int
	neighbors []string
}

func ParseLine(line string) *Line {
	fields := strings.Fields(line)
	x, err := strconv.ParseInt(fields[3], 0, 32)
	if err != nil {
		panic(err)
	}
	y, err := strconv.ParseInt(fields[4], 0, 32)
	if err != nil {
		panic(err)
	}

	return &Line{
		addr:      fields[0],
		name:      fields[1],
		realAddr:  fields[2],
		x:         int(x),
		y:         int(y),
		neighbors: []string{},
	}
}

func (l *Line) String() string {
	return fmt.Sprintf("%s %s %s %d %d %s", l.addr, l.name, l.realAddr, l.x, l.y, strings.Join(l.neighbors, " "))
}

func ReadLines(fn string) []*Line {
	out := []*Line{}
	file, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		out = append(out, ParseLine(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return out
}

func ChangeLines(lines []*Line) {
	for _, line := range lines {
		if line.name != "Sensor" && line.name != "Display" {
			line.x += rand.Intn(7) - 3
			if line.x < 0 {
				line.x = 0
			} else if line.x > 300 {
				line.x = 300
			}

			line.y += rand.Intn(7) - 3
			if line.y < 0 {
				line.y = 0
			} else if line.y > 300 {
				line.y = 300
			}
		}
	}

	for i := range lines {
		for j := range lines {
			if i == j {
				continue
			}

			dx := lines[i].x - lines[j].x
			dy := lines[i].y - lines[j].y
			if dx*dx+dy*dy <= 100*100 {
				lines[i].neighbors = append(lines[i].neighbors, lines[j].addr)
			}
		}
	}
}

func SaveLines(fn string, lines []*Line) {

	out := ""
  for _, line := range lines {
    write := fmt.Sprintf("%s\n", line)
    fmt.Printf("Writing to %s %s", fn, write)
		out += write
  }

	if err := ioutil.WriteFile(fn, []byte(out), 0755); err != nil {
		panic(err)
	}
}

func main() {
  ticker := time.NewTicker(20 * time.Second)
  for {
		fmt.Println("Updating config file!")
    lines := ReadLines("config.txt")
    ChangeLines(lines)
    SaveLines("config.txt", lines)
    <- ticker.C
  }
}
