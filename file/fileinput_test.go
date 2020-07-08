package file

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

// ファイル読み込み自体のテスト
func TestLoadCsv(t *testing.T) {
	path := "../testdata/fileinput.csv"
	actualMap, err := LoadCsv(path)

	if err == nil {
		if 0 == len(actualMap) {
			t.Errorf("actualMap size:%d", len(actualMap))
		} else {
			fmt.Println("actualMap size:", len(actualMap))
		}
	} else {
		t.Error(err)
	}
}

// ファイルを読み込んだ中身を検証するテスト
func TestCsvToMap(t *testing.T) {

	cases := []struct {
		in       io.Reader
		expected map[string]Data
	}{
		struct {
			in       io.Reader
			expected map[string]Data
		}{
			in: strings.NewReader(
				"id,name,age\n" +
					"1001,foo,20\n" +
					"1002,bar,35\n"),
			expected: map[string]Data{"id": Data{"1001", "1002"}, "name": Data{"foo", "bar"}, "age": Data{"20", "35"}},
		},
	}

	for cidx, c := range cases {
		actual, err := csvToMap(c.in)
		if err == nil {
			if len(c.expected) != len(actual) {
				t.Errorf("case:%d, len(c.expected):%d, len(actual):%d",
					cidx, len(c.expected), len(actual))
			}

			for key := range c.expected {
				for index, e := range c.expected[key] {
					if e != actual[key][index] {
						t.Errorf("key:%s, expected:%s, actual:%s",
							key, e, actual[key][index])
					}
				}
			}
		} else {
			t.Error(err)
		}
	}
}
