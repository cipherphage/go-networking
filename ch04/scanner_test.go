package main

import (
	"bufio"
	"net"
	"reflect"
	"testing"
)

const payload = "The bigger the interface, the weaker the abstraction."

func TestScaner(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:")
	HandleTestError(t, err, "fatal")

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			t.Error(err)
			return
		}
		defer conn.Close()

		_, err = conn.Write([]byte(payload))
		HandleTestError(t, err, "error")
	}()

	conn, err := net.Dial("tcp", listener.Addr().String())
	HandleTestError(t, err, "fatal")
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	scanner.Split(bufio.ScanWords)

	var words []string

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	err = scanner.Err()
	HandleTestError(t, err, "error")

	expected := []string{"The", "bigger", "the", "interface,", "the", "weaker", "the", "abstraction."}

	if !reflect.DeepEqual(words, expected) {
		t.Fatal("inaccurate scanned word list")
	}

	t.Logf("Scanned words: %#v", words)
}
