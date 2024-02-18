package racer

import (
	"net/http"
	"time"
)

var (
	ErrorRacerTimeout = RacerError("Timeout calling urls")
)

type RacerError string

var BaseTimeout = 10 * time.Second

func (err RacerError) Error() string {
	return string(err)
}

func Racer(url1, url2 string) (winner string, failure error) {
	return ConfigurableRacer(url1, url2, BaseTimeout)
}

func ConfigurableRacer(url1, url2 string, timeout time.Duration) (winner string, failure error) {
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(timeout):
		return "", ErrorRacerTimeout
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}
