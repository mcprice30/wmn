// Package chief contains logic for running the fire chief's end of the network.
package chief

import (
	"fmt"

	"github.com/mcprice30/wmn/data"
	"github.com/mcprice30/wmn/transport"
)

// RunListener will listen for incoming packets on the incoming manet address.
// It will print out the buffers of the packets recieved.
func RunListener(address data.ManetAddr) {
	rec := transport.CreateReliableReceiver(address)
	defer rec.Close()

	buffers := make([]*data.Buffer, data.NumSensorTypes)
	for i := range buffers {
		buffers[i] = data.CreateBuffer()
	}

	for {
		data := rec.Listen()
		fmt.Println("GOT: ", data.Header.SequenceNumber)
		for _, d := range data.Body {
			buffers[d.Type()].Add(d)
		}
		for _, b := range buffers {
			for _, d := range b.GetData() {
				fmt.Print(d.Id(), " ")
			}
			fmt.Println()
		}
	}

}
