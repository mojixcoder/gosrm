package gosrm

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type (
	// HTTPClient is the interface that can be used to do HTTP calls.
	HTTPClient interface {
		Do(req *http.Request) (*http.Response, error)
	}

	// OSRMClient is the base type with helper methods to call OSRM APIs.
	// It only holds the base OSRM URL.
	OSRMClient struct {
		baseURL *url.URL

		client HTTPClient
	}

	// Request is the OSRM's request structure.
	// It can be used with all services except tile service.
	// Note that for nearest request you have to pass only a coordinate.
	Request struct {
		// Profile is the profile of the request.
		Profile Profile

		// Coordinates is the coordinate of the request.
		Coordinates []Coordinate
	}
)

// New returns a new OSRM client.
func New(baseURL string) (OSRMClient, error) {
	var client OSRMClient

	u, err := url.Parse(baseURL)
	if err != nil {
		return client, err
	}

	client.baseURL = u
	client.SetHTTPClient(NewHTTPClient(HTTPClientConfig{}))

	return client, nil
}

// SetHTTPClient sets the HTTP client that will be used to call OSRM.
func (osrm *OSRMClient) SetHTTPClient(client HTTPClient) {
	if client == nil {
		panic("http client can't be nil")
	}
	osrm.client = client
}

// get calls the given URL and parses the response.
func (osrm OSRMClient) get(ctx context.Context, url string, out any) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req = req.WithContext(ctx)

	res, err := osrm.client.Do(req)
	if err != nil {
		return err
	}

	return json.NewDecoder(res.Body).Decode(out)
}

// applyOpts applys options to the URL.
func (osrm OSRMClient) applyOpts(u *url.URL, opts []Option) {
	for i := 0; i < len(opts); i++ {
		opts[i].apply(u)
	}
}

// buildURLPath builds the path of OSRM's services.
func (req Request) buildURLPath(u url.URL, servicePath string) *url.URL {
	path := strings.TrimSuffix(u.Path, "/")
	coordinates := "/" + convertSliceToStr(req.Coordinates, ";")
	profile := "/" + string(req.Profile)

	u.Path = path + servicePath + profile + coordinates + ".json"

	return &u
}
