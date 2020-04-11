package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	t.Run("reflection basic", func(t *testing.T){
		expected := "Chris"
		var got []string

		x := struct {
			Name string
		}{expected}

		walk(x, func(input string) {
			got = append(got, input)
		})

		if got[0] != expected {
			t.Errorf("wrong. 	got : %s, want : %s\n", got, expected)
		}
	})

	cases := []struct {
		Name string
		Input interface{}
		ExpectedCalls []string
	} {
		{
			"test for one string field struct",
			struct{
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"test for two string field struct",
			struct {
				Name string
				City string
			}{"Chris", "Seoul"},
			[]string{"Chris", "Seoul"},
		},
		{
			"test for various type struct",
			struct {
				Name string
				Age int
			}{"Chris", 30},
			[]string{"Chris"},
		},
		{
			"Nested fields",
			struct {
				Name string
				Profile struct {
					Age  int
					City string
				}
			}{"Chris", struct {
				Age  int
				City string
			}{33, "London"}},
			[]string{"Chris", "London"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got : %v, expected : %v\n", got, test.ExpectedCalls)
			}
		})
	}
}