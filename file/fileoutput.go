package file

import (
	"fmt"
	"io"
	"os"
)

var osCreate func(name string) (closer io.WriteCloser, err error)

func int() {
	osCreate = func(name string) (closer io.WriteCloser, err error) {
		return os.Create(name)
	}
}

func Output(filename string, strs []string) error {
	f, err := osCreate(filename)
	if err != nil {
		return err
	}

	defer f.Close()
	for _, str := range strs {
		fmt.Fprint(f, str)
	}
	return nil
}
