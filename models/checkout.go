package models

type Item struct {
	SKU string `json:"sku"`
	Qty int    `json:"qty"`
}

type CheckoutResult struct {
	FinalPrice float64 `json:"final_price"`
}
