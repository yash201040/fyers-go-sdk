package tests

import (
	"fyers-go-sdk/utils"
	"net/http"
)

// MockTransport is a custom RoundTripper for mocking HTTP responses.
type MockTransport struct {
	RoundTripFunc func(req *http.Request) *http.Response
}

func (m *MockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.RoundTripFunc(req), nil
}

// MockClient returns a mock Fyers API client with a custom HTTP transport.
func MockClientWithTransport(transport http.RoundTripper) *utils.Client {
	mockHTTPClient := &http.Client{Transport: transport}
	return &utils.Client{
		AppID:       "mock-app-id",
		AccessToken: "mock-access-token",
		BaseURL:     "https://mock-api.fyers.in",
		HttpClient:  mockHTTPClient,
	}
}
