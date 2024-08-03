package util

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"time"
)

// GenPort Get the random port number.
func GenPort(count int) []string {
	ports := make([]string, count)
	for i := 0; i < count; i++ {
		port := getRangeUnusedPort(5000, 9999)
		ports[i] = strconv.Itoa(port)
	}
	return ports
}

// getRangeUnusedPort Get the random unoccupied port number in the specified interval.
func getRangeUnusedPort(startPort, endPort int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		port := r.Intn(endPort-startPort+1) + startPort
		addr := fmt.Sprintf(":%d", port)
		listen, err := net.Listen("tcp", addr)
		if err == nil {
			_ = listen.Close()
			return port
		}
	}
}
