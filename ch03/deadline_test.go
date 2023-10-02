package ch03

import (
	"io"
	"net"
	"testing"
	"time"
)

func TestDeadline(t *testing.T) {
	sync := make(chan struct{})

	checkTestError := func(err error, fatal bool) {
		if err != nil {
			if fatal {
				t.Fatal(err)
			} else {
				t.Error(err)
				return
			}
		}
	}

	listener, err := net.Listen("tcp", "127.0.0.1:")
	checkTestError(err, true)

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			t.Log(err)
			return
		}
		defer func() {
			conn.Close()
			close(sync) // Read from sync shouldn't block due to early return.
		}()

		err = conn.SetDeadline(time.Now().Add(5 * time.Second))
		checkTestError(err, false)

		buf := make([]byte, 1)
		_, err = conn.Read(buf) // Blocked until remote node sends data.
		nErr, ok := err.(net.Error)
		if !ok || !nErr.Timeout() {
			t.Errorf("expected timeout error; actual: %v", err)
		}

		sync <- struct{}{}

		err = conn.SetDeadline(time.Now().Add(5 * time.Second))
		checkTestError(err, false)

		_, err = conn.Read(buf)
		if err != nil {
			t.Error(err)
		}
	}()

	conn, err := net.Dial("tcp", listener.Addr().String())
	checkTestError(err, true)
	defer conn.Close()

	<-sync
	_, err = conn.Write([]byte("1"))
	checkTestError(err, true)

	buf := make([]byte, 1)
	_, err = conn.Read(buf) // Blocked until remote node sends data.
	if err != io.EOF {
		t.Errorf("expected server termination; actual: %v", err)
	}
}
