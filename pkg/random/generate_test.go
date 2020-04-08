package random

import (
	"testing"
)

var length = 12
var number = 1
var strs = Bytes()
var lett = Letter()
var uppe = Uppercase()
var lowe = Lowercase()
var num = Number()

// cmd: go test -bench=. -run=non
func BenchmarkDictionary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strs.Generate(length)
	}
}

func BenchmarkLetter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lett.Generate(length)
	}
}

func BenchmarkUppercase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uppe.Generate(length)
	}
}

func BenchmarkLowercase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lowe.Generate(length)
	}
}

func BenchmarkNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		num.Generate(length)
	}
}

func Test_Dictionary(t *testing.T) {
	for i := 0; i < number; i++ {
		strs.Generate(length)
	}
	t.Log("ok")
}

func Test_Letter(t *testing.T) {
	for i := 0; i < number; i++ {
		lett.Generate(length)
	}
	t.Log("ok")
}

func Test_Uppercase(t *testing.T) {
	for i := 0; i < number; i++ {
		uppe.Generate(length)
	}
	t.Log("ok")
}

func Test_Lowercase(t *testing.T) {
	for i := 0; i < number; i++ {
		lowe.Generate(length)
	}
	t.Log("ok")
}

func Test_Number(t *testing.T) {
	for i := 0; i < number; i++ {
		num.Generate(length)
	}
	t.Log("ok")
}