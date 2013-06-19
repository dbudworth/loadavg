// +build linux darwin

package loadavg

import "testing"

func TestLoadAvg(t *testing.T) {
	loadavg, err := LoadAvg()
	if err != nil {
		t.Fatal(err)
	}
	for _, l := range loadavg[:] {
		if l < 0 {
			t.Errorf("expected loadavg >= 0, got %v", l)
		}
	}
}
