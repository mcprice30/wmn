package network

import (
	"fmt"
	"math/rand"
	"net"

	"github.com/mcprice30/wmn/data"
)

const dropChance = 0.5

type ReceiverConnection struct {
	address data.ManetAddr
	conn    *net.UDPConn
}

func Listen(address data.ManetAddr) *ReceiverConnection {
	addr := ToUDPAddr(address)
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}

	return &ReceiverConnection{
		address: address,
		conn:    conn,
	}
}

func (c *ReceiverConnection) Reply(data []byte, addr net.Addr) {
	if rand.Float64() < dropChance {
		fmt.Println("Gremlin!")
		return
	}
	raddr := addr.(*net.UDPAddr)
	if _, err := c.conn.WriteToUDP(data, raddr); err != nil {
		panic(err)
	}
}

func (c *ReceiverConnection) Receive() ([]byte, net.Addr) {
	buffer := make([]byte, data.MaxPacketBytes)
	if n, addr, err := c.conn.ReadFrom(buffer); err != nil {
		panic(err)
	} else {
		return buffer[:n], addr
	}
}

func (c *ReceiverConnection) Close() {
	c.conn.Close()
}
