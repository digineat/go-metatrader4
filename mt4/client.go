package mt4

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/valentineus/go-metatrader4/internal/conn"
	"github.com/valentineus/go-metatrader4/internal/proto"
)

// Client provides access to a MetaTrader4 server.
type Client struct {
	addr         string
	port         int
	dialTimeout  time.Duration
	readTimeout  time.Duration
	writeTimeout time.Duration
	autoClose    bool
	dialer       net.Dialer
	c            *conn.Conn
}

// Option configures the Client.
type Option func(*Client)

// WithDialTimeout sets timeout for establishing connections.
func WithDialTimeout(d time.Duration) Option { return func(c *Client) { c.dialTimeout = d } }

// WithReadTimeout sets timeout for reading responses.
func WithReadTimeout(d time.Duration) Option { return func(c *Client) { c.readTimeout = d } }

// WithWriteTimeout sets timeout for writing requests.
func WithWriteTimeout(d time.Duration) Option { return func(c *Client) { c.writeTimeout = d } }

// WithAutoClose enables or disables automatic connection close after Execute.
func WithAutoClose(b bool) Option { return func(c *Client) { c.autoClose = b } }

// NewClient creates a new Client with optional configuration.
func NewClient(addr string, port int, opts ...Option) *Client {
	cl := &Client{
		addr:         addr,
		port:         port,
		dialTimeout:  5 * time.Second,
		readTimeout:  5 * time.Second,
		writeTimeout: 5 * time.Second,
		autoClose:    true,
	}
	for _, o := range opts {
		o(cl)
	}
	return cl
}

// Connect establishes connection to the MT4 server if not already connected.
func (c *Client) Connect(ctx context.Context) error {
	if c.c != nil {
		return nil
	}
	address := fmt.Sprintf("%s:%d", c.addr, c.port)
	cn, err := conn.Dial(ctx, address, c.dialTimeout)
	if err != nil {
		return err
	}
	c.c = cn
	return nil
}

// Close closes underlying connection.
func (c *Client) Close() error {
	if c.c == nil {
		return nil
	}
	err := c.c.Close()
	c.c = nil
	return err
}

// Execute sends command with params to the server and returns decoded response.
func (c *Client) Execute(ctx context.Context, command string, params map[string]string) (string, error) {
	if err := c.Connect(ctx); err != nil {
		return "", fmt.Errorf("connect: %w", err)
	}

	encoded, err := proto.EncodeParams(params)
	if err != nil {
		if c.autoClose {
			c.Close()
		}
		return "", err
	}
	req := proto.BuildRequest(command, encoded, c.autoClose)

	if err := c.c.Send(ctx, req, c.writeTimeout); err != nil {
		if c.autoClose {
			c.Close()
		}
		return "", fmt.Errorf("send: %w", err)
	}

	respBytes, err := c.c.Receive(ctx, c.readTimeout)
	if c.autoClose {
		c.Close()
	}
	if err != nil {
		return "", fmt.Errorf("receive: %w", err)
	}

	resp, err := proto.DecodeResponse(string(respBytes))
	if err != nil {
		return "", err
	}
	return resp, nil
}
