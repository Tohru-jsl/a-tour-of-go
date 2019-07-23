package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}


func (r13 *rot13Reader) Read(b []byte) (int, error) {
	n, e := r13.r.Read(b)
	for i := 0; i < n; i++ {
		if b[i] > 'A' && b[i] < 'z' {
			if b[i] > 'z' - 13 {
				b[i] = b[i] - 13
			} else {
				b[i] = b[i] + 13
			}
		}
	}
	return n, e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}