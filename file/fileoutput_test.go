package file

import (
	"bytes"
	"io"
	"testing"
)

type mockCloser struct {
	io.Writer
}

func (f mockCloser) Close() error {
	return nil
}

func TestOutput(t *testing.T) {

	var outputMap map[string]*bytes.Buffer
	SetOsCreate(func(name string) (io.WriteCloser, error) {
		b := &bytes.Buffer{}
		outputMap[name] = b
		return mockCloser{b}, nil
	})

	var patterns = []struct {
		filename string
		data     []string
		expected map[string]string
	}{
		{filename: "hoge.txt",
			data: []string{"foo", "bar"},
			expected: map[string]string{
				"hoge.txt": "foobar",
			},
		},
	}

	for _, p := range patterns {
		outputMap = map[string]*bytes.Buffer{}
		Output(p.filename, p.data)
		for key, value := range p.expected {
			if b, ok := outputMap[key]; !ok {
				t.Errorf("file doesn't exist filename:%s", key)
			} else if b.String() != value {
				t.Errorf("file contents don't match filename:%s, "+
					"expected:%s, actual:%s", key, value, b.String())
			}
		}
	}
}
