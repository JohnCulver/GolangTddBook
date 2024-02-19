package counter

import (
	"sync"
	"testing"
)

func TestIncrement(t *testing.T) {

	assertCounter := func(t testing.TB, value int, counter *Counter) {
		t.Helper()
		got := counter.Value()
		want := value
		if value != counter.Value() {
			t.Errorf("Got %d but want %d", got, want)
		}
	}

	t.Run("Test increment to 3", func(t *testing.T) {
		c := &Counter{}
		wantedCount := 3
		for i := 0; i < wantedCount; i++ {
			c.Increment()
		}

		assertCounter(t, 3, c)

	})

	t.Run("Works concurrently", func(t *testing.T) {
		c := &Counter{}
		wantedCount := 1000

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				c.Increment()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, wantedCount, c)

	})

}
