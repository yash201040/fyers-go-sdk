package main

import (
	"errors"
	"fmt"
	"fyers-go-sdk/api"
	"fyers-go-sdk/utils"
	"os"
)

func main() {
	logger, err := utils.InitializeLogger("fyers.log")
	if err != nil {
		fmt.Printf("ERROR: Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Info("Application Started")

	// Generate Auth URL
	appID := "C26U8H9DO5-100"
	redirectURI := "your-redirect-uri"
	state := "random-state"
	authURL := api.GenerateAuthURL(appID, redirectURI, state)
	fmt.Printf("Visit this URL to authorize the app: %s\n", authURL)

	// Wait for user to input the authorization code received after login
	var authCode string
	fmt.Print("Enter the authorization code: ")
	fmt.Scan(&authCode)

	// Exchange auth code for access token
	secretID := "your-secret-id"
	tokenResp, err := api.GetAccessToken(appID, secretID, authCode, redirectURI)
	if err != nil {
		logger.Error("Failed to fetch access token: %v", err)
		return
	}
	logger.Info("Access token fetched successfully: %+v", tokenResp.AccessToken)

	// Use your live credentials here
	client := utils.NewClient(appID, tokenResp.AccessToken, "https://api.fyers.in")

	// Fetch Profile
	profile, err := api.GetProfile(client, logger)
	if err != nil {
		if errors.Is(err, os.ErrPermission) {
			logger.Error("Permission error while fetching profile: %v", err)
		} else {
			logger.Error("Failed to fetch profile: %v", err)
		}
		return
	}
	logger.Info("Fetched profile: %+v", profile)

	// Fetch Funds
	funds, err := api.GetFunds(client, logger)
	if err != nil {
		logger.Error("Failed to fetch funds: %v", err)
		return
	}
	logger.Info("Fetched funds: %+v", funds)

	// Fetch Market Quotes
	symbol := "NSE:SBIN-EQ"
	quotes, err := api.GetQuotes(client, symbol, logger)
	if err != nil {
		logger.Error("Failed to fetch quotes for symbol: %v, error: %v", symbol, err)
		return
	}
	logger.Info("Fetched quotes for symbol %v: %+v", symbol, quotes)

	// Place Order
	// order := models.OrderRequest{
	// 	Symbol:      "NSE:IDEA-EQ",
	// 	Qty:         1,
	// 	Type:        1, // Limit order
	// 	Side:        1, // Buy
	// 	ProductType: "INTRADAY",
	// 	LimitPrice:  7.9,
	// 	Validity:    "DAY",
	// }

	// orderResp, err := api.PlaceLimitOrder(client, order)
	// if err != nil {
	// 	logger.Error("Failed to place limit order: %v", err)
	// 	return
	// }
	// logger.Info("Order placed successfully: %+v", orderResp)

	logger.Info("Application Finished")
}
