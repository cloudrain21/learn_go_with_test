package main

import (
	"bytes"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestCountDown(t *testing.T) {
	myAssert := func(t *testing.T, want, got string) {
		t.Helper()
		if got != want {
			t.Errorf("want : %s, got : %s\n", want, got)
		}
	}

	t.Run("countdown", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		c := &SpySleeper{}
		CountDown(buffer, c)

		got := buffer.String()
		want := `3
2
1
Go!`
		myAssert(t, want, got)
	})

	t.Run("countdown", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		c := &SpySleeper{}
		CountDown(buffer, c)

		got := strconv.Itoa(c.SleepTime)
		want := strconv.Itoa(4)

		myAssert(t, want, got)
	})

	t.Run("real sleep", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		c := DefaultSleeper{}
		CountDown(buffer, c)

		got := buffer.String()
		want := `3
2
1
Go!`
		myAssert(t, want, got)
	})

	t.Run("sequential check", func(t *testing.T) {
		sleeper := &CheckSequenceSleeper{}
		CountDown(sleeper, sleeper)

		got := sleeper.Seq
		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want : %v, got : %v\n", want, got)
		}
	})

	t.Run("configurable sleeper", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		c := &ConfigurableSleeper{1, time.Sleep}
		CountDown(buffer, c)

		got := buffer.String()
		want := `3
2
1
Go!`
		myAssert(t, want, got)
	})
}
