package portscan

import (
	"net"
	"testing"
	"time"
)

func TestIsOpenDetectsLocalIPv4Listener(t *testing.T) {
	listener, err := net.Listen("tcp4", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("net.Listen: %v", err)
	}
	defer listener.Close()

	accepted := make(chan struct{})
	go func() {
		conn, err := listener.Accept()
		if err == nil {
			conn.Close()
		}
		close(accepted)
	}()

	host, port, err := net.SplitHostPort(listener.Addr().String())
	if err != nil {
		t.Fatalf("SplitHostPort: %v", err)
	}

	if !IsOpen(host, port) {
		t.Fatalf("IsOpen(%q, %q) = false, want true", host, port)
	}

	select {
	case <-accepted:
	case <-time.After(time.Second):
		t.Fatal("listener did not accept the connection")
	}
}

func TestIsOpenReturnsFalseForClosedIPv4Port(t *testing.T) {
	listener, err := net.Listen("tcp4", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("net.Listen: %v", err)
	}

	host, port, err := net.SplitHostPort(listener.Addr().String())
	if err != nil {
		listener.Close()
		t.Fatalf("SplitHostPort: %v", err)
	}

	if err := listener.Close(); err != nil {
		t.Fatalf("listener.Close: %v", err)
	}

	if IsOpen(host, port) {
		t.Fatalf("IsOpen(%q, %q) = true, want false", host, port)
	}
}

func TestTargetAddressFormatsIPv4(t *testing.T) {
	target := targetAddress("127.0.0.1", "8840")
	if target != "127.0.0.1:8840" {
		t.Fatalf("targetAddress returned %q, want 127.0.0.1:8840", target)
	}
}

func TestTargetAddressFormatsIPv6(t *testing.T) {
	target := targetAddress("2001:db8::1", "443")
	if target != "[2001:db8::1]:443" {
		t.Fatalf("targetAddress returned %q, want [2001:db8::1]:443", target)
	}

	host, port, err := net.SplitHostPort(target)
	if err != nil {
		t.Fatalf("SplitHostPort(%q): %v", target, err)
	}
	if host != "2001:db8::1" || port != "443" {
		t.Fatalf("SplitHostPort(%q) = (%q, %q), want (2001:db8::1, 443)", target, host, port)
	}
}
