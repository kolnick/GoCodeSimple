package server

import "testing"

func TestServer(t *testing.T) {
	server := NewServer("127.0.0.1", 8888)
	server.Start()
}
