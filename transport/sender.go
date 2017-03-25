package transport

import (
	"fmt"
	"sync"
	"time"

	"github.com/mcprice30/wmn/data"
	"github.com/mcprice30/wmn/network"
)

const resendDelay = "25ms"

type ReliableSender struct {
	conn               *network.SenderConnection
	seqNum             uint16
	outstandingPackets map[uint16]bool
	bufferLock         *sync.Mutex
	interval           time.Duration
}

func CreateReliableSender(src, dst data.ManetAddr) *ReliableSender {
	conn := network.Connect(src, dst)
	duration, err := time.ParseDuration(resendDelay)
	if err != nil {
		panic(err)
	}
	out := &ReliableSender{
		conn:               conn,
		outstandingPackets: map[uint16]bool{},
		interval:           duration,
		bufferLock:         &sync.Mutex{},
	}
	go out.listenForAck()
	return out
}

func (rc *ReliableSender) Transmit(packet *data.DataPacket) {
	packet.Header.SequenceNumber = rc.seqNum
	packet.Header.TTL = data.MaxTTL
	go rc.sendBytes(packet.ToBytes(), rc.seqNum)
	rc.seqNum++
}

func (rc *ReliableSender) listenForAck() {
	for {
		ack := data.DataPacketFromBytes(rc.conn.Receive())
		fmt.Println("Got ack for", ack.Header.SequenceNumber)
		rc.bufferLock.Lock()
		delete(rc.outstandingPackets, ack.Header.SequenceNumber)
		rc.bufferLock.Unlock()
	}
}

func (rc *ReliableSender) sendBytes(data []byte, seqNum uint16) {
	rc.bufferLock.Lock()
	rc.outstandingPackets[seqNum] = true
	rc.bufferLock.Unlock()
	fmt.Printf("Transmitting packet #%d\n", seqNum)
	t := time.NewTicker(rc.interval)
	rc.conn.Send(data)
	for {
		<-t.C
		rc.bufferLock.Lock()
		if _, outstanding := rc.outstandingPackets[seqNum]; outstanding {
			rc.bufferLock.Unlock()
			fmt.Printf("Re-transmitting packet #%d\n", seqNum)
			rc.conn.Send(data)
		} else {
			rc.bufferLock.Unlock()
			t.Stop()
			return
		}
	}
}
