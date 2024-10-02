package api

import (
	"encoding/json"
	"fmt"
	"fyers-go-sdk/models"
	"fyers-go-sdk/utils"
	"io"
	"net/http"
)

// GetProfile fetches the account profile information.
func GetProfile(client *utils.Client, logger *utils.Logger) (*models.ProfileResponse, error) {
	req, err := http.NewRequest("GET", client.BaseURL+"/api/v3/profile", nil)
	if err != nil {
		logger.Error("Failed to create profile request: %v", err)
		return nil, fmt.Errorf("getProfile: failed to create request: %w", err)
	}

	req.Header.Set("Authorization", client.AppID+":"+client.AccessToken)

	resp, err := client.HttpClient.Do(req)
	if err != nil {
		logger.Error("Error making profile request: %v", err)
		return nil, fmt.Errorf("getProfile: error making request: %w", err)
	}
	defer resp.Body.Close()

	if err := utils.CheckHTTPResponse(resp); err != nil {
		body, _ := io.ReadAll(resp.Body)
		logger.Error("Profile request returned non-OK status: %v, body: %s", resp.StatusCode, string(body))
		return nil, err
	}

	var profileResp models.ProfileResponse
	err = json.NewDecoder(resp.Body).Decode(&profileResp)
	if err != nil {
		logger.Error("Failed to decode profile response: %v", err)
		return nil, fmt.Errorf("getProfile: failed to decode response: %w", err)
	}

	logger.Info("Fetched profile successfully")
	return &profileResp, nil
}

// GetFunds fetches the funds available in the account.
func GetFunds(client *utils.Client, logger *utils.Logger) (*models.FundsResponse, error) {
	req, err := http.NewRequest("GET", client.BaseURL+"/api/v3/funds", nil)
	if err != nil {
		logger.Error("Failed to create funds request: %v", err)
		return nil, fmt.Errorf("getFunds: failed to create request: %w", err)
	}

	req.Header.Set("Authorization", client.AppID+":"+client.AccessToken)

	resp, err := client.HttpClient.Do(req)
	if err != nil {
		logger.Error("Error making funds request: %v", err)
		return nil, fmt.Errorf("getFunds: error making request: %w", err)
	}
	defer resp.Body.Close()

	if err := utils.CheckHTTPResponse(resp); err != nil {
		logger.Error("Funds request returned non-OK status: %v", resp.StatusCode)
		return nil, err
	}

	var fundsResp models.FundsResponse
	err = json.NewDecoder(resp.Body).Decode(&fundsResp)
	if err != nil {
		logger.Error("Failed to decode funds response: %v", err)
		return nil, fmt.Errorf("getFunds: failed to decode response: %w", err)
	}

	logger.Info("Fetched funds successfully")
	return &fundsResp, nil
}
