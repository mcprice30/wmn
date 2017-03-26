package network

type Connection interface {
	Send(bytes []byte)
	Receive() []byte
	Close()
}
