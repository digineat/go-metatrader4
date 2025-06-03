package conn

import (
	"context"
	"io"
	"net"
	"time"
)

type Conn struct {
	netConn net.Conn
}

// FromNetConn wraps an existing net.Conn. Useful for tests.
func FromNetConn(n net.Conn) *Conn { return &Conn{netConn: n} }

func Dial(ctx context.Context, addr string, timeout time.Duration) (*Conn, error) {
	d := net.Dialer{Timeout: timeout}
	c, err := d.DialContext(ctx, "tcp", addr)
	if err != nil {
		return nil, err
	}
	return &Conn{netConn: c}, nil
}

func (c *Conn) Close() error {
	if c.netConn == nil {
		return nil
	}
	return c.netConn.Close()
}

func (c *Conn) Send(ctx context.Context, data []byte, timeout time.Duration) error {
	if dl, ok := ctx.Deadline(); ok {
		c.netConn.SetWriteDeadline(dl)
	} else {
		c.netConn.SetWriteDeadline(time.Now().Add(timeout))
	}
	_, err := c.netConn.Write(data)
	return err
}

func (c *Conn) Receive(ctx context.Context, timeout time.Duration) ([]byte, error) {
	if dl, ok := ctx.Deadline(); ok {
		c.netConn.SetReadDeadline(dl)
	} else {
		c.netConn.SetReadDeadline(time.Now().Add(timeout))
	}
	return io.ReadAll(c.netConn)
}
