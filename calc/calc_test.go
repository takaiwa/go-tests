package calc

import "testing"

func TestSum(t *testing.T) {
	patterns := map[string]struct {
		x        int
		y        int
		expected int
	}{
		"初歩的なテスト": {
			x:        1,
			y:        2,
			expected: 3,
		},
		"初歩的なテストをもう一つ": {
			x:        2,
			y:        3,
			expected: 5,
		},
	}

	for index, p := range patterns {
		actual := sum(p.x, p.y)
		if actual != p.expected {
			t.Errorf("index:%d, x:%d, y:%d, expected:%d, actual:%d",
				index, p.x, p.y, p.expected, actual)
		}
	}
}
