package generator

import "testing"

func BenchmarkRandomBool(b *testing.B) {
	for i := 0; i < 128; i++ {
		RandomBool()
	}
}

func BenchmarkRandomNum(b *testing.B) {
	for i := 0; i < 128; i++ {
		RandomNum(0, 128)
	}
}
