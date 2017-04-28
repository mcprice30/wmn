package network

import (
  "fmt"
  "time"

  "github.com/mcprice30/wmn/data"
)

var neighborTable *data.NeighborTable = data.CreateNeighborTable()

func (c *ManetConnection) HelloLoop() {

  t := time.NewTicker(10 * time.Second)


  for {
    c.SendHelloPacket()
    <-t.C
  }
}

func (c *ManetConnection) SendHelloPacket() {
  outgoingHeader := data.PacketHeader{
    SourceAddress:      GetMyAddress(),
    DestinationAddress: data.BroadcastAddress,
    PreviousHop:        GetMyAddress(),
    TTL:                data.MaxTTL,
    PacketType:         data.PacketTypeHello,
  }

  unidirectionalLinks := []data.ManetAddr{}
  bidirectionalLinks := []data.ManetAddr{}
  mprLinks := []data.ManetAddr{}

  for neighbor, row := range neighborTable.Table {
    if row.LinkType == data.LinkTypeUnidirectional {
      unidirectionalLinks = append(unidirectionalLinks, neighbor)
    } else if row.LinkType == data.LinkTypeBidirectional {
      bidirectionalLinks = append(bidirectionalLinks, neighbor)
    } else if row.LinkType == data.LinkTypeMPR {
      mprLinks = append(mprLinks, neighbor)
    }
  }

  sendPacket := &data.HelloPacket {
    Header: outgoingHeader,
    NumBidirectional: uint8(len(bidirectionalLinks)),
    BidirectionalLinks: bidirectionalLinks,
    NumHeard: uint8(len(unidirectionalLinks)),
    HeardLinks: unidirectionalLinks,
    NumMPR: uint8(len(mprLinks)),
    MPRLinks: mprLinks,
  }

  fmt.Println("Sending Hello Packet", sendPacket)
  c.Send(sendPacket.ToBytes())
}

func handleHelloPacket(packet []byte) {
  fmt.Printf("Got Hello Packet at 0x%04x\n", GetMyAddress())
  hp := data.HelloPacketFromBytes(packet)

  bidirectionalConnection := false
  isMPRSelector := false
  twoHopNeighbors := []data.ManetAddr{}
  for _, neighbor := range hp.HeardLinks {
    if neighbor == GetMyAddress() {
      bidirectionalConnection = true
    }
    twoHopNeighbors = append(twoHopNeighbors, neighbor)
  }

  for _, neighbor := range hp.BidirectionalLinks {
    if neighbor == GetMyAddress() {
      bidirectionalConnection = true
    }
    twoHopNeighbors = append(twoHopNeighbors, neighbor)
  }

  for _, neighbor := range hp.MPRLinks {
    if neighbor == GetMyAddress() {
      bidirectionalConnection = true
      isMPRSelector = true
    }
    twoHopNeighbors = append(twoHopNeighbors, neighbor)
  }

  neighborTable.Selectors[hp.Header.SourceAddress] = isMPRSelector
  if bidirectionalConnection {
    neighborTable.Table[hp.Header.SourceAddress] = &data.NeighborTableRow {
      LinkType: data.LinkTypeBidirectional,
      TwoHopNeigbhors: twoHopNeighbors,
    }
  } else {
    neighborTable.Table[hp.Header.SourceAddress] = &data.NeighborTableRow {
      LinkType: data.LinkTypeUnidirectional,
      TwoHopNeigbhors: twoHopNeighbors,
    }
  }

  findMPRs()
}

func findMPRs() {

  for _, row := range neighborTable.Table {
    if row.LinkType == data.LinkTypeMPR {
      row.LinkType = data.LinkTypeBidirectional
    }
  }

  twoHopNeighbors := map[data.ManetAddr]bool{}

  for neighbor, row := range neighborTable.Table {
    twoHopNeighbors[neighbor] = true
    for _, twoHop := range row.TwoHopNeigbhors {
      twoHopNeighbors[twoHop] = true
    }
  }

  for neighbor, row := range neighborTable.Table {
    if row.LinkType == data.LinkTypeBidirectional {
      delete(twoHopNeighbors, neighbor)
      for _, twoHop := range row.TwoHopNeigbhors {
        delete(twoHopNeighbors, twoHop)
      }
      row.LinkType = data.LinkTypeMPR
      if len(twoHopNeighbors) == 0 {
        break
      }
    }
  }
}
