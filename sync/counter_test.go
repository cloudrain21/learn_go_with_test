package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	countAssert := func(t *testing.T, want int, got *Counter) {
		if got.Value() != want {
			t.Errorf("want : %d, got : %d\n", want, got.Value())
		}
	}

	t.Run("basic", func(t *testing.T) {
		counter := &Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		countAssert(t, 3, counter)
	})

	t.Run("test by goroutine", func(t *testing.T) {
		counter := NewCounter()
		want := 1000000
		wg := sync.WaitGroup{}
		wg.Add(want)

		for i := 0; i < want; i++ {
			go func(c *Counter) {
				defer wg.Done()
				c.Inc()
			}(counter)
		}
		wg.Wait()
		countAssert(t, want, counter)
	})
}
