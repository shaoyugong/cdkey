package random

import (
	"math/rand"
	"strings"
	"sync"
	"unsafe"
)

const UppercaseBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const LowercaseBytes = "abcdefghijklmnopqrstuvwxyz"
const NumberBytes    = "1234567890"

var (
	typesOnce sync.Once
	typesBuilder = &builder{}
	letterOnce sync.Once
	letterBuilder = &builder{}
	uppercaseOnce sync.Once
	uppercaseBuilder = &builder{}
	lowercaseOnce sync.Once
	lowercaseBuilder = &builder{}
)

type Builder interface {
	Generate(length int) string
}

type builder struct {
	dictionary *Dictionary
	source rand.Source
}

// create builder
func NewBuilder(bytes ...string) *builder {
	var bud strings.Builder
	for _, b := range bytes {
		bud.WriteString(b)
	}

	strs := bud.String()
	bits := GetBytesBits(strs)
	dict := NewDictionary(strs, bits)
	return &builder{
		dictionary: dict,
		source: source,
	}
}

func Bytes() *builder {
	typesOnce.Do(func() {
		typesBuilder = NewBuilder(UppercaseBytes, LowercaseBytes, NumberBytes)
	})

	return typesBuilder
}

// create a dictionary by Uppercase and Lowercase
func Letter() *builder {
	letterOnce.Do(func() {
		letterBuilder = NewBuilder(UppercaseBytes, LowercaseBytes)
	})

	return letterBuilder
}

// create a dictionary by Uppercase
func Uppercase() *builder {
	uppercaseOnce.Do(func() {
		uppercaseBuilder = NewBuilder(UppercaseBytes)
	})

	return uppercaseBuilder
}

// create a dictionary by Lowercase
func Lowercase() *builder {
	lowercaseOnce.Do(func() {
		lowercaseBuilder = NewBuilder(LowercaseBytes)
	})
	return lowercaseBuilder
}

// generate string according to length
func (b *builder) Generate(length int) string {
	bytes := make([]byte, length)

	//  A rand.dictionary.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := length - 1, b.source.Int63(), b.dictionary.index.max; i >= 0; {
		if remain == 0 {
			cache, remain = b.source.Int63(), b.dictionary.index.max
		}

		if idx := int(cache & b.dictionary.index.mask); idx < len(b.dictionary.bytes) {
			bytes[i] = b.dictionary.bytes[idx]
			i--
		}

		cache >>= b.dictionary.index.bits
		remain--
	}

	return *(*string)(unsafe.Pointer(&bytes))
}