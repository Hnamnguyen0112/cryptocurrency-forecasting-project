package websocket

import (
	"net/http"
	"net/http/httptest"
	"os"
	"os/signal"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockConn struct {
	mock.Mock
}

func (m *MockConn) WriteJSON(v interface{}) error {
	args := m.Called(v)
	return args.Error(0)
}

func (m *MockConn) WriteMessage(messageType int, data []byte) error {
	args := m.Called(messageType, data)
	return args.Error(0)
}

func (m *MockConn) Close() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockConn) ReadMessage() (messageType int, p []byte, err error) {
	args := m.Called()
	return args.Int(0), args.Get(1).([]byte), args.Error(2)
}

var upgrader = websocket.Upgrader{}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			break
		}
		err = c.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}

func TestConnect(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(echo))
	defer s.Close()

	u := "ws" + strings.TrimPrefix(s.URL, "http")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	ws := Connect(u, interrupt)

	assert.NotNil(t, ws.Conn)
	assert.Equal(t, interrupt, ws.Interrupt)
	assert.Equal(t, 0, ws.TotalChannels)
}

func TestSubscribe(t *testing.T) {
	mockConn := new(MockConn)
	ws := &Websocket{
		Conn:          mockConn,
		Done:          make(chan struct{}),
		Interrupt:     make(chan os.Signal, 1),
		TotalChannels: 0,
	}

	request := map[string]interface{}{"channel": "test"}
	mockConn.On("WriteJSON", request).Return(nil)

	ws.Subscribe(request)

	mockConn.AssertExpectations(t)
}

func TestHandleInterrupt(t *testing.T) {
	mockConn := new(MockConn)
	ws := &Websocket{
		Conn:          mockConn,
		Done:          make(chan struct{}),
		Interrupt:     make(chan os.Signal, 1),
		TotalChannels: 0,
	}

	mockConn.On("WriteMessage", websocket.CloseMessage, mock.Anything).Return(nil)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ws.Interrupt <- os.Interrupt

		close(ws.Done)
	}()

	ws.HandleInterrupt()

	mockConn.AssertExpectations(t)
}

func TestClose(t *testing.T) {
	mockConn := new(MockConn)
	ws := &Websocket{
		Conn:          mockConn,
		Done:          make(chan struct{}),
		Interrupt:     make(chan os.Signal, 1),
		TotalChannels: 0,
	}

	mockConn.On("Close").Return(nil)

	ws.Close()

	mockConn.AssertExpectations(t)
}
