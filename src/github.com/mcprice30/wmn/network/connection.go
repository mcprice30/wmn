// Package network handles the network layer of the system, including both the
// ethernet and manet connections and the gremlin functions.
package network

// Connection is used to encapsulate a network interface. The transport layer
// is responsible for ensuring reliablity on top of this, which will unreliably
// attempt to communicate data over the network.
type Connection interface {
	// Send will attempt to send the given bytes over the network. In order for
	// the routing to work correctly, the source and destination manet addresses
	// should be included in the byte's header, as defined in Bytes.PacketHeader.
	Send(bytes []byte)

	// Receive will block until bytes arrive, at which point it will return
	// the received bytes, which should be marshallable into a packet.
	Receive() []byte

	// Close will close the connection. This should be called once the connection
	// will no longer be used.
	Close()
}
