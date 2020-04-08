package random

import (
	"math/rand"
	"time"
	"unsafe"
)

var src = rand.NewSource(time.Now().UnixNano())

// Generating random string of length n
func (d *Dictionary) Generate(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), d.Max; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), d.Max
		}
		if idx := int(cache & d.Mask); idx < len(d.Bytes) {
			b[i] = d.Bytes[idx]
			i--
		}
		cache >>= d.Bits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}