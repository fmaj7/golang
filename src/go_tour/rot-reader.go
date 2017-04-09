package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader) Read(bytes []byte) (int, error) {
	count, ok := reader.r.Read(bytes)
	if ok == nil {
		for index, byte := range bytes[:count] {
			bytes[index] = (byte-'A'+13)%26 + 'A'
		}
		return count, nil
	}
	return 0, ok
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
