package arp

import (
	"reflect"
	"testing"
)

func TestSelectScanInterfacesKeepsValidConfiguredInterfaces(t *testing.T) {
	available := map[string]bool{
		"enp6s0": true,
		"wlo1":   true,
	}

	selected, invalid, usedAuto := selectScanInterfaces(
		[]string{"enp6s0", "missing0"},
		available,
		[]string{"wlo1"},
	)

	if !reflect.DeepEqual(selected, []string{"enp6s0"}) {
		t.Fatalf("selected = %#v, want enp6s0", selected)
	}
	if !reflect.DeepEqual(invalid, []string{"missing0"}) {
		t.Fatalf("invalid = %#v, want missing0", invalid)
	}
	if usedAuto {
		t.Fatal("usedAuto = true, want false")
	}
}

func TestSelectScanInterfacesFallsBackWhenAllConfiguredInterfacesAreInvalid(t *testing.T) {
	available := map[string]bool{
		"enp6s0": true,
	}

	selected, invalid, usedAuto := selectScanInterfaces(
		[]string{"enP8p1s0"},
		available,
		[]string{"enp6s0"},
	)

	if !reflect.DeepEqual(selected, []string{"enp6s0"}) {
		t.Fatalf("selected = %#v, want enp6s0", selected)
	}
	if !reflect.DeepEqual(invalid, []string{"enP8p1s0"}) {
		t.Fatalf("invalid = %#v, want enP8p1s0", invalid)
	}
	if !usedAuto {
		t.Fatal("usedAuto = false, want true")
	}
}

func TestSelectScanInterfacesUsesAutoWhenConfigIsBlank(t *testing.T) {
	selected, invalid, usedAuto := selectScanInterfaces(
		nil,
		map[string]bool{"enp6s0": true},
		[]string{"enp6s0"},
	)

	if !reflect.DeepEqual(selected, []string{"enp6s0"}) {
		t.Fatalf("selected = %#v, want enp6s0", selected)
	}
	if len(invalid) != 0 {
		t.Fatalf("invalid = %#v, want empty", invalid)
	}
	if !usedAuto {
		t.Fatal("usedAuto = false, want true")
	}
}
