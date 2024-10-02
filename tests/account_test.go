package tests

import (
	"bytes"
	"fyers-go-sdk/api"
	"fyers-go-sdk/utils"
	"io"
	"net/http"
	"testing"
)

func TestGetProfile(t *testing.T) {
	logger, err := utils.InitializeLogger("test.log")
	if err != nil {
		t.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Close()

	// Simulate successful response with profile data
	mockTransport := &MockTransport{
		RoundTripFunc: func(req *http.Request) *http.Response {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body: io.NopCloser(bytes.NewBufferString(`{
                    "data": {
                        "client_id": "12345",
                        "name": "Test User"
                    },
                    "status": "ok"
                }`)),
				Header: make(http.Header),
			}
		},
	}
	client := MockClientWithTransport(mockTransport)

	// Pass nil as transport if you don't need to mock HTTP responses in this test
	// client := MockClientWithTransport(nil)

	profile, err := api.GetProfile(client, logger)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if profile == nil {
		t.Fatalf("Expected profile data, got nil")
	}

	t.Logf("Profile fetched successfully: %+v", profile)
}

func TestGetFunds(t *testing.T) {
	logger, err := utils.InitializeLogger("test.log")
	if err != nil {
		t.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Close()

	// Simulate successful response with funds data
	mockTransport := &MockTransport{
		RoundTripFunc: func(req *http.Request) *http.Response {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body: io.NopCloser(bytes.NewBufferString(`{
					"data": {
						"balance": 10000
					},
					"status": "ok"
				}`)),
				Header: make(http.Header),
			}
		},
	}
	client := MockClientWithTransport(mockTransport)
	// client := MockClientWithTransport(nil)

	funds, err := api.GetFunds(client, logger)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if funds == nil || funds.Data.Balance != 10000 {
		t.Fatalf("Expected balance of 10000, got: %+v", funds)
	}

	t.Logf("Funds fetched successfully: %+v", funds)
}
