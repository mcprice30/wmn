package transport

import (
	"fmt"
	"sync"

	"github.com/mcprice30/wmn/data"
	"github.com/mcprice30/wmn/network"
)

const maxBufferSize = 64

type ReliableReceiver struct {
	conn            network.Connection
	buffer          map[uint16]*data.DataPacket
	bufferLock      *sync.Mutex
	nextSequenceNum uint16
	receivedChan    chan *data.DataPacket
}

func CreateReliableReceiver(address data.ManetAddr) *ReliableReceiver {
	// SWITCHED TO MANET
	conn := network.BindManet()
	conn.SetNeighbors([]data.ManetAddr{0x0003})
	// END SWITCHED TO MANET
	out := &ReliableReceiver{
		conn:            conn,
		buffer:          map[uint16]*data.DataPacket{},
		bufferLock:      &sync.Mutex{},
		nextSequenceNum: 0,
		receivedChan:    make(chan *data.DataPacket, maxBufferSize),
	}
	go out.runListen()
	return out
}

func (rr *ReliableReceiver) Listen() *data.DataPacket {
	return <-rr.receivedChan
}

func (rr *ReliableReceiver) Close() {
	rr.conn.Close()
}

func (rr *ReliableReceiver) runListen() {
	for {
		bytes := rr.conn.Receive()
		packet := data.DataPacketFromBytes(bytes)
		ackPacket := createAck(packet)
		rr.conn.Send(ackPacket.ToBytes())
		rr.bufferPacket(packet)
	}
}

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

func createAck(packet *data.DataPacket) *data.DataPacket {
	return &data.DataPacket{
		Header: data.PacketHeader{
			SourceAddress:      packet.Header.DestinationAddress,
			DestinationAddress: packet.Header.SourceAddress,
			PreviousHop:        network.GetMyAddress(),
			PacketType:         data.PacketTypeAck,
			SequenceNumber:     packet.Header.SequenceNumber,
			TTL:                data.MaxTTL,
			SendKey:            packet.Header.SendKey ^ 0x8000,
		},
		Body: []data.SensorData{},
	}
}

func RunEcho(address data.ManetAddr) {
	rec := CreateReliableReceiver(address)

	defer rec.Close()

	for {
		data := rec.Listen()
		fmt.Println("GOT: ", data.Header.SequenceNumber)
	}
}
