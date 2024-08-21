package main

import (
	"log"
	"net/url"
	"time"

	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector/config"
	binanceworker "github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector/internal/binance_worker"
	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector/pkg/database"
	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector/pkg/entities"
	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector/pkg/websocket"
)

const idleTimeout = 5 * time.Second

func main() {
  dbUser := config.Config("BINANCE_DB_USER")
	dbPassword := config.Config("BINANCE_DB_PASSWORD")
	dbHost := config.Config("BINANCE_DB_HOST")
	dbPort := config.Config("BINANCE_DB_PORT")
	dbName := config.Config("BINANCE_DB_NAME")

  connectParams := database.ConnectParams{
    Host:     dbHost,
    Port:     dbPort,
    User:     dbUser,
    Password: dbPassword,
    Name:   dbName,
  }

  database.Connect(connectParams)

  database.DB.AutoMigrate(&entities.BinanceTicker{},&entities.BinanceCandlestick{})

	u := url.URL{
		Scheme: "wss",
		Host:   "fstream.binance.com",
		Path:   "/ws",
	}
	log.Printf("Connecting to %s", u.String())

	ws := websocket.Connect(u.String()) 
	defer ws.Conn.Close()

  svc := binanceworker.NewService(ws)

  svc.Subcribe("btcusdt", "ticker") 
  svc.Subcribe("btcusdt", "kline_1m")

  go func() { 
    defer close(ws.Done)
    svc.Listen()
  }()

  ws.HandleInterrupt()	
}
