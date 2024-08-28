package websocket

import (
	"github.com/gorilla/websocket"
	"log"
	"os"
)

type WebsocketConn interface {
	WriteJSON(v interface{}) error
	WriteMessage(messageType int, data []byte) error
	ReadMessage() (messageType int, p []byte, err error)
	Close() error
}

type Websocket struct {
	Conn          WebsocketConn
	Done          chan struct{}
	Interrupt     chan os.Signal
	TotalChannels int
}

func Connect(u string, interrupt chan os.Signal) *Websocket {
	c, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		log.Fatal("Dial error:", err)
		os.Exit(1)
	}

	return &Websocket{
		Conn:          c,
		Done:          make(chan struct{}),
		Interrupt:     interrupt,
		TotalChannels: 0,
	}
}

func (ws *Websocket) Subscribe(request map[string]interface{}) {
	err := ws.Conn.WriteJSON(request)
	if err != nil {
		log.Fatal("Subscribe error:", err)
	}
}

func (ws *Websocket) HandleInterrupt() {
	for {
		select {
		case <-ws.Done:
			return
		case <-ws.Interrupt:
			log.Println("Interrupt received, websocket closing connection ...")

			// Cleanly close the connection
			err := ws.Conn.WriteMessage(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
			)
			if err != nil {
				log.Println("Close message error:", err)
				return
			}
			select {
			case <-ws.Done:
			}
			return
		}
	}
}

func (ws *Websocket) Close() {
	log.Println("Closing websocket connection ...")
	ws.Conn.Close()
	log.Println("Websocket connection closed")
}
