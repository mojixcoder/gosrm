package gosrm

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const invalidURL string = "postgres://user:abc{DEf1=ghi@example.com:5432/db?sslmode=require"

func newOSRMClient() OSRMClient {
	return OSRMClient{client: NewHTTPClient(HTTPClientConfig{})}
}

func TestNew(t *testing.T) {
	testCases := []struct {
		name, baseURL string
		ok            bool
	}{
		{name: "ok", baseURL: "http://www.test.com:5000", ok: true},
		{name: "not_ok", baseURL: invalidURL, ok: false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			cli, err := New(testCase.baseURL)

			if testCase.ok {
				assert.NoError(t, err)
				assert.Equal(t, "http", cli.baseURL.Scheme)
				assert.Equal(t, "www.test.com:5000", cli.baseURL.Host)
			} else {
				assert.NotNil(t, err)
				assert.Error(t, err)
			}
		})
	}
}

func TestOSRMClient_SetHTTPClient(t *testing.T) {
	osrm := newOSRMClient()

	assert.PanicsWithValue(t, "http client can't be nil", func() {
		osrm.SetHTTPClient(nil)
	})

	client := NewHTTPClient(HTTPClientConfig{MaxConcurrency: 10})

	osrm.SetHTTPClient(client)

	assert.Equal(t, client, osrm.client)
}

func TestOSRMClient_get(t *testing.T) {
	osrm := newOSRMClient()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("{\"message\": \"Ok\"}"))
	}))
	defer srv.Close()

	res := make(map[string]any)

	err := osrm.get(context.Background(), invalidURL, &res)
	assert.Error(t, err)

	err = osrm.get(context.Background(), "http://invalid_host", &res)
	assert.Error(t, err)

	err = osrm.get(context.Background(), srv.URL, &res)
	assert.NoError(t, err)
	assert.Equal(t, "Ok", res["message"])
}

func TestOSRMClient_applyOpts(t *testing.T) {
	osrm := newOSRMClient()
	u := url.URL{}

	osrm.applyOpts(&u, []Option{
		optionImpl(func(u *url.URL) {
			u.Path = "test"
		}),
		optionImpl(func(u *url.URL) {
			u.Path += "/test2"
		}),
	})

	assert.Equal(t, "test/test2", u.Path)
}

func TestRequest_buildURLPath(t *testing.T) {
	osrm, err := New("http://127.0.0.1:5000")
	assert.NoError(t, err)

	req := Request{
		Coordinates: []Coordinate{{13.388860, 52.517037}, {13.397634, 52.529407}, {13.428555, 52.523219}},
		Profile:     ProfileCar,
	}

	u := req.buildURLPath(*osrm.baseURL, tripServiceURL)

	assert.Equal(t, "/trip/v1/car/13.388860,52.517037;13.397634,52.529407;13.428555,52.523219.json", u.Path)
}

func getOSRMAddress() string {
	return os.Getenv("OSRM_ADDRESS")
}
