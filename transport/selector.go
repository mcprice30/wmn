package transport

import (
	"math/rand"
	"net"

	"github.com/mcprice30/wmn/network"
)

const numConnectionOptions = 2
const successMemory = 32

type Selector struct {
	connections  []*ConnectionOption
	successful   [][]bool
	successRates []int
	idx          []int
}

func CreateSelector() *Selector {

	addr := network.ToUDPAddr(network.GetMyAddress())
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}

	connections := make([]*ConnectionOption, numConnectionOptions)
	connections[0] = &ConnectionOption{
		Conn: network.BindManet(conn),
		Id:   0,
	}
	connections[1] = &ConnectionOption{
		Conn: network.Bind(conn),
		Id:   1,
	}
	successful := make([][]bool, numConnectionOptions)
	for i := range successful {
		successful[i] = make([]bool, successMemory)
	}

	successRates := make([]int, numConnectionOptions)
	idx := make([]int, numConnectionOptions)

	return &Selector{
		connections:  connections,
		successful:   successful,
		successRates: successRates,
		idx:          idx,
	}

}

func (s *Selector) GetOption() *ConnectionOption {
	r := rand.Int31n(2*successMemory + 2)
	g := s.successRates[0] - s.successRates[1] + successMemory + 1
	if int(r) < g {
		return s.connections[0]
	} else {
		return s.connections[1]
	}
}

func (s *Selector) Receive() []byte {
	return s.connections[0].Conn.Receive()
}

func (s *Selector) Succeeded(c *ConnectionOption) {
	old := s.successRates[c.Id]
	if s.successful[c.Id][s.idx[c.Id]] {
		old--
	}
	s.successRates[c.Id] = old + 1
	s.successful[c.Id][s.idx[c.Id]] = true
	s.idx[c.Id] = (s.idx[c.Id] + 1) % successMemory
}

func (s *Selector) Failed(c *ConnectionOption) {
	old := s.successRates[c.Id]
	if s.successful[c.Id][s.idx[c.Id]] {
		old--
	}
	s.successRates[c.Id] = old
	s.successful[c.Id][s.idx[c.Id]] = false
	s.idx[c.Id] = (s.idx[c.Id] + 1) % successMemory
}

func (s *Selector) Close() {
	for _, conn := range s.connections {
		conn.Conn.Close()
	}
}

type ConnectionOption struct {
	Conn network.Connection
	Id   int
}
