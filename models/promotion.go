package models

import "database/sql"

type Promotion struct {
	MinQty         int             `json:"min_qty"`
	Type           string          `json:"typ"`
	UOM            sql.NullString  `json:"uom"`
	BonusItemSKU   sql.NullString  `json:"bonus_item_sku"`
	BonusItemPrice sql.NullFloat64 `json:"bonus_item_price"`
	Value          float64         `json:"value"`
}
