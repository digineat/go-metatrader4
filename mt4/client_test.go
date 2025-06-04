package mt4

import (
	"context"
	"net"
	"strings"
	"testing"
	"time"

	ic "go.popov.link/metatrader4/internal/conn"
	"go.popov.link/metatrader4/internal/proto"
)

// mockServer returns net.Pipe connections with server writing resp to client.
func mockServer(response string) (net.Conn, net.Conn) {
	server, client := net.Pipe()
	go func() {
		defer server.Close()
		buf := make([]byte, 1024)
		server.Read(buf) // read request ignoring
		server.Write([]byte(response))
	}()
	return client, server
}

func TestClientExecute(t *testing.T) {
	reqParams := map[string]string{"A": "1"}
	encoded, err := proto.EncodeParams(reqParams)
	if err != nil {
		t.Fatalf("encode params: %v", err)
	}
	resp := encoded
	clientConn, _ := mockServer(resp)

	c := &Client{addr: "", port: 0, autoClose: true, readTimeout: time.Second, writeTimeout: time.Second, dialTimeout: time.Second}
	c.c = ic.FromNetConn(clientConn)

	res, err := c.Execute(context.Background(), "CMD", reqParams)
	if err != nil {
		t.Fatalf("execute: %v", err)
	}
	if !strings.Contains(res, "1") {
		t.Fatalf("unexpected response %q", res)
	}
}
