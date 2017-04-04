package network

import (
	"fmt"
	"net"

	"github.com/mcprice30/wmn/data"
)

// EthernetConnection implements the Connection interface, simulating a network
// occurring over the default ethernet link.
type EthernetConnection struct {
	conn *net.UDPConn
}

// Bind will create an EthernetConnection that listens on the given address.
func Bind(conn *net.UDPConn) *EthernetConnection {
	return &EthernetConnection{
		conn: conn,
	}
}

// Send will attempt to send the given packet to the address specified in the
// packet's header, as specified by the Connection interface.
func (c *EthernetConnection) Send(bytes []byte) {
	if dropFixedRate() {
		fmt.Println("DROP ETHERNET!")
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
