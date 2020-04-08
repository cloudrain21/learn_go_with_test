package structs

import (
	"strconv"
	"testing"
)

func TestPerimeter(t *testing.T) {
	r := GetPerimeter(10.0, 10.0)
	w := 40.0
	if r != w {
		t.Errorf("w : %.2f, r : %.2f", w, r)
	}
}

func TestArea(t *testing.T) {
	myAssert := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		ret := shape.Area()
		if ret != want {
			t.Errorf("ret : %g want : %g", ret, want)
		}
	}
	rec := Rectangle{10.0, 20.0}
	myAssert(t, rec, 200.0)

	cir := Circle{10.0}
	myAssert(t, cir, 314.1592653589793)

	// table driven test
	areaTests := []struct{
		shape Shape
		want float64
	} {
		{"Rectangle", Rectangle{10.0,20.0}, 200.0},
		{"Circle", Circle{10.0}, 314.1592653589793},
		{"Triangle", Triangle{10.0, 10.0}, 50.0},
	}
	for i, tt := range areaTests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ret := tt.shape.Area()
			want := tt.want
			if ret != want {
				t.Errorf("want : %g ret : %g", want, ret)
			}
		})
	}
}
