package client

import (
	"github.com/gorilla/websocket"
)

const (
	WsServer = "ws://rchat.gagalive.kr:8082/"
)

type GagaliveClient struct {
	Connection      *websocket.Conn
	Messages        []string
	LastMessageChan chan string
	Disconnected    bool
	Chan            chan bool

	OnDisconnected func()
	OnConnected    func()
	OnMessage      func(string)
}

func (client *GagaliveClient) SetOnMessage(fn func(string)) {
	client.OnMessage = fn
}

func (client *GagaliveClient) SetOnConnected(fn func()) {
	client.OnConnected = fn
}

func (client *GagaliveClient) SetOnDisconnected(fn func()) {
	client.OnDisconnected = fn
}

func (client *GagaliveClient) Send(msg string) error {
	return client.Connection.WriteMessage(websocket.TextMessage, []byte("#"+msg))
}

func (client *GagaliveClient) Connect() error {
	conn, _, err := websocket.DefaultDialer.Dial(WsServer, nil)
	if err != nil {
		return err
	}

	_ = conn.WriteMessage(websocket.TextMessage, []byte("Y"))
	_ = conn.WriteMessage(websocket.TextMessage, []byte("LGuest|@@@randomchat"))

	client.Connection = conn
	client.Disconnected = false
	client.OnConnected()

	go func() {
		client.Read()
	}()

	return nil
}

func (client *GagaliveClient) Read() {
	if client.Connection == nil {
		return
	}

	for {
		_, msg, err := client.Connection.ReadMessage()

		if err != nil {
			client.Disconnected = true
			client.OnDisconnected()
			return
		}

		client.OnMessage(string(msg))
	}
}

func NewGagaClient() GagaliveClient {
	return GagaliveClient{
		Disconnected:    true,
		Chan:            make(chan bool),
		LastMessageChan: make(chan string),
	}
}
