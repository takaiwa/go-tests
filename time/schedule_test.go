package time

import (
	"testing"
	"time"
)

func TestCheck(t *testing.T) {
	deliveryTime := time.Date(2020, 7, 31, 10, 0, 0, 0, time.Local)

	cases := []struct {
		hour     int
		min      int
		expected bool
	}{
		{hour: 9, min: 20, expected: false},
		{hour: 9, min: 30, expected: true},
		{hour: 9, min: 40, expected: true},
		{hour: 10, min: 0, expected: false},
		{hour: 10, min: 10, expected: false},
	}

	for idx, c := range cases {

		now = func() time.Time { return time.Date(2020, 7, 31, c.hour, c.min, 0, 0, time.Local) }

		actual := CanNotify(deliveryTime)
		if actual != c.expected {
			t.Error("case:", idx, "actual:", actual, ", c.expected:", c.expected)
		}
	}

}
