package websocket

import (
	"context"
	"errors"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader *websocket.Upgrader

// SetUpgrader 设置 websocket Upgrader
func SetUpgrader(up *websocket.Upgrader) {
	upgrader = up
}

// Conn websocket连接
type Conn interface {
	// Read 读消息
	Read(ctx context.Context, handler func(ctx context.Context, msg *Message) (*Message, error)) error
	// Write 写消息
	Write(ctx context.Context, msg *Message) error
	// Close 关闭连接
	Close(ctx context.Context) error
}

type conn struct {
	conn   *websocket.Conn
	authOK bool
	authFn func(ctx context.Context, msg *Message) (*Message, error)
}

func (c *conn) Read(ctx context.Context, handler func(ctx context.Context, msg *Message) (*Message, error)) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			t, b, err := c.conn.ReadMessage()
			if err != nil {
				return err
			}

			var msg *Message

			// if `authFunc` is not nil and unauthorized, need to authorize first.
			if c.authFn != nil && !c.authOK {
				msg, err = c.authFn(ctx, NewMessage(t, b))
				if err != nil {
					msg = NewTextMsg(err.Error())
				} else {
					c.authOK = true
				}
			} else {
				if handler != nil {
					msg, err = handler(ctx, NewMessage(t, b))
					if err != nil {
						msg = NewTextMsg(err.Error())
					}
				}
			}

			if msg != nil {
				if err = c.conn.WriteMessage(msg.T(), msg.V()); err != nil {
					return err
				}
			}
		}
	}
}

func (c *conn) Write(ctx context.Context, msg *Message) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	// if `authFn` is not nil and unauthorized, disable to write message.
	if c.authFn != nil && !c.authOK {
		return errors.New("write msg disabled due to unauthorized")
	}

	return c.conn.WriteMessage(msg.T(), msg.V())
}

func (c *conn) Close(ctx context.Context) error {
	return c.conn.Close()
}

// NewConn 生成一个websocket连接
func NewConn(w http.ResponseWriter, r *http.Request, authFn func(ctx context.Context, msg *Message) (*Message, error)) (Conn, error) {
	if upgrader == nil {
		return nil, errors.New("upgrader is nil (forgotten set?)")
	}

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}

	conn := &conn{
		conn:   c,
		authFn: authFn,
	}

	return conn, nil
}
