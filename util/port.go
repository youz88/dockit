package util

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"sync"
	"time"
)

// GenPorts Get the random port number.
func GenPorts(count int) []string {
	if count < 1 {
		return nil
	}

	if count == 1 {
		return []string{getRangeUnusedPort(5000, 9999)}
	}

	// Use coroutine concurrency to obtain unused ports.
	ports := make([]string, count)
	wg := sync.WaitGroup{}
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(i int) {
			defer wg.Done()
			ports[i] = getRangeUnusedPort(5000, 9999)
		}(i)
	}
	wg.Wait()
	return ports
}

// GenPort Get the random port number.
func GenPort() string {
	return GenPorts(1)[0]
}

// getRangeUnusedPort Get the random unoccupied port number in the specified interval.
func getRangeUnusedPort(startPort, endPort int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		port := r.Intn(endPort-startPort+1) + startPort
		addr := fmt.Sprintf(":%d", port)
		listen, err := net.Listen("tcp", addr)
		if err == nil {
			_ = listen.Close()
			return strconv.Itoa(port)
		}
	}
}
