package fixture

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"testing"
)

type Fixture struct {
	TestCases []TestCase `json:"test_cases"`
}
type TestCase struct {
	X        int `json:"x"`
	Y        int `json:"y"`
	Expected int `json:"expected"`
}

func TestMul(t *testing.T) {
	b, err := ioutil.ReadFile("../testdata/fixture.json")
	if err != nil {
		log.Fatal(err)
	}

	f := new(Fixture)
	if err := json.Unmarshal(b, &f); err != nil {
		log.Fatal(err)
	}

	for index, p := range f.TestCases {
		actual := Mul(p.X, p.Y)
		if actual != p.Expected {
			t.Errorf("index:%d, x:%d, y:%d, expected:%d, actual:%d",
				index, p.X, p.Y, p.Expected, actual)
		}
	}
}
