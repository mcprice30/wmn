package network

import (
	"fmt"
	"math/rand"
	"net"

	"github.com/mcprice30/wmn/data"
)

const dropChance = 0.90

type EthernetConnection struct {
	laddr data.ManetAddr
	conn  *net.UDPConn
}

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

func (c *EthernetConnection) Receive() []byte {
	buffer := make([]byte, data.MaxPacketBytes)
	if n, err := c.conn.Read(buffer); err != nil {
		panic(err)
	} else {
		return buffer[:n]
	}
}

func (c *EthernetConnection) Close() {
	c.conn.Close()
}
