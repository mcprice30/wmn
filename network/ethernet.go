package network

import (
	"fmt"
	"math/rand"
	"net"

	"github.com/mcprice30/wmn/data"
)

// dropChance indicates the probability that a given packet is dropped by
// the ethernet connection. Note that the same probability is applied to both
// the data packet and the acknowledgement packet. For example, if this is 50%,
// then the probability of a completely successful transmission is only 25%.
const dropChance = 0.90

// EthernetConnection implements the Connection interface, simulating a network
// occurring over the default ethernet link.
type EthernetConnection struct {
	laddr data.ManetAddr
	conn  *net.UDPConn
}

// Bind will create an EthernetConnection that listens on the given address.
func Bind(address data.ManetAddr) *EthernetConnection {
	addr := ToUDPAddr(address)
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}

	return &EthernetConnection{
		laddr: address,
		conn:  conn,
	}
}

// Send will attempt to send the given packet to the address specified in the
// packet's header, as specified by the Connection interface.
func (c *EthernetConnection) Send(bytes []byte) {
	if rand.Float64() < dropChance {
		fmt.Println("Gremlin!")
		return
	}
	header := data.PacketHeaderFromBytes(bytes)
	raddr := ToUDPAddr(header.DestinationAddress)
	if _, err := c.conn.WriteToUDP(bytes, raddr); err != nil {
		panic(err)
	}
}

// Receive will listen at the given address for incoming packets, and will
// return any that it recieves, as specified by the Connection interface.
func (c *EthernetConnection) Receive() []byte {
	buffer := make([]byte, data.MaxPacketBytes)
	if n, err := c.conn.Read(buffer); err != nil {
		panic(err)
	} else {
		return buffer[:n]
	}
}

// Close will close this connection, as specified by the Connection interface.
func (c *EthernetConnection) Close() {
	c.conn.Close()
}
