// Package chief contains logic for running the fire chief's end of the network.
package chief

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"

	"github.com/mcprice30/wmn/data"
	"github.com/mcprice30/wmn/transport"
)

var sendChan = make(chan []byte, 8)

// RunListener will listen for incoming packets on our local address.
// It will print out the buffers of the packets recieved.
func RunListener(port string) {

	go runServer(port)

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
		for i, b := range buffers {
			sendData := map[string]interface{}{
				"id":   i,
				"data": b.GetData(),
			}
			if js, err := json.Marshal(sendData); err != nil {
				fmt.Println("ERROR: ", err)
			} else {
				sendChan <- js
			}
		}
	}
}

func wsHandler(ws *websocket.Conn) {
	for {
		if _, err := ws.Write(<-sendChan); err != nil {
			fmt.Println(err)
			break
		}
	}
}

func runServer(port string) {
	prt := ":" + port
	http.Handle("/ws", websocket.Handler(wsHandler))
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/home", http.StripPrefix("/home", fs))
	err := http.ListenAndServe(prt, nil)
	if err != nil {
		fmt.Println("WS ERR: ", err)
	}
}
