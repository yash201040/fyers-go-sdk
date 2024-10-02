package tests

import (
	"fyers-go-sdk/api"
	"fyers-go-sdk/models"
	"testing"
)

func TestPlaceLimitOrder(t *testing.T) {
	// mockTransport := &MockTransport{
	// 	RoundTripFunc: func(req *http.Request) *http.Response {
	// 		return &http.Response{
	// 			StatusCode: http.StatusOK,
	// 			Body: io.NopCloser(bytes.NewBufferString(`{
	//                 "id": "order123",
	//                 "status": "success"
	//             }`)),
	// 			Header: make(http.Header),
	// 		}
	// 	},
	// }
	// client := MockClientWithTransport(mockTransport)

	// Pass nil as transport if you don't need to mock HTTP responses in this test
	client := MockClientWithTransport(nil)

	order := models.OrderRequest{
		Symbol:      "NSE:IDEA-EQ",
		Qty:         1,
		Type:        1,
		Side:        1,
		ProductType: "INTRADAY",
		LimitPrice:  7.9,
		Validity:    "DAY",
	}

	orderResp, err := api.PlaceLimitOrder(client, order)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if orderResp == nil {
		t.Fatalf("Expected order response, got nil")
	}

	t.Logf("Order placed successfully: %+v", orderResp)
}

func TestPlaceLimitOrder_Failure(t *testing.T) {
	// mockTransport := &MockTransport{
	// 	RoundTripFunc: func(req *http.Request) *http.Response {
	// 		return &http.Response{
	// 			StatusCode: http.StatusBadRequest,
	// 			Body:       io.NopCloser(bytes.NewBufferString(`{"error":"Invalid order parameters"}`)),
	// 			Header:     make(http.Header),
	// 		}
	// 	},
	// }
	// client := MockClientWithTransport(mockTransport)
	client := MockClientWithTransport(nil)

	order := models.OrderRequest{
		Symbol:      "NSE:IDEA-EQ",
		Qty:         1,
		Type:        1,
		Side:        1,
		ProductType: "INTRADAY",
		LimitPrice:  7.9,
		Validity:    "DAY",
	}

	_, err := api.PlaceLimitOrder(client, order)
	if err == nil {
		t.Fatalf("Expected error for bad request, got nil")
	}

	t.Logf("Received expected error: %v", err)
}
