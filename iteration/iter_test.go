package iteration

import (
	"fmt"
	"testing"
)

func TestIter(t *testing.T) {
	repeated := Repeat(4,"a")
	expected := "aaaa"

	if repeated != expected {
		t.Errorf("expected : %q got : %q", expected, repeated)
	}
}

func TestCheckString(t *testing.T) {
	count := CheckString([]string{"aaa","b", "ccc"}, "aacbccc")
	expected := 2
	if count != expected {
		t.Errorf("expected : %d got : %d", expected, count)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i<b.N; i++ {
		Repeat(4,"a")
	}
}

// This is example test for document
func ExampleRepeat() {
	r := Repeat(4,"a")
	fmt.Println(r)
	// Output: aaaa
}