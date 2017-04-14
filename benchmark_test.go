package optional

import "testing"

func BenchmarkEmpty(b *testing.B) {
	for n := 0; n < b.N; n++ {
		EmptyString()
	}
}

func BenchmarkOf(b *testing.B) {
	for n := 0; n < b.N; n++ {
		OfString("hello")
	}
}
