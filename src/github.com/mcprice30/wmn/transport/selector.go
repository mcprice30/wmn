package transport

import (
	"fmt"
	"math/rand"
	"net"

	"github.com/mcprice30/wmn/network"
)

const numConnectionOptions = 2
const successMemory = 32

// Selector is used to automatically delegate the the best network for
// transmission.
type Selector struct {
	connections  []*ConnectionOption
	successful   [][]bool
	successRates []int
	idx          []int
	pastManet    int
	pastEthernet int
}

// CreateSelector will instantiate and return a new selector.
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

// GetOption will randomly select an available connection option, favoring
// the network that has been more reliable recently.
func (s *Selector) GetOption(disp bool) *ConnectionOption {

	if s.pastManet+s.pastEthernet >= 100 && disp {
		fmt.Println("Past 100 transmissions:")
		fmt.Printf("%d%% Manet\n", s.pastManet)
		fmt.Printf("%d%% Ethernet\n", s.pastEthernet)
		s.pastManet = 0
		s.pastEthernet = 0
	}

	r := rand.Int31n(2*successMemory + 2)
	g := s.successRates[0] - s.successRates[1] + successMemory + 1
	if int(r) < g {
		s.pastManet++
		return s.connections[0]
	} else {
		s.pastEthernet++
		return s.connections[1]
	}
}

// Recieve will receive and return bytes from either connection.
func (s *Selector) Receive() []byte {
	return s.connections[0].Conn.Receive()
}

// Succeeded should be called when a transmission over the given connection was
// successful.
func (s *Selector) Succeeded(c *ConnectionOption) {
	old := s.successRates[c.Id]
	if s.successful[c.Id][s.idx[c.Id]] {
		old--
	}
	s.successRates[c.Id] = old + 1
	s.successful[c.Id][s.idx[c.Id]] = true
	s.idx[c.Id] = (s.idx[c.Id] + 1) % successMemory
}

// Failed should be called when a transmission over the given connection was
// unsuccessful.
func (s *Selector) Failed(c *ConnectionOption) {
	old := s.successRates[c.Id]
	if s.successful[c.Id][s.idx[c.Id]] {
		old--
	}
	s.successRates[c.Id] = old
	s.successful[c.Id][s.idx[c.Id]] = false
	s.idx[c.Id] = (s.idx[c.Id] + 1) % successMemory
}

// Close will close all potential connections in the selector.
func (s *Selector) Close() {
	for _, conn := range s.connections {
		conn.Conn.Close()
	}
}

// ConnectionOption represents a pairing of a connection to send across with
// its associated id.
type ConnectionOption struct {
	Conn network.Connection
	Id   int
}
