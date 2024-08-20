package websocket

import (
	"log"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

type Websocket struct {
  Conn *websocket.Conn
  Done chan struct{}
  Interrupt chan os.Signal
}

func Connect(u string) *Websocket {
  interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

  c, _, err := websocket.DefaultDialer.Dial(u, nil)
  if err != nil {
    log.Fatal("Dial error:", err)
  }
  
  return &Websocket{
    Conn: c,
    Done: make(chan struct{}),
    Interrupt: interrupt,
  }
}

func (ws *Websocket) Listen() {
  for {
    _, message, err := ws.Conn.ReadMessage()
  if err != nil {
      log.Println("Read error:", err)
      return
    }
    log.Printf("Received: %s", message)
  }
}

func (ws *Websocket) HandleInterrupt() {
  for {
    select {
    case <-ws.Done:
      return
    case <-ws.Interrupt:
      log.Println("Interrupt received, closing connection")

      // Cleanly close the connection
      err := ws.Conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
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
