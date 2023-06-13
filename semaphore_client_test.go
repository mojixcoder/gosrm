package gosrm

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHTTPClient(t *testing.T) {
	cfg := HTTPClientConfig{}

	c := NewHTTPClient(cfg).(httpClient)
	assert.Equal(t, http.DefaultClient, c.client)
	assert.Equal(t, 0, cap(c.pool))

	cfg.MaxConcurrency = 100

	c = NewHTTPClient(cfg).(httpClient)
	assert.Equal(t, 100, cap(c.pool))

	cfg.HTTPClient = &http.Client{}
	c = NewHTTPClient(cfg).(httpClient)
	assert.Equal(t, cfg.HTTPClient, c.client)
}

func TestHTTPClient_acquire_and_release(t *testing.T) {
	cfg := HTTPClientConfig{}
	c := NewHTTPClient(cfg).(httpClient)

	for i := 0; i < 5; i++ {
		c.acquire()
	}

	// chan is not used since it's disabled.
	assert.Len(t, c.pool, 0)

	c.release()
	assert.Len(t, c.pool, 0)

	cfg.MaxConcurrency = 2
	c = NewHTTPClient(cfg).(httpClient)

	for i := 0; i < int(cfg.MaxConcurrency); i++ {
		c.acquire()
	}

	// chan is full.
	assert.Len(t, c.pool, 2)

	for i := cfg.MaxConcurrency; i > 0; i-- {
		c.release()
		assert.Len(t, c.pool, int(i-1))
	}

	// chan is empty.
	assert.Len(t, c.pool, 0)
}

func TestHTTPClient_Do(t *testing.T) {
	client := NewHTTPClient(HTTPClientConfig{MaxConcurrency: 1})

	testsrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(201)
	}))

	req, err := http.NewRequest("GET", testsrv.URL, nil)
	assert.NoError(t, err)

	res, err := client.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, 201, res.StatusCode)
}
