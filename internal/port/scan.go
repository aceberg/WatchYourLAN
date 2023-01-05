package port

import (
	"fmt"
	"net"
	"time"
)

func isOpen(host, proto string, port int) bool {

	timeout := 3 * time.Second
	target := fmt.Sprintf("%s:%d", host, port)

	conn, err := net.DialTimeout(proto, target, timeout)
	if err != nil {
		return false
	}

	if conn != nil {
		conn.Close()
		return true
	}

	return false
}

// Scan - scan all TCP ports of a host
func Scan(host, proto string, begin, end int) []string {
	var onePort string
	var ports []string

	ports = []string{}

	for i := begin; i < end+1; i++ {
		if isOpen(host, proto, i) {
			onePort = fmt.Sprintf("%s:%d", host, i)
			ports = append(ports, onePort)
		}
	}

	return ports
}
