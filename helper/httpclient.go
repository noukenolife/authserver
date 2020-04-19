package helper

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

type HTTPClientInterface interface {
	Do(req *http.Request) (*http.Response, error)
}

type MockHTTPClient struct {
	HTTPClientInterface
	mock.Mock
}

func (m *MockHTTPClient) Do(req *http.Request) (res *http.Response, err error) {
	c := m.Called(req)
	if r, ok := c.Get(0).(*http.Response); ok {
		res = r
	} else {
		err = c.Error(1)
	}
	return
}
