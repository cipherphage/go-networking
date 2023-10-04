package main

import (
	"crypto/rand"
	"io"
	"net"
	"testing"
)

func TestReadIntoBuffer(t *testing.T) {
	payload := make([]byte, 1<<24) // 16 MB.
	_, err := rand.Read(payload)   // Generate a random payload.
	HandleTestError(t, err, "fatal")

	listener, err := net.Listen("tcp", "127.0.0.1:")
	HandleTestError(t, err, "fatal")

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			t.Log(err)
			return
		}
		defer conn.Close()

		_, err = conn.Write(payload)
		HandleTestError(t, err, "error")
	}()

	conn, err := net.Dial("tcp", listener.Addr().String())
	HandleTestError(t, err, "fatal")

	buf := make([]byte, 1<<19) // 512 KB.

	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				t.Error(err)
			}
			break
		}

		t.Logf("read %d bytes", n) // buf[:n] is the data read from conn.
	}

	conn.Close()
}
