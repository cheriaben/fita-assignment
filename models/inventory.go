package models

type Inventory struct {
	SKU   string  `json:"item_sku"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Qty   int     `json:"qty"`
}
