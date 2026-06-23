package routines

import (
	"testing"

	"github.com/aceberg/WatchYourLAN/internal/models"
)

func TestFoundHostIndexKeepsSameMACDifferentIPs(t *testing.T) {
	index := newFoundHostIndex([]models.Host{
		{IP: "192.168.57.146", Mac: "48:b0:2d:eb:f5:15"},
		{IP: "192.168.57.209", Mac: "48:b0:2d:eb:f5:15"},
	})

	host, ok := index.match(models.Host{IP: "192.168.57.209", Mac: "48:b0:2d:eb:f5:15"}, false)
	if !ok {
		t.Fatal("expected exact MAC+IP match")
	}
	if host.IP != "192.168.57.209" {
		t.Fatalf("matched wrong IP: %s", host.IP)
	}

	index.delete(host)

	remaining := index.remaining()
	if len(remaining) != 1 {
		t.Fatalf("expected one remaining host, got %d", len(remaining))
	}
	if remaining[0].IP != "192.168.57.146" {
		t.Fatalf("expected 192.168.57.146 to remain, got %s", remaining[0].IP)
	}
}

func TestFoundHostIndexUsesMACForSingleAddressChange(t *testing.T) {
	index := newFoundHostIndex([]models.Host{
		{IP: "192.168.57.210", Mac: "48:b0:2d:eb:f5:15"},
	})

	host, ok := index.match(models.Host{IP: "192.168.57.209", Mac: "48:b0:2d:eb:f5:15"}, true)
	if !ok {
		t.Fatal("expected single MAC match for address change")
	}
	if host.IP != "192.168.57.210" {
		t.Fatalf("matched wrong IP: %s", host.IP)
	}
}

func TestFoundHostIndexDoesNotUseMACFallbackForKnownMultiIPHost(t *testing.T) {
	index := newFoundHostIndex([]models.Host{
		{IP: "192.168.57.146", Mac: "48:b0:2d:eb:f5:15"},
	})

	_, ok := index.match(models.Host{IP: "192.168.57.209", Mac: "48:b0:2d:eb:f5:15"}, false)
	if ok {
		t.Fatal("did not expect MAC fallback when same MAC is already known on multiple IPs")
	}
}
