package gosrm

import "net/http"

type (
	// httpClient is the default implementation of HTTPClient interface.
	httpClient struct {
		client *http.Client
		pool   chan struct{}
	}

	// HTTPClientConfig is the config used to customize http client.
	HTTPClientConfig struct {
		// MaxConcurrency is the max number of concurrent requests.
		// If it's 0 then there is no limit.
		//
		// Defaults to 0.
		MaxConcurrency uint

		// HTTPClient is the client which will be used to do HTTP calls.
		//
		// Defaults to http.DefaultClient
		HTTPClient *http.Client
	}
)

// acquire acquires a spot in the pool.
func (c httpClient) acquire() {
	if cap(c.pool) == 0 {
		return
	}
	c.pool <- struct{}{}
}

// release releases a spot from the pool.
func (c httpClient) release() {
	if cap(c.pool) == 0 {
		return
	}
	<-c.pool
}

// Do does the HTTP call.
func (c httpClient) Do(req *http.Request) (*http.Response, error) {
	c.acquire()
	defer c.release()

	return c.client.Do(req)
}

// NewHTTPClient returns a new HTTP client.
func NewHTTPClient(cfg HTTPClientConfig) HTTPClient {
	var c httpClient

	if cfg.HTTPClient != nil {
		c.client = cfg.HTTPClient
	} else {
		c.client = http.DefaultClient
	}

	c.pool = make(chan struct{}, cfg.MaxConcurrency)

	return c
}
