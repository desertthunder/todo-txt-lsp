package jrpc

import (
	"testing"
)

func TestJRPC(t *testing.T) {
	// header := "Content-Length: {000}\r\n\r\n"
	// content := "{\"jsonrpc\":\"2.0\",\"id\":1,\"method\":\"textDocument/someMethod\",\"params\":{\"position\":{\"line\":0,\"character\":0},\"text\":\"test\"}}"
	// example := []byte(strings.Replace(header, "{000}", strconv.Itoa(len(content)), 1) + content)

	t.Run("SplitFunc implementation for stdin scanner", func(t *testing.T) {
		t.Skip("Not yet implemented")

	})

	t.Run("Decoding of messages", func(t *testing.T) {
		t.Skip("Not yet implemented")

	})

	t.Run("Handling of messages", func(t *testing.T) {
		t.Skip("Not yet implemented")
	})
}
