package transport

import (
	"sync"

	"github.com/mcprice30/wmn/data"
	"github.com/mcprice30/wmn/network"
)

const maxBufferSize = 64

// ReliableReciever is used to listen reliably to packets on the given
// connection.
type ReliableReceiver struct {
	selector *Selector
	buffer          map[uint16]*data.DataPacket
	bufferLock      *sync.Mutex
	nextSequenceNum uint16
	receivedChan    chan *data.DataPacket
}

// CreateReliableReceiver will instantiate a new reliable receiver listening and
// transmitting the given manet address.
func CreateReliableReceiver() *ReliableReceiver {
	// SWITCHED TO MANET
	selector := CreateSelector()
//	conn.SetNeighbors([]data.ManetAddr{0x0003})
	// END SWITCHED TO MANET
	out := &ReliableReceiver{
		selector:            selector,
		buffer:          map[uint16]*data.DataPacket{},
		bufferLock:      &sync.Mutex{},
		nextSequenceNum: 0,
		receivedChan:    make(chan *data.DataPacket, maxBufferSize),
	}
	go out.runListen()
	return out
}

// Listen will block until the next packet is received, and then return that
// packet.
func (rr *ReliableReceiver) Listen() *data.DataPacket {
	return <-rr.receivedChan
}

// Close will close the connection, once transmission is done.
func (rr *ReliableReceiver) Close() {
	rr.selector.Close()
}

// runListen will listen on the given address, sending acknowledgements and
// buffering all recieved packets.
func (rr *ReliableReceiver) runListen() {
	for {
		bytes := rr.selector.Receive()
		packet := data.DataPacketFromBytes(bytes)
		ackPacket := createAck(packet)
		rr.selector.GetOption().Conn.Send(ackPacket.ToBytes())
		rr.bufferPacket(packet)
	}
}

// bufferPacket will store a packet in the buffer, ensuring packets are
// returned in the proper order, relative to their sequence numbers.
func (rr *ReliableReceiver) bufferPacket(packet *data.DataPacket) {
	rr.bufferLock.Lock()
	defer rr.bufferLock.Unlock()
	if packet.Header.SequenceNumber == rr.nextSequenceNum {
		rr.receivedChan <- packet
		rr.nextSequenceNum++
		nextPacket, exists := rr.buffer[rr.nextSequenceNum]
		for exists {
			rr.receivedChan <- nextPacket
			delete(rr.buffer, rr.nextSequenceNum)
			rr.nextSequenceNum++
			nextPacket, exists = rr.buffer[rr.nextSequenceNum]
		}
	} else {
		rr.buffer[packet.Header.SequenceNumber] = packet
	}
}

// createAck will create a data packet that serves as an acknowledgement of the
// for the received packet.
func createAck(packet *data.DataPacket) *data.DataPacket {
	return &data.DataPacket{
		Header: data.PacketHeader{
			SourceAddress:      packet.Header.DestinationAddress,
			DestinationAddress: packet.Header.SourceAddress,
			PreviousHop:        network.GetMyAddress(),
			PacketType:         data.PacketTypeAck,
			SequenceNumber:     packet.Header.SequenceNumber,
			TTL:                data.MaxTTL,
			SendKey:            packet.Header.SendKey | 0x8000,
		},
		Body: []data.SensorData{},
	}
}
