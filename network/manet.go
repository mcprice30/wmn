package network

import (
	"fmt"
	"math/rand"
	"net"

	"github.com/mcprice30/wmn/data"
)

const cacheDepth = 32

type ManetConnection struct {
	laddr     data.ManetAddr
	conn      *net.UDPConn
	neighbors []data.ManetAddr
	cache     map[uint16]map[uint16]bool
}

func BindManet() *ManetConnection {
	addr := ToUDPAddr(GetMyAddress())
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}

	return &ManetConnection{
		laddr:     GetMyAddress(),
		conn:      conn,
		neighbors: []data.ManetAddr{},
		cache:     map[uint16]map[uint16]bool{},
	}
}

func (c *ManetConnection) SetNeighbors(neighbors []data.ManetAddr) {
	c.neighbors = neighbors
}

func (c *ManetConnection) Send(bytes []byte) {
	if rand.Float64() < dropChance {
		fmt.Println("Gremlin!")
		return
	}
	for _, neighbor := range c.neighbors {
		raddr := ToUDPAddr(neighbor)
		if _, err := c.conn.WriteToUDP(bytes, raddr); err != nil {
			panic(err)
		}
	}
}

func (c *ManetConnection) Receive() []byte {
	for {
		buffer := make([]byte, data.MaxPacketBytes)
		if n, err := c.conn.Read(buffer); err != nil {
			panic(err)
		} else {
			header := data.PacketHeaderFromBytes(buffer[:n])
			if header.DestinationAddress == c.laddr {
				return buffer[:n]
			} else {
				c.forward(buffer[:n])
			}
		}
	}
}

func (c *ManetConnection) forward(bytes []byte) {
	incomingHeader := data.PacketHeaderFromBytes(bytes)

	inCache := false
	if seqCache, exists := c.cache[incomingHeader.SequenceNumber]; exists {
		_, inCache = seqCache[incomingHeader.SendKey]
	} else {
		c.cache[incomingHeader.SequenceNumber] = map[uint16]bool{}
	}

	if inCache || incomingHeader.TTL <= 1 {
		return
	}

	c.cache[incomingHeader.SequenceNumber][incomingHeader.SendKey] = true
	delete(c.cache, incomingHeader.SequenceNumber-cacheDepth)
	delete(c.cache[incomingHeader.SequenceNumber],
		incomingHeader.SendKey-cacheDepth)

	outgoingHeader := &data.PacketHeader{
		SourceAddress:      incomingHeader.SourceAddress,
		DestinationAddress: incomingHeader.DestinationAddress,
		PreviousHop:        GetMyAddress(),
		TTL:                incomingHeader.TTL - 1,
		PacketType:         incomingHeader.PacketType,
		SequenceNumber:     incomingHeader.SequenceNumber,
		NumBytes:           incomingHeader.NumBytes,
	}

	for i, b := range outgoingHeader.ToBytes() {
		bytes[i] = b
	}

	for _, neighbor := range c.neighbors {
		if neighbor == incomingHeader.PreviousHop {
			continue
		}
		raddr := ToUDPAddr(neighbor)
		if _, err := c.conn.WriteToUDP(bytes, raddr); err != nil {
			panic(err)
		}
	}
}

func (c *ManetConnection) Close() {
	c.conn.Close()
}
