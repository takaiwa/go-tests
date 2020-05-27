package fileoutput

import "io"

func SetOsCreate(f func(name string) (closer io.WriteCloser, err error)) {
	osCreate = f
}
