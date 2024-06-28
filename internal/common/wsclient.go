package common

import (
	"bufio"
	"github.com/gorilla/websocket"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	PingMsg    = "ping"
	PingPeriod = 3 * time.Second
)

var upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WsClient struct {
	client           *websocket.Conn
	message          chan wsMessage
	stopCh           chan struct{}
	checkFailedCount int
}

type wsMessage struct {
	data    []byte
	msgType int
}

func NewWsClient(client *websocket.Conn) *WsClient {
	wsClient := &WsClient{
		client:           client,
		message:          make(chan wsMessage, 5),
		stopCh:           make(chan struct{}),
		checkFailedCount: 0,
	}

	client.SetCloseHandler(func(code int, text string) error {
		//println("websocket: user client send close event")
		wsClient.Close()
		return nil
	})

	return wsClient
}

func (ws *WsClient) Close() {
	close(ws.stopCh)
	close(ws.message)
}

func (ws *WsClient) HandleLogs(reader io.Reader) {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()

	ws.readMessage()
	ws.writeMessage()

	go func() {
		defer func() {
			if err := recover(); err != nil {
				return
			}
		}()

		ticker := time.NewTicker(PingPeriod)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				ws.message <- wsMessage{
					data:    []byte(PingMsg),
					msgType: websocket.TextMessage,
				}
			case <-ws.stopCh:
				return
			}
		}
	}()

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		select {
		case <-ws.stopCh:
			return
		default:
			del003EStr := strings.ReplaceAll(scanner.Text(), "\\u003e", ">")
			delN := strings.ReplaceAll(del003EStr, "\\n", "")
			ws.message <- wsMessage{
				data:    []byte(delN),
				msgType: websocket.TextMessage,
			}
		}
	}
	close(ws.message)
}

func (ws *WsClient) writeMessage() {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				return
			}
		}()

		for {
			select {
			case msg, ok := <-ws.message:
				if !ok && string(msg.data) == "" {
					ws.Close()
					return
				}
				if err := ws.client.WriteMessage(msg.msgType, msg.data); err != nil {
					return
				}
				if string(msg.data) == PingMsg {
					_ = ws.client.SetReadDeadline(time.Now().Add(2*PingPeriod + time.Second))
				}
			case <-ws.stopCh:
				return
			}
		}
	}()
}

func (ws *WsClient) readMessage() {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				return
			}
		}()
		for {
			select {
			case <-ws.stopCh:
				return
			default:
				if _, _, err := ws.client.ReadMessage(); err != nil {
					if ws.checkFailedCount > 10 {
						ws.Close()
						break
					}
					ws.checkFailedCount++
					time.Sleep(300 * time.Millisecond)
				}
				if ws.checkFailedCount != 0 {
					ws.checkFailedCount = 0
				}
			}
		}
	}()
}
