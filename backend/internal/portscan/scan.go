package portscan

import (
	"fmt"
	"net"
	"time"
)

// IsOpen - check one tcp port
func IsOpen(host, port string) bool {

	timeout := 3 * time.Second
	target := fmt.Sprintf("%s:%s", host, port)

	conn, err := net.DialTimeout("tcp", target, timeout)

	if err == nil {
		err = conn.Close()
		if err == nil {
			return true
		}
	}

	return false
}
