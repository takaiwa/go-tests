package encapsulation

import "testing"

func TestParseUserData(t *testing.T) {

	cases := []struct {
		data     []byte
		expected string
	}{
		{
			data:     []byte("{\"name\": \"Uncle Sam\"}"),
			expected: "Mr. Uncle Sam",
		},
	}

	for index, c := range cases {
		actual, err := ParseUserData(c.data)

		// actual.Username は参照できないのがポイント（他から書き換えられない）

		if err != nil {
			t.Fatal(err.Error())
		} else if actual.Name() != c.expected {
			t.Fatalf("index:%d expected:%s, actual:%s", index, c.expected, actual.Name())
		}
	}
}
