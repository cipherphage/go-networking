package ch03

import (
	"context"
	"fmt"
	"io"
	"time"
)

func ExamplePinger() {
	ctx, cancel := context.WithCancel(context.Background())
	r, w := io.Pipe() // In lieu of net.Conn.
	done := make(chan struct{})
	resetTimer := make(chan time.Duration, 1)
	resetTimer <- time.Second // Initial ping interval.

	go func() {
		Pinger(ctx, w, resetTimer)
		close(done)
	}()

	receivePing := func(d time.Duration, r io.Reader) {
		if d >= 0 {
			fmt.Printf("resetting timer (%s)\n", d)
			// t.Logf("resetting timer (%s)\n", d)
			resetTimer <- d
		}

		now := time.Now()
		buf := make([]byte, 1024)
		n, err := r.Read(buf)

		if err != nil {
			fmt.Println(err)
			//t.Error(err)
		}

		fmt.Printf("received %q (%s)\n", buf[:n], time.Since(now).Round(100*time.Millisecond))
		// t.Logf("received %q (%s)\n", buf[:n], time.Since(now).Round(100*time.Millisecond))
	}

	for i, v := range []int64{0, 200, 300, 0, -1, -1, -1} {
		fmt.Printf("Run %d:\n", i+1)
		// t.Logf("Run %d:\n", i+1)
		receivePing(time.Duration(v)*time.Millisecond, r)
	}

	cancel()
	<-done // Ensures the pinger exits after cancelling the context.

	// Output:
	// Run 1:
	// resetting timer (0s)
	// received "ping" (1s)
	// Run 2:
	// resetting timer (200ms)
	// received "ping" (200ms)
	// Run 3:
	// resetting timer (300ms)
	// received "ping" (300ms)
	// Run 4:
	// resetting timer (0s)
	// received "ping" (300ms)
	// Run 5:
	// received "ping" (300ms)
	// Run 6:
	// received "ping" (300ms)
	// Run 7:
	// received "ping" (300ms)
}

// EXPECTED OUTPUT
// go test -v -timeout 300s -race -bench=. ./ch03/ping.go ./ch03/ping_example_test.go
// === RUN   TestExamplePinger
//     ping_example_test.go:40: Run 1:
//     ping_example_test.go:24: resetting timer (0s)
//     ping_example_test.go:36: received "ping" (1s)
//     ping_example_test.go:40: Run 2:
//     ping_example_test.go:24: resetting timer (200ms)
//     ping_example_test.go:36: received "ping" (200ms)
//     ping_example_test.go:40: Run 3:
//     ping_example_test.go:24: resetting timer (300ms)
//     ping_example_test.go:36: received "ping" (300ms)
//     ping_example_test.go:40: Run 4:
//     ping_example_test.go:24: resetting timer (0s)
//     ping_example_test.go:36: received "ping" (300ms)
//     ping_example_test.go:40: Run 5:
//     ping_example_test.go:36: received "ping" (300ms)
//     ping_example_test.go:40: Run 6:
//     ping_example_test.go:36: received "ping" (300ms)
//     ping_example_test.go:40: Run 7:
//     ping_example_test.go:36: received "ping" (300ms)
// --- PASS: TestExamplePinger (2.72s)
// PASS
// ok      command-line-arguments  3.409s
