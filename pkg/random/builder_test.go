package random

import (
	"testing"
)

const (
	UppercaseBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LowercaseBytes = "abcdefghijklmnopqrstuvwxyz"
	NumberBytes    = "1234567890"
)

var codeBuilder, _ = BuilderManagerInstance.Get(UppercaseBytes + LowercaseBytes + NumberBytes)
var letterBuilder, _ = BuilderManagerInstance.Get(UppercaseBytes + LowercaseBytes)
var numberBuilder, _ = BuilderManagerInstance.Get(NumberBytes)

var length = 12
var number = 1


// cmd: go test -bench=. -run=non
func BenchmarkCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		codeBuilder.Generate(length)
	}
}

func BenchmarkLetter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		letterBuilder.Generate(length)
	}
}

func BenchmarkNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numberBuilder.Generate(length)
	}
}

func Test_Code(t *testing.T) {
	for i := 0; i < number; i++ {
		codeBuilder.Generate(length)
	}
	t.Log("ok")
}

func Test_Letter(t *testing.T) {
	for i := 0; i < number; i++ {
		letterBuilder.Generate(length)
	}
	t.Log("ok")
}

func Test_Number(t *testing.T) {
	for i := 0; i < number; i++ {
		numberBuilder.Generate(length)
	}
	t.Log("ok")
}