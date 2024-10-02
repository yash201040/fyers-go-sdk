package tests

import (
	"fyers-go-sdk/api"
	"fyers-go-sdk/utils"
	"testing"
)

func TestGetQuotes_NonOKStatus(t *testing.T) {
	// Use the MockTransport for simulating responses

	// mockTransport := &MockTransport{
	// 	RoundTripFunc: func(req *http.Request) *http.Response {
	// 		// Simulate a non-OK status code
	// 		return &http.Response{
	// 			StatusCode: http.StatusInternalServerError,
	// 			Body:       io.NopCloser(bytes.NewBufferString("")),
	// 			Header:     make(http.Header),
	// 		}
	// 	},
	// }
	// client := MockClientWithTransport(mockTransport)
	client := MockClientWithTransport(nil)

	logger, _ := utils.InitializeLogger("test.log")
	defer logger.Close()

	_, err := api.GetQuotes(client, "NSE:SBIN-EQ", logger)
	if err == nil {
		t.Fatalf("Expected error for non-OK status code, got nil")
	} else {
		t.Logf("Received expected error: %v", err)
	}
}

func TestGetHistoricalData(t *testing.T) {
	logger, err := utils.InitializeLogger("test.log")
	if err != nil {
		t.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Close()

	// Simulate successful historical data response
	// mockTransport := &MockTransport{
	// 	RoundTripFunc: func(req *http.Request) *http.Response {
	// 		return &http.Response{
	// 			StatusCode: http.StatusOK,
	// 			Body: io.NopCloser(bytes.NewBufferString(`{
	//                 "data": [{
	//                     "time": 1627814400,
	//                     "open": 100,
	//                     "high": 110,
	//                     "low": 90,
	//                     "close": 105,
	//                     "volume": 5000
	//                 }],
	//                 "status": "ok"
	//             }`)),
	// 			Header: make(http.Header),
	// 		}
	// 	},
	// }
	// client := MockClientWithTransport(mockTransport)
	client := MockClientWithTransport(nil)

	historicalData, err := api.GetHistoricalData(client, "NSE:SBIN-EQ", "2021-01-01", "2021-01-02", logger)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if len(historicalData.Data) == 0 || historicalData.Data[0].Close != 105 {
		t.Fatalf("Expected closing price of 105, got: %+v", historicalData)
	}

	t.Logf("Historical data fetched successfully: %+v", historicalData)
}
