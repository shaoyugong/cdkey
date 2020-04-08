package random

import (
	"bytes"
	"strconv"
)

const UppercaseBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const LowercaseBytes = "abcdefghijklmnopqrstuvwxyz"
const NumberBytes    = "1234567890"

type Dictionary struct {
	Bytes string	// Bytes
	Bits int 		// binLen bits to represent a Bytes index
	Mask int64 		// All 1-bits, as many as Bits
	Max  int 		// # of Bytes indices fitting in 63 bits
}

func New(bytes string) *Dictionary {
	binLen := len(strconv.FormatInt(int64(len(bytes)), 2))
	return &Dictionary{
		Bytes: bytes,
		Bits: binLen,
		Mask: 1<<binLen - 1,
		Max:  63 / binLen,
	}
}

// Set up a dictionary containing letters and Numbers
func Bytes() *Dictionary {
	var buf bytes.Buffer
	buf.WriteString(UppercaseBytes)
	buf.WriteString(LowercaseBytes)
	buf.WriteString(NumberBytes)
	str := buf.String()
	return New(str)
}

// Set up a dictionary containing Uppercase and Lowercase
func Letter() *Dictionary {
	var buf bytes.Buffer
	buf.WriteString(UppercaseBytes)
	buf.WriteString(LowercaseBytes)
	str := buf.String()
	return New(str)
}

// Set up a dictionary containing Uppercase
func Uppercase() *Dictionary {
	var buf bytes.Buffer
	buf.WriteString(UppercaseBytes)
	str := buf.String()
	return New(str)
}

// Set up a dictionary containing Lowercase
func Lowercase() *Dictionary {
	var buf bytes.Buffer
	buf.WriteString(LowercaseBytes)
	str := buf.String()
	return New(str)
}

// Set up a dictionary containing Number
func Number() *Dictionary {
	var buf bytes.Buffer
	buf.WriteString(NumberBytes)
	str := buf.String()
	return New(str)
}