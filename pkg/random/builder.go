package random

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"unsafe"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewBuilder(bytes string) (*builder, error) {
	if bytes == "" {
		return nil, fmt.Errorf("source: %s can not be empty", bytes)
	}

	indexBits := len(strconv.FormatInt(int64(len(bytes)), 2))
	p := &builder{
		source: bytes,
		indexBits: indexBits,
		indexMask: 1 << indexBits - 1,
		indexMax:  63 / indexBits,
	}

	return p, nil
}

type builder struct {
	source    string
	indexBits int
	indexMask int64
	indexMax  int
}

// generate string according to indexBits
func (p *builder) Generate(length int) string {
	codes := make([]byte, length)
	for i, number, remain := length - 1, rand.Int63(), p.indexMax; i >= 0; {
		if remain == 0 {
			number, remain = rand.Int63(), p.indexMax
		}

		if idx := int(number & p.indexMask); idx < len(p.source) {
			codes[i] = p.source[idx]
			i--
		}

		number >>= p.indexBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&codes))
}