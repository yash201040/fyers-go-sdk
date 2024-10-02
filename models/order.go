package models

// OrderRequest represents a request to place an order.
type OrderRequest struct {
	Symbol       string  `json:"symbol"`
	Qty          int     `json:"qty"`
	Type         int     `json:"type"`
	Side         int     `json:"side"`
	ProductType  string  `json:"productType"`
	LimitPrice   float64 `json:"limitPrice,omitempty"`
	StopPrice    float64 `json:"stopPrice,omitempty"`
	DisclosedQty int     `json:"disclosedQty,omitempty"`
	Validity     string  `json:"validity"`
	OfflineOrder bool    `json:"offlineOrder,omitempty"`
}

// OrderResponse represents a response from the API after placing an order.
type OrderResponse struct {
	OrderID string `json:"id"`
	Status  string `json:"status"`
}
