package usb

import (
	"fmt"
	"io"
	"os"
)

type StdWriter struct{}

func (s *StdWriter) Write(p []byte) (n int, err error) {
	return fmt.Println(p)
}

func (s *StdWriter) Close() error {
	return nil
}

func NewHID() (io.WriteCloser, error) {
	var w io.WriteCloser
	hidfp, err := os.OpenFile("/dev/hidg0", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		w = &StdWriter{}
	} else {
		w = hidfp

	}
	return w, err
}
