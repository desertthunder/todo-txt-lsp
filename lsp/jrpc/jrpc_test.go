package jrpc

import (
	"bufio"
	"bytes"
	"strconv"
	"strings"
	"testing"

	"github.com/charmbracelet/log"
)

func TestJRPC(t *testing.T) {
	header := "Content-Length: {000}\r\n\r\n"
	content := "{\"jsonrpc\":\"2.0\",\"id\":1,\"method\":\"textDocument/someMethod\",\"params\":{\"position\":{\"line\":0,\"character\":0},\"text\":\"test\"}}"
	example := []byte(strings.Replace(header, "{000}", strconv.Itoa(len(content)), 1) + content)

	t.Run("SplitFunc implementation for stdin scanner", func(t *testing.T) {

		r := bytes.NewReader(example)
		s := bufio.NewScanner(r)
		s.Split(Split)

		for s.Scan() {
			b := s.Bytes()

			got := strings.Trim(string(b), "\x00") // Remove null bytes
			want := content

			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		}

	})

	t.Run("Decoding of messages", func(t *testing.T) {

		m, err := DecodeMessage(example, log.Default())

		if err != nil {
			t.Errorf("error decoding message: %v", err)
		}

		got := m.Method
		want := "textDocument/someMethod"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		got = m.JSONRPC
		want = "2.0"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Handling of messages", func(t *testing.T) {
		t.Skip("Not yet implemented")
	})
}
