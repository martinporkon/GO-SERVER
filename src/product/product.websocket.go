package product

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/net/websocket"
)

type Message struct {
	Data string `json:"Data"`
	Type string `json:"type"`
}

func productSocket(ws *websocket.Conn) {
	// listen or inoming data
	// as we dont want to block, we will put that into a gorutine

	done := make(chan struct{})
	fmt.Println("new websocket connection established")
	go func(c *websocket.Conn) {
		for {
			var msg Message
			if err := websocket.JSON.Receive(ws, &msg); err != nil {
				log.Println(err)
				break
			}
			fmt.Printf("received message %s\n", msg.Data)
		} // if thes would be an event system. Typically a queue would be used to listen for product data.
		close(done)
	}(ws)
	for {
		select {
		case <-done:
			fmt.Println("connection was closed, lets break out of here")
		default:
			for {
				products, err := GetTopTenProducts()
				if err != nil {
					log.Println(err)
					break
				}
				if err := websocket.JSON.Send(ws, products); err != nil {
					log.Println(err)
					break
				}
				time.Sleep(10 * time.Second)
			}
		}
	}
	fmt.Println("closing the websocket")
	defer ws.Close()
}

// JavaScript in the Chrome Browser Developer tools
//> let ws = new WebSocket("ws://localhost:5000/websocket")
//<. undefined
//> ws.send(JSON.stringify({data: "test message from browser", type: "test"}))
// This will be now send to the web service
// the websockets can be easily tested in a browser
