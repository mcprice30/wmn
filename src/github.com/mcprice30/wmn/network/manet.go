package network

import (
	"fmt"
	"net"

	"github.com/mcprice30/wmn/data"
)

const cacheDepth = 32

var myNeighbors map[data.ManetAddr]float64 = map[data.ManetAddr]float64{}

// ManetConnection implements a connection over the Manet.
type ManetConnection struct {
	laddr data.ManetAddr
	conn  *net.UDPConn
	cache map[uint16]uint16
}

// BindManet will instantiate a connection to the manet on the address specified
// by this process's local address, as determined by SetMyAddress.
func BindManet(conn *net.UDPConn) *ManetConnection {

	out := &ManetConnection{
		laddr: GetMyAddress(),
		conn:  conn,
		cache: map[uint16]uint16{},
	}

	go out.HelloLoop()
	return out
}

// BindManet will instantiate a connection to the manet on the address specified
// by this process's local address, as determined by SetMyAddress.
func CreateManet() *ManetConnection {

	addr := ToUDPAddr(GetMyAddress())
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}

	return BindManet(conn)
}

// SetNeighbors will update all the neighbors of the given simulated manet node
// to be the set of neighbors with the given addresses.
func SetNeighbors(neighbors map[data.ManetAddr]float64) {
	myNeighbors = neighbors
}

// Send will attempt to transmit the given packet bytes over the manet, as
// specified by the Connection interface.
func (c *ManetConnection) Send(bytes []byte) {
	for neighbor := range myNeighbors {
		if dropManetChance() {
			continue
		}
		raddr := ToUDPAddr(neighbor)
		if _, err := c.conn.WriteToUDP(bytes, raddr); err != nil {
			panic(err)
		}
	}
}

// Recieve will attempt to recieve packets on the address this connection
// was set up on, as specified by the Connection interface.
func (c *ManetConnection) Receive() []byte {
	for {
		buffer := make([]byte, data.MaxPacketBytes)
		if n, err := c.conn.Read(buffer); err != nil {
			panic(err)
		} else {
			header := data.PacketHeaderFromBytes(buffer[:n])

			if header.PacketType == data.PacketTypeHello {
				handleHelloPacket(buffer[:n])
				continue
			}

			if header.DestinationAddress == c.laddr {
				return buffer[:n]
			} else if mpr, exists := neighborTable.Selectors[header.SourceAddress]; exists && mpr {
				fmt.Printf("0x%04x: forwarding as MPR for 0x%04x\n", GetMyAddress(), header.SourceAddress)
				c.forward(buffer[:n])
			}
		}
	}
}

// inCache returns true iff the given sequence number and send key have already
// been transmitted by this device.
func (c *ManetConnection) inCache(seq, sendKey uint16) bool {
	return (c.cache[seq] == sendKey) || c.cache[seq] == (sendKey|0x8000)
}

// forward will forward the given packet to all neighboring manet nodes.
func (c *ManetConnection) forward(bytes []byte) {
	incomingHeader := data.PacketHeaderFromBytes(bytes)

	cached := c.inCache(incomingHeader.SequenceNumber, incomingHeader.SendKey)

	if cached || incomingHeader.TTL <= 1 {
		return
	}

	c.cache[incomingHeader.SequenceNumber] = incomingHeader.SendKey
	delete(c.cache, incomingHeader.SequenceNumber-cacheDepth)

	outgoingHeader := &data.PacketHeader{
		SourceAddress:      incomingHeader.SourceAddress,
		DestinationAddress: incomingHeader.DestinationAddress,
		PreviousHop:        GetMyAddress(),
		TTL:                incomingHeader.TTL - 1,
		PacketType:         incomingHeader.PacketType,
		SequenceNumber:     incomingHeader.SequenceNumber,
		NumBytes:           incomingHeader.NumBytes,
		SendKey:						incomingHeader.SendKey,
	}

	for i, b := range outgoingHeader.ToBytes() {
		bytes[i] = b
	}

	for neighbor := range myNeighbors {
		if neighbor == incomingHeader.PreviousHop {
			continue
		}
		raddr := ToUDPAddr(neighbor)
		if _, err := c.conn.WriteToUDP(bytes, raddr); err != nil {
			panic(err)
		}
	}
}

// Close will close the connection, as specified in the Connection interface.
func (c *ManetConnection) Close() {
	c.conn.Close()
}
