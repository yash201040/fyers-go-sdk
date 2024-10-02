package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fyers-go-sdk/models"
	"fyers-go-sdk/utils"
	"net/http"
	"time"
)

// PlaceLimitOrder places a limit order.
func PlaceLimitOrder(client *utils.Client, order models.OrderRequest) (*models.OrderResponse, error) {
	reqBody, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	request, err := http.NewRequestWithContext(ctx, "POST", client.BaseURL+"/api/v3/orders/sync", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", client.AppID+":"+client.AccessToken)
	request.Header.Set("Content-Type", "application/json")

	response, err := client.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if err := utils.CheckHTTPResponse(response); err != nil {
		return nil, err
	}

	var orderResp models.OrderResponse
	err = json.NewDecoder(response.Body).Decode(&orderResp)
	if err != nil {
		return nil, err
	}

	return &orderResp, nil
}
