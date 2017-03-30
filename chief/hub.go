// Package chief contains logic for running the fire chief's end of the network.
package chief

import (
	"fmt"

	"github.com/mcprice30/wmn/data"
	"github.com/mcprice30/wmn/transport"
)

// RunListener will listen for incoming packets on our local address.
// It will print out the buffers of the packets recieved.
func RunListener() {
	rec := transport.CreateReliableReceiver()
	defer rec.Close()

	buffers := make([]*data.Buffer, data.NumSensorTypes)
	for i := range buffers {
		buffers[i] = data.CreateBuffer()
	}

	for {
		data := rec.Listen()
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
