package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("Testing  basic race", func(t *testing.T) {
		fastServer := getNewServer(0 * time.Second)
		slowServer := getNewServer(5 * time.Millisecond)

		defer fastServer.Close()
		defer slowServer.Close()

		fastUrl := fastServer.URL
		slowUrl := slowServer.URL

		want := fastUrl

		got, _ := Racer(slowUrl, fastUrl)

		if got != want {
			t.Errorf("Got %q wanted %q", got, want)
		}
	})

	t.Run("Timeout error if it takes longer than 10 seconds", func(t *testing.T) {
		fastServer := getNewServer(11 * time.Millisecond)
		slowServer := getNewServer(12 * time.Millisecond)

		defer fastServer.Close()
		defer slowServer.Close()

		_, err := ConfigurableRacer(fastServer.URL, slowServer.URL, 5*time.Millisecond)

		if err != ErrorRacerTimeout {
			t.Errorf("Racer should timeout if request takes longer than 10s")
		}

	})
}

func getNewServer(duration time.Duration) *httptest.Server {
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(duration)
		w.WriteHeader(http.StatusOK)
	}))
	return slowServer
}
