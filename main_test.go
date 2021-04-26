package main

import (
	"testing"
)

func Test_slicePorts_single(t *testing.T) {
	want := []string{
		"80",
	}

	portsToSlice := "80"
	got := slicePorts(&portsToSlice)

	if len(got) == len(want) {
		for i, results := range got {
			if results != want[i] {
				t.Errorf("scanPorts(%q) == %q, want %q", want, got, want)
				return
			}
		}
	} else {
		t.Errorf("scanPorts(%q) == slice len:%d, want len:%d", want, len(got), len(want))
	}
	return
}

func Test_slicePorts_range(t *testing.T) {
	want := []string{
		"80",
		"81",
		"82",
	}

	portsToSlice := "80-82"
	got := slicePorts(&portsToSlice)

	if len(got) == len(want) {
		for i, results := range got {
			if results != want[i] {
				t.Errorf("scanPorts(%q) == %q, want %q", want, got, want)
				return
			}
		}
	} else {
		t.Errorf("scanPorts(%q) == slice len:%d, want len:%d", want, len(got), len(want))
	}
	return
}
