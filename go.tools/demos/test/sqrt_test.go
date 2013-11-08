package sqrt

import (
	"fmt"
	"math/rand"
	"testing"
)

var testCases = []struct {
	got  int
	want int
}{
	{got: 1, want: 1},
	{got: 2, want: 4},
	{got: 9, want: 81},
}

// TextXXX
func TestSqrt(t *testing.T) {
	for _, tc := range testCases {
		result := Sqrt(tc.got)
		if tc.want != result {
			t.Errorf("Didn't match: want(%d), result(%d)", tc.want, result)
			continue
		}
	}
}

// BenchMarkXXX
func BenchmarkSqrt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sqrt(random())
	}
}

// Example
func ExampleHello() {
	fmt.Println("hello")
	fmt.Println("goodbye")
	// Output:
	// hello
	// goodbye
}

func random() int {
	return rand.Intn(100)
}
