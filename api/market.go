package api

import (
	"context"
	"encoding/json"
	"fyers-go-sdk/models"
	"fyers-go-sdk/utils"
	"net/http"
	"time"
)

// GetQuotes fetches the market quotes for a given symbol.
func GetQuotes(client *utils.Client, symbol string, logger *utils.Logger) (*models.QuoteResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", client.BaseURL+"/data/quotes/?symbols="+symbol, nil)
	if err != nil {
		logger.Error("Failed to create request for symbol quote: %v, error: %v", symbol, err)
		return nil, err
	}

	req.Header.Set("Authorization", client.AppID+":"+client.AccessToken)

	resp, err := client.HttpClient.Do(req)
	if err != nil {
		logger.Error("Error fetching quotes for symbol: %v, error: %v", symbol, err)
		return nil, err
	}
	defer resp.Body.Close()

	if err := utils.CheckHTTPResponse(resp); err != nil {
		logger.Error("Received non-OK response for symbol quote: %v, error: %v", symbol, err)
		return nil, err
	}

	var quoteResp models.QuoteResponse
	err = json.NewDecoder(resp.Body).Decode(&quoteResp)
	if err != nil {
		logger.Error("Failed to decode quotes for symbol: %v, error: %v", symbol, err)
		return nil, err
	}

	logger.Info("Successfully fetched quotes for symbol: %v", symbol)
	return &quoteResp, nil
}

// GetMarketDepth fetches the market depth for a given symbol.
func GetMarketDepth(client *utils.Client, symbol string, logger *utils.Logger) (*models.MarketDepthResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", client.BaseURL+"/data/depth/?symbol="+symbol+"&ohlcv_flag=1", nil)
	if err != nil {
		logger.Error("Failed to create request for symbol market depth: %v, error: %v", symbol, err)
		return nil, err
	}

	req.Header.Set("Authorization", client.AppID+":"+client.AccessToken)

	resp, err := client.HttpClient.Do(req)
	if err != nil {
		logger.Error("Error fetching market depth for symbol: %v, error: %v", symbol, err)
		return nil, err
	}
	defer resp.Body.Close()

	if err := utils.CheckHTTPResponse(resp); err != nil {
		logger.Error("Received non-OK response for symbol market depth: %v, error: %v", symbol, err)
		return nil, err
	}

	var depthResp models.MarketDepthResponse
	err = json.NewDecoder(resp.Body).Decode(&depthResp)
	if err != nil {
		logger.Error("Failed to decode market depth for symbol: %v, error: %v", symbol, err)
		return nil, err
	}

	return &depthResp, nil
}

// GetHistoricalData fetches historical data for a given symbol.
func GetHistoricalData(client *utils.Client, symbol string, from string, to string, logger *utils.Logger) (*models.HistoricalDataResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", client.BaseURL+"/data/history?symbol="+symbol+"&resolution=30&date_format=1&range_from="+from+"&range_to="+to+"&cont_flag=1", nil)
	if err != nil {
		logger.Error("Failed to create request for symbol historical data: %v, error: %v", symbol, err)
		return nil, err
	}

	req.Header.Set("Authorization", client.AppID+":"+client.AccessToken)

	resp, err := client.HttpClient.Do(req)
	if err != nil {
		logger.Error("Error fetching historical data for symbol: %v, error: %v", symbol, err)
		return nil, err
	}
	defer resp.Body.Close()

	if err := utils.CheckHTTPResponse(resp); err != nil {
		logger.Error("Received non-OK response for symbol historical data: %v, error: %v", symbol, err)
		return nil, err
	}

	var historicalResp models.HistoricalDataResponse
	err = json.NewDecoder(resp.Body).Decode(&historicalResp)
	if err != nil {
		logger.Error("Failed to decode historical data for symbol: %v, error: %v", symbol, err)
		return nil, err
	}

	return &historicalResp, nil
}
