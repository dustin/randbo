package random

import (
	"io"
	"io/ioutil"
	"testing"
)

func TestRandbo(t *testing.T) {
	buf := make([]byte, 16)
	n, err := New().Read(buf)
	if err != nil {
		t.Fatalf("Error reading: %v", err)
	}
	if n != len(buf) {
		t.Fatalf("Short read: %v", n)
	}
	t.Logf("Read %x", buf)
}

func BenchmarkRandbo(b *testing.B) {
	b.SetBytes(int64(b.N))
	r := New()
	io.CopyN(ioutil.Discard, r, int64(b.N))
}
