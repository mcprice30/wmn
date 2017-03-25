package network

import (
	"fmt"
	"math/rand"
	"net"

	"github.com/mcprice30/wmn/data"
)

type SenderConnection struct {
	src  data.ManetAddr
	dst  data.ManetAddr
	conn net.Conn
}

func Connect(src, dst data.ManetAddr) *SenderConnection {
	srcAddr := ToUDPAddr(src)
	dstAddr := ToUDPAddr(dst)
	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		panic(err)
	}

	return &SenderConnection{
		src:  src,
		dst:  dst,
		conn: conn,
	}
}

func (c *SenderConnection) Send(data []byte) {
	if rand.Float64() < dropChance {
		fmt.Println("Gremlin!")
		return
	}
	if _, err := c.conn.Write(data); err != nil {
		panic(err)
	}
}

func (c *SenderConnection) Receive() []byte {
	buffer := make([]byte, data.MaxPacketBytes)
	if n, err := c.conn.Read(buffer); err != nil {
		panic(err)
	} else {
		return buffer[:n]
	}
}

func (c *SenderConnection) Close() {
	err := c.conn.Close()
	if err != nil {
		panic(err)
	}
}
