package models

// AuthRequest represents the request body for generating an auth code.
type AuthRequest struct {
	ClientID     string `json:"client_id"`
	RedirectURI  string `json:"redirect_uri"`
	ResponseType string `json:"response_type"`
	State        string `json:"state"`
}

// AuthResponse represents the response from the API for auth code generation.
type AuthResponse struct {
	Code  string `json:"code"`
	State string `json:"state"`
}

// TokenRequest represents the payload for the access token request.
type TokenRequest struct {
	ClientID    string `json:"client_id"`
	SecretID    string `json:"secret_id"`
	GrantType   string `json:"grant_type"`
	Code        string `json:"code"`
	RedirectURI string `json:"redirect_uri"`
}

// TokenResponse represents the response structure from the token endpoint.
type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Error       string `json:"error"`
}
