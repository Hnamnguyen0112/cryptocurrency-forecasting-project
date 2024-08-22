package main

import (
	"log"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector/config"
	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector/pkg/database"
	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector/pkg/websocket"
)

const idleTimeout = 5 * time.Second

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	dbUser := config.Config("COINBASE_DB_USER")
	dbPassword := config.Config("COINBASE_DB_PASSWORD")
	dbHost := config.Config("COINBASE_DB_HOST")
	dbPort := config.Config("COINBASE_DB_PORT")
	dbName := config.Config("COINBASE_DB_NAME")

	connectParams := database.ConnectParams{
		Host:     dbHost,
		Port:     dbPort,
		User:     dbUser,
		Password: dbPassword,
		Name:     dbName,
	}

	database.Connect(connectParams)

	u := url.URL{
		Scheme: "wss",
		Host:   "advanced-trade-ws.coinbase.com",
		Path:   "/",
	}
	log.Printf("Connecting to %s", u.String())

	ws := websocket.Connect(u.String(), interrupt)
	defer ws.Conn.Close()

	request := map[string]interface{}{
		"type":        "subscribe",
		"product_ids": []string{"BTC-USDT"},
		"channel":     "ticker",
	}

	err := ws.Conn.WriteJSON(request)
	if err != nil {
		log.Fatal("Write error:", err)
	}

	go func() {
		defer close(ws.Done)
		ws.Listen()
	}()

	ws.HandleInterrupt()
}
