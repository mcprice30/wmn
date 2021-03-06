package transport

import (
	"fmt"
	"sync"
	"time"

	"github.com/mcprice30/wmn/data"
	"github.com/mcprice30/wmn/network"
)

// resendDelay is the delay between resending packets.
const resendDelay = "25ms"

// ReliableSender will send packets with to a ReliableReceiver. It will continue
// sending each packet until it receives an acknowledgement.
type ReliableSender struct {
	selector           *Selector
	seqNum             uint16
	outstandingPackets map[uint16]bool
	bufferLock         *sync.Mutex
	interval           time.Duration
}

// CreateReliableSender will create a new instance of a reliable sender
// that sends data from the given manet address.
func CreateReliableSender() *ReliableSender {
	selector := CreateSelector()
	duration, err := time.ParseDuration(resendDelay)
	if err != nil {
		panic(err)
	}
	out := &ReliableSender{
		selector:           selector,
		outstandingPackets: map[uint16]bool{},
		interval:           duration,
		bufferLock:         &sync.Mutex{},
	}
	go out.listenForAck()
	return out
}

// Transmit will reliably transmit a data packet across the network.
func (rc *ReliableSender) Transmit(packet *data.DataPacket) {
	packet.Header.SourceAddress = network.GetMyAddress()
	packet.Header.SequenceNumber = rc.seqNum
	packet.Header.TTL = data.MaxTTL
	packet.Header.PreviousHop = network.GetMyAddress()
	go rc.sendBytes(packet.ToBytes(), rc.seqNum)
	rc.seqNum++
}

// listenForAck will run in the background and listen for acknowledgments for
// all outstanding packets.
func (rc *ReliableSender) listenForAck() {
	for {
		ack := data.DataPacketFromBytes(rc.selector.Receive())
		fmt.Println("Got ack for", ack.Header.SequenceNumber)
		rc.bufferLock.Lock()
		delete(rc.outstandingPackets, ack.Header.SequenceNumber)
		rc.bufferLock.Unlock()
	}
}

// sendBytes will repeatedly transmit the given bytes, until an acknowledgement
// is received for the given sequence number.
func (rc *ReliableSender) sendBytes(bytes []byte, seqNum uint16) {
	rc.bufferLock.Lock()
	rc.outstandingPackets[seqNum] = true
	rc.bufferLock.Unlock()
	fmt.Printf("Transmitting packet #%d\n", seqNum)
	t := time.NewTicker(rc.interval)
	conn := rc.selector.GetOption(true)
	conn.Conn.Send(bytes)
	for {
		<-t.C
		rc.bufferLock.Lock()
		if _, outstanding := rc.outstandingPackets[seqNum]; outstanding {
			rc.bufferLock.Unlock()
			header := data.PacketHeaderFromBytes(bytes)
			header.SendKey = header.SendKey + 1
			for i, b := range header.ToBytes() {
				bytes[i] = b
			}
			rc.selector.Failed(conn)
			conn = rc.selector.GetOption(true)
			conn.Conn.Send(bytes)
		} else {
			rc.selector.Succeeded(conn)
			rc.bufferLock.Unlock()
			t.Stop()
			return
		}
	}
}
