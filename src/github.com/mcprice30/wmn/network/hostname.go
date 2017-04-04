package network

import (
	"net"

	"github.com/mcprice30/wmn/data"
)

// cache contains a mapping of ManetAddresses to actual addresses.
var cache = map[data.ManetAddr]string{}

// dsn contains a mapping of manet host names to their address.
var dns = map[string]data.ManetAddr{}

// myHostname indicates what name in the manet this node is.
var myHostname string = ""

// GetMyAddress will return the address that this device has been assigned.
func GetMyAddress() data.ManetAddr {
	return dns[myHostname]
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

// GetAddrFromHostname will get the hostname associated with the given
// manet address.
func GetAddrFromHostname(hostname string) data.ManetAddr {
	return dns[hostname]
}

// SetAddress will map a manet address, to a location string, such as
// 'tux054:10010'
func SetAddress(addr data.ManetAddr, location string) {
	cache[addr] = location
}

// SetHostname will map the given hostname to the appropriate manet address.
func SetHostname(hostname string, addr data.ManetAddr) {
	dns[hostname] = addr
}

// SetMyHostname will remember what this device's hostname is.
func SetMyHostname(in string) {
	myHostname = in
}
