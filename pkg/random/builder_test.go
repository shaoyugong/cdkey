package random

import "testing"

var length = 12
var number = 1

//// cmd: go test -bench=. -run=non
//func BenchmarkDictionary(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Bytes().Generate(length)
//	}
//}
//
//func BenchmarkLetter(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Letter().Generate(length)
//	}
//}
//
//func BenchmarkUppercase(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Uppercase().Generate(length)
//	}
//}
//
//func BenchmarkLowercase(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Lowercase().Generate(length)
//	}
//}

func Test_Dictionary(t *testing.T) {
	for i := 0; i < number; i++ {
		Bytes().Generate(length)
	}
	Bytes().Generate(length)
	t.Log("ok")
}

func Test_Letter(t *testing.T) {
	for i := 0; i < number; i++ {
		Letter().Generate(length)
	}
	Letter().Generate(length)
	t.Log("ok")
}

func Test_Uppercase(t *testing.T) {
	for i := 0; i < number; i++ {
		Uppercase().Generate(length)
	}
	Uppercase().Generate(length)
	t.Log("ok")
}

func Test_Lowercase(t *testing.T) {
	for i := 0; i < number; i++ {
		Lowercase().Generate(length)
	}
	Lowercase().Generate(length)
	t.Log("ok")
}