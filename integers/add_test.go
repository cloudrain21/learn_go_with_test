package ttt

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(2,4)
	expected := 6
	if sum != expected {
		t.Errorf("expected : %d sum %d", expected, sum)
	}
}

// 여기에 comment 를 남기면 go doc 에서 문서로 볼 수 있다.
// Example 도 함께 go doc 에서 보여준다.
func ExampleAdd() {
	sum := Add(1,4)
	fmt.Println(sum)
	// Output: 5
}