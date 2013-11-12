package random

import (
	"io"
	"math/rand"
	"time"
)

// Randbo creates a stream of non-crypto quality random bytes
type randbo struct {
	rand.Source
}

// New creates a new random reader with a time source.
func New() io.Reader {
	return NewFrom(rand.NewSource(time.Now().UnixNano()))
}

func NewFrom(src rand.Source) io.Reader {
	return &randbo{src}
}

// Read satisfies io.Reader
func (r *randbo) Read(p []byte) (n int, err error) {
	todo := len(p)
	offset := 0
	for {
		val := int64(r.Int63())
		for i := 0; i < 8; i++ {
			p[offset] = byte(val & 0xff)
			todo--
			if todo == 0 {
				return len(p), nil
			}
			offset++
			val >>= 8
		}
	}
}
