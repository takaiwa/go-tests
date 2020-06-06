package initialisation

import "testing"

func TestGetStruct1(t *testing.T) {
	expected := Struct1{
		x: "A",
		y: []string{"B", "C"},
	}

	actual := GetStruct1()

	if expected.x != actual.x {
		t.Errorf("expected.y:%s, actual:%s", expected.x, actual.x)
	}

	for index, ax := range actual.y {
		if expected.y[index] != ax {
			t.Errorf("index%d expected.y:%s, actual:%s", index, expected.y[index], ax)
		}
	}
}
