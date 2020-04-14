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
	typesBuilder = &builder{}
)

type Builder interface {
	Generate(length int) string
}

type builder struct {
	mutex sync.Mutex
	dictionary *Dictionary
	source rand.Source
}

func Bytes() *builder {
	typesBuilder.mutex.Lock()
	defer typesBuilder.mutex.Unlock()
	if typesBuilder.dictionary != nil && typesBuilder.source != nil {
		return typesBuilder
	}

	var b strings.Builder
	b.WriteString(UppercaseBytes)
	b.WriteString(LowercaseBytes)
	b.WriteString(NumberBytes)
	typesBuilder.dictionary = New(b.String(),6)
	typesBuilder.source = source
	return typesBuilder
}

// create a dictionary by Uppercase and Lowercase
func Letter() *builder {
	typesBuilder.mutex.Lock()
	defer typesBuilder.mutex.Unlock()
	if typesBuilder.dictionary != nil && typesBuilder.source != nil {
		return typesBuilder
	}

	var b strings.Builder
	b.WriteString(UppercaseBytes)
	b.WriteString(LowercaseBytes)
	typesBuilder.dictionary = New(b.String(),6)
	typesBuilder.source = source
	return typesBuilder
}

// create a dictionary by Uppercase
func Uppercase() *builder {
	typesBuilder.mutex.Lock()
	defer typesBuilder.mutex.Unlock()
	if typesBuilder.dictionary != nil && typesBuilder.source != nil {
		return typesBuilder
	}

	var b strings.Builder
	b.WriteString(UppercaseBytes)
	typesBuilder.dictionary = New(b.String(),5)
	typesBuilder.source = source
	return typesBuilder
}

// create a dictionary by Lowercase
func Lowercase() *builder {
	typesBuilder.mutex.Lock()
	defer typesBuilder.mutex.Unlock()
	if typesBuilder.dictionary != nil && typesBuilder.source != nil {
		return typesBuilder
	}

	var b strings.Builder
	b.WriteString(LowercaseBytes)
	typesBuilder.dictionary = New(b.String(),5)
	typesBuilder.source = source
	return typesBuilder
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