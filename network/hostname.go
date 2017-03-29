package network

import (
	"net"

	"github.com/mcprice30/wmn/data"
)

// cache contains a mapping of ManetAddresses to actual addresses.
var cache = map[data.ManetAddr]string{
	0x0001: "localhost:5009", // "Sensor Hub" (transmission only)
	0x0002: "localhost:5010", // "Display Hub"
	0x0003: "localhost:5011", // "Manet Node 1"
}

// myAddress indicates which manet address belongs to this device.
var myAddress data.ManetAddr = 0

// SetMyAddress will set whatever address this device owns.
func SetMyAddress(in data.ManetAddr) {
	myAddress = in
}

// GetMyAddress will return the address that this device has been assigned.
func GetMyAddress() data.ManetAddr {
	return myAddress
}

// ToUDPAddr will convert a manet address to a udp address, to be used for
// actually sending the packets.
func ToUDPAddr(addr data.ManetAddr) *net.UDPAddr {
	if res, err := net.ResolveUDPAddr("udp", cache[addr]); err != nil {
		panic(err)
	} else {
		return res
	}
}

// SetAddress will map a manet address, to a location string, such as
// 'tux054:10010'
func SetAddress(addr data.ManetAddr, location string) {
	cache[addr] = location
}
