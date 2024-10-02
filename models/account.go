package models

// ProfileResponse represents the response from the API for fetching account profile.
type ProfileResponse struct {
	Data struct {
		ClientID string `json:"client_id"`
		Name     string `json:"name"`
	} `json:"data"`
	Status string `json:"status"`
}

// FundsResponse represents the response from the API for fetching account funds.
type FundsResponse struct {
	Data struct {
		Balance float64 `json:"balance"`
	} `json:"data"`
	Status string `json:"status"`
}
