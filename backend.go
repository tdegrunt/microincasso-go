package microincasso

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// Backend is an interface for making calls against a MicroIncasso service.
// This interface exists to enable mocking for during testing if needed.
type Backend interface {
	Call(method, path string, mireq *Request) (*Response, error)
}

// InternalBackend is the internal implementation for making HTTP calls to MicroIncasso.
type InternalBackend struct {
	url        string
	httpClient *http.Client
}

// NewInternalBackend returns a customized backend used for making calls in this binding.
func NewInternalBackend(httpClient *http.Client, url string) *InternalBackend {
	if len(url) == 0 {
		url = defaultURL
	}

	return &InternalBackend{
		url:        url,
		httpClient: httpClient,
	}
}

// SetDebug enables additional tracing globally.
// The method is designed for used during testing.
func SetDebug(value bool) {
	debug = value
}

// GetBackend returns the currently used backend in the binding.
func GetBackend() Backend {
	if backend == nil {
		backend = NewInternalBackend(http.DefaultClient, "")
	}

	return backend
}

// SetBackend sets the backend used in the binding.
func SetBackend(b Backend) {
	backend = b
}

func (s *InternalBackend) Call(method, path string, mireq *Request) (*Response, error) {
	var miresp Response

	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = s.url + path

	b, err := xml.Marshal(mireq)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer([]byte{})
	buf.WriteString(xml.Header)
	buf.Write(b)

	if debug {
		log.Printf("Request body: %s\n", buf.String())
	}

	req, err := http.NewRequest(method, path, buf)
	if err != nil {
		log.Printf("Cannot create request: %v\n", err)
		return nil, err
	}

	req.Header.Add("User-Agent", "Microincasso/v1") // GoBindings/"+clientversion)
	req.Header.Add("Content-Type", "text/xml")

	var start time.Time
	if debug {
		log.Printf("Requesting %v %q\n", method, path)
		start = time.Now()
	}

	res, err := s.httpClient.Do(req)

	if debug {
		log.Printf("Completed in %v\n", time.Since(start))
	}

	if err != nil {
		log.Printf("Request failed: %v\n", err)
		return nil, err
	}
	defer res.Body.Close()

	if debug {
		log.Printf("Response status: %v\n", res.StatusCode)
		log.Printf("Response ContentLength: %v\n", res.ContentLength)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Cannot parse response: %v\n", err)
		return nil, err
	}

	if debug {
		log.Printf("Response: %q\n", resBody)
	}

	if len(resBody) > 2 {
		if err := xml.Unmarshal(resBody, &miresp); err != nil {
			return nil, err
		}
	} else {
		miresp.Status = res.StatusCode
	}

	return &miresp, nil

}
