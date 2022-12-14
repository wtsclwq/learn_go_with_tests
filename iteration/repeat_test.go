package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat(3, "a")
	expected := "aaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat(555, "a")
	}
}

func ExampleRepeat() {
	repeated := Repeat(6, "s")
	fmt.Println(repeated)
	// Output: ssssss
}
