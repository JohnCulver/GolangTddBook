package context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type StubStore struct {
	response string
}

func (s *StubStore) Fetch() string {
	return s.response
}

func (s *StubStore) Cancel() {}

func TestServer(t *testing.T) {
	t.Run("Test basic response", func(t *testing.T) {
		data := "hello, world"
		svr := Server(&StubStore{data})

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}
	})
}

type SpyStore struct {
	response  string
	cancelled bool
}

func (store *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return store.response
}

func (store *SpyStore) Cancel() {
	store.cancelled = true
}

func TestCancel(t *testing.T) {
	t.Run("Test cancel", func(t *testing.T) {
		data := "Hello World!"
		store := SpyStore{response: data}
		server := Server(&store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

	})
}
