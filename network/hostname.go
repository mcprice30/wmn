package network

import (
	"net"

	"github.com/mcprice30/wmn/data"
)

var cache = map[data.ManetAddr]string{
	0x0001: "localhost:5009", // "Sensor Hub" (transmission only)
	0x0002: "localhost:5010", // "Display Hub"
	0x0003: "localhost:5011", // "Manet Node 1"
}

var myAddress data.ManetAddr = 0

func SetMyAddress(in data.ManetAddr) {
	myAddress = in
}

func GetMyAddress() data.ManetAddr {
	return myAddress
}

func ToUDPAddr(addr data.ManetAddr) *net.UDPAddr {
	if res, err := net.ResolveUDPAddr("udp", cache[addr]); err != nil {
		panic(err)
	} else {
		return res
	}
}

func SetAddress(addr data.ManetAddr, location string) {
	cache[addr] = location
}
