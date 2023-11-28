package websocket

import "github.com/gorilla/websocket"

// Message websocket消息
type Message struct {
	t int
	v []byte
}

func (m *Message) T() int {
	return m.t
}

func (m *Message) V() []byte {
	return m.v
}

// NewMessage 返回一个websocket消息
func NewMessage(t int, v []byte) *Message {
	return &Message{
		t: t,
		v: v,
	}
}

// NewTextMsg 返回一个websocket.TextMessage
func NewTextMsg(s string) *Message {
	return &Message{
		t: websocket.TextMessage,
		v: []byte(s),
	}
}

// NewBinaryMsg 返回一个websocket.BinaryMessage
func NewBinaryMsg(v []byte) *Message {
	return &Message{
		t: websocket.BinaryMessage,
		v: v,
	}
}
