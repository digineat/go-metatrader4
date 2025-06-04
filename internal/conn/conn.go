package conn

import (
	"context"
	"io"
	"net"
	"time"
)

// Conn is a wrapper around net.Conn with convenience helpers.
type Conn struct {
	netConn net.Conn
}

// FromNetConn wraps an existing net.Conn. Useful for tests.
func FromNetConn(n net.Conn) *Conn { return &Conn{netConn: n} }

// Dial opens a TCP connection to addr using the given timeout.
func Dial(ctx context.Context, addr string, timeout time.Duration) (*Conn, error) {
	d := net.Dialer{Timeout: timeout}
	c, err := d.DialContext(ctx, "tcp", addr)
	if err != nil {
		return nil, err
	}
	return &Conn{netConn: c}, nil
}

// Close closes the underlying network connection.
func (c *Conn) Close() error {
	if c.netConn == nil {
		return nil
	}
	return c.netConn.Close()
}

// Send writes data to the connection with the provided timeout.
func (c *Conn) Send(ctx context.Context, data []byte, timeout time.Duration) error {
	if dl, ok := ctx.Deadline(); ok {
		if err := c.netConn.SetWriteDeadline(dl); err != nil {
			return err
		}
	} else {
		if err := c.netConn.SetWriteDeadline(time.Now().Add(timeout)); err != nil {
			return err
		}
	}
	_, err := c.netConn.Write(data)
	return err
}

// Receive reads all data from the connection with the provided timeout.
func (c *Conn) Receive(ctx context.Context, timeout time.Duration) ([]byte, error) {
	if dl, ok := ctx.Deadline(); ok {
		if err := c.netConn.SetReadDeadline(dl); err != nil {
			return nil, err
		}
	} else {
		if err := c.netConn.SetReadDeadline(time.Now().Add(timeout)); err != nil {
			return nil, err
		}
	}
	return io.ReadAll(c.netConn)
}
