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

func Bytes() *builder {
	typesOnce.Do(func() {
		var b strings.Builder
		b.WriteString(UppercaseBytes)
		b.WriteString(LowercaseBytes)
		b.WriteString(NumberBytes)
		typesBuilder.dictionary = New(b.String(),6)
		typesBuilder.source = source
	})
	return typesBuilder
}

// create a dictionary by Uppercase and Lowercase
func Letter() *builder {
	letterOnce.Do(func() {
		var b strings.Builder
		b.WriteString(UppercaseBytes)
		b.WriteString(LowercaseBytes)
		letterBuilder.dictionary = New(b.String(),6)
		letterBuilder.source = source
	})
	return letterBuilder
}

// create a dictionary by Uppercase
func Uppercase() *builder {
	uppercaseOnce.Do(func() {
		var b strings.Builder
		b.WriteString(UppercaseBytes)
		uppercaseBuilder.dictionary = New(b.String(),5)
		uppercaseBuilder.source = source
	})
	return uppercaseBuilder
}

// create a dictionary by Lowercase
func Lowercase() *builder {
	lowercaseOnce.Do(func() {
		var b strings.Builder
		b.WriteString(LowercaseBytes)
		lowercaseBuilder.dictionary = New(b.String(),5)
		lowercaseBuilder.source = source
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