package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"fyers-go-sdk/models"
	"fyers-go-sdk/utils"
	"net/http"
)

// GenerateAuthCode generates an auth code for OAuth.
func GenerateAuthCode(client *utils.Client, req models.AuthRequest, logger *utils.Logger) (*models.AuthResponse, error) {
	reqBody, err := json.Marshal(req)
	if err != nil {
		logger.Error("Failed to marshal auth request: %v", err)
		return nil, fmt.Errorf("generateAuthCode: failed to marshal auth request: %w", err)
	}

	request, err := http.NewRequest("GET", client.BaseURL+"/api/v3/generate-authcode", bytes.NewBuffer(reqBody))
	if err != nil {
		logger.Error("Failed to create auth code request: %v", err)
		return nil, fmt.Errorf("generateAuthCode: failed to create auth code request: %w", err)
	}

	request.Header.Set("Authorization", client.AppID+":"+client.AccessToken)
	request.Header.Set("Content-Type", "application/json")

	response, err := client.HttpClient.Do(request)
	if err != nil {
		logger.Error("Error making auth code request: %v", err)
		return nil, fmt.Errorf("generateAuthCode: error making auth code request: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		logger.Warn("Auth code request returned non-OK status: %v", response.StatusCode)
		return nil, fmt.Errorf("generateAuthCode: received non-OK status: %d", response.StatusCode)
	}

	var authResp models.AuthResponse
	err = json.NewDecoder(response.Body).Decode(&authResp)
	if err != nil {
		logger.Error("Failed to decode auth response: %v", err)
		return nil, fmt.Errorf("generateAuthCode: failed to decode auth response: %w", err)
	}

	logger.Info("Successfully generated auth code for state: %v", authResp.State)
	return &authResp, nil
}

// GenerateAuthURL generates the OAuth URL to get the authorization code.
// `redirectURI` is the URI where users will be redirected after login.
func GenerateAuthURL(appID, redirectURI, state string) string {
	return fmt.Sprintf("https://api.fyers.in/api/v2/generate-authcode?client_id=%s&redirect_uri=%s&response_type=code&state=%s", appID, redirectURI, state)
}

// GetAccessToken exchanges the authorization code for an access token.
func GetAccessToken(appID, secretID, authCode, redirectURI string) (*models.TokenResponse, error) {
	tokenReq := models.TokenRequest{
		ClientID:    appID,
		SecretID:    secretID,
		GrantType:   "authorization_code",
		Code:        authCode,
		RedirectURI: redirectURI,
	}

	reqBody, err := json.Marshal(tokenReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal token request: %w", err)
	}

	req, err := http.NewRequest("POST", "https://api.fyers.in/api/v2/token", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create token request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error during HTTP request: %w", err)
	}
	defer resp.Body.Close()

	var tokenResp models.TokenResponse
	err = json.NewDecoder(resp.Body).Decode(&tokenResp)
	if err != nil {
		return nil, fmt.Errorf("failed to decode token response: %w", err)
	}

	if tokenResp.Error != "" {
		return nil, fmt.Errorf("token response error: %s", tokenResp.Error)
	}

	return &tokenResp, nil
}
