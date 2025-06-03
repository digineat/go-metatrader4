package mt4

import (
	"bufio"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net"
	"strings"
	"time"
	"unicode"

	"golang.org/x/text/encoding/charmap"
)

// Client provides access to a MetaTrader4 server.
type Client struct {
	addr   string
	port   int
	dialer net.Dialer
}

// NewClient creates a new Client.
func NewClient(addr string, port int) *Client {
	return &Client{
		addr:   addr,
		port:   port,
		dialer: net.Dialer{},
	}
}

// Execute sends a raw command with parameters to the MT4 server.
// params will be encoded using Windows-1251 charset and base64.
func (c *Client) Execute(ctx context.Context, command string, params map[string]string) (string, error) {
	var sb strings.Builder
	for k, v := range params {
		if sb.Len() > 0 {
			sb.WriteByte('|')
		}
		sb.WriteString(k)
		sb.WriteByte('=')
		sb.WriteString(v)
	}
	sb.WriteByte('|')

	enc := charmap.Windows1251.NewEncoder()
	encodedParams, err := enc.String(sb.String())
	if err != nil {
		return "", fmt.Errorf("encode params: %w", err)
	}

	payload := command + " " + base64.StdEncoding.EncodeToString([]byte(encodedParams)) + "\nQUIT\n"

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	conn, err := c.dialer.DialContext(ctx, "tcp", fmt.Sprintf("%s:%d", c.addr, c.port))
	if err != nil {
		return "", fmt.Errorf("dial server: %w", err)
	}
	defer conn.Close()

	if deadline, ok := ctx.Deadline(); ok {
		conn.SetDeadline(deadline)
	}

	if _, err := fmt.Fprint(conn, payload); err != nil {
		return "", fmt.Errorf("write request: %w", err)
	}

	res, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil && err != io.EOF {
		return "", fmt.Errorf("read response: %w", err)
	}

	dec := charmap.Windows1251.NewDecoder()
	decodedStr, err := dec.String(res)
	if err != nil {
		return "", fmt.Errorf("decode response charset: %w", err)
	}

	bytesRes, err := base64.StdEncoding.DecodeString(decodedStr)
	if err != nil {
		return "", fmt.Errorf("base64 decode: %w", err)
	}

	cleaned := strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return -1
	}, string(bytesRes))

	return cleaned, nil
}
