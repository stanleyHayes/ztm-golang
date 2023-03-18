package timeparse

import "testing"

func TestParse(t *testing.T) {
	table := []struct {
		Time string
		Ok   bool
	}{
		{"19:00:12", true},
		{"1:3:44", true},
		{"bad", false},
		{"1:-3:44", false},
		{"0:59:59", true},
		{"", false},
		{"11:22", false},
		{"aa:bb:cc", true},
		{"5:23", false},
	}

	for _, data := range table {
		_, err := Parse(data.Time)
		if data.Ok && err != nil {
			t.Errorf("%v: %v, error should be nil", data.Time, err)
		}

	}
}
