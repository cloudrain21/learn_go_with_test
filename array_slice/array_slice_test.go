package array_slice

import (
	"reflect"
	"testing"
)

func TestArraySum(t *testing.T) {
	numbers := [5]int{1,2,3,4,5}
	r := ArraySum(numbers)
	w := 15
	if r != w {
		t.Errorf("want : %d, got : %d", w, r)
	}
}

func TestSliceSum(t *testing.T) {
	numbers := []int{1,2,3,4,5}
	r := SliceSum(numbers)
	w := 15
	if r != w {
		t.Errorf("w : %d, r : %d", w, r)
	}
}

func TestSum(t *testing.T) {
	myAssert := func(t *testing.T, w, r int) {
		t.Helper()
		if r != w {
			t.Errorf("w : %d, r : %d", w, r)
		}
	}
	t.Run("test1", func(t *testing.T) {
		numbers := []int{1,2,3,4,5}
		r := Sum(numbers)
		w := 15
		myAssert(t, w, r)
	})

	t.Run("test2", func(t *testing.T) {
		numbers := []int{1,2,3}
		r := Sum(numbers)
		w := 6
		myAssert(t, w, r)
	})
}

func TestSumAll(t *testing.T) {
	r := SumAll([]int{1,2}, []int{0,9})
	w := []int{3, 9}
	if !reflect.DeepEqual(w,r) {
		t.Errorf("w : %d, r : %d", w, r)
	}
}

func TestSumAllTails(t *testing.T) {
	myAssert := func(t *testing.T, w, r []int) {
		t.Helper()
		if !reflect.DeepEqual(r,w) {
			t.Errorf("w : %d, r : %d", w, r)
		}
	}
	t.Run("test1", func(t *testing.T) {
		r := SumAllTails([]int{1,2}, []int{0,9})
		w := []int{2,9}
		myAssert(t, w, r)
	})

	t.Run("test2", func(t *testing.T) {
		r := SumAllTails([]int{}, []int{2,3,6})
		w := []int{0,9}
		myAssert(t, w, r)
	})

	t.Run("test3", func(t *testing.T) {
		r := SumAllTails([]int{}, []int{})
		w := []int{0,0}
		myAssert(t,w,r)
	})
}