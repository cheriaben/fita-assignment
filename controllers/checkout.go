package controllers

import (
	"context"
	"fita-assignment/database"
	"fita-assignment/models"
	"math"

	log "github.com/sirupsen/logrus"
)

func Checkout(ctx context.Context, request []models.Item) (models.CheckoutResult, error) {

	var inv models.Inventory

	var result models.CheckoutResult

	scannedItems := make(map[string]int)
	possiblePromoItems := make(map[string]models.Item)

	tx, err := database.DBConn.BeginTx(ctx, nil)
	if err != nil {
		log.Error(err)
		return result, err
	}

	for _, v := range request {
		var promotion models.Promotion
		rows, err := tx.Query(getInventorySQLQuery, v.SKU)
		if err != nil {
			tx.Rollback()
			log.Error(err)
			return result, err
		}

		defer rows.Close()

		for rows.Next() {
			err := rows.Scan(
				&inv.SKU,
				&inv.Name,
				&inv.Price,
				&inv.Qty,
			)
			if err != nil {
				tx.Rollback()
				log.Error(err)
				return result, err
			}
		}

		rows, err = tx.Query(getPromotionDiscountSQLQuery, v.SKU, v.Qty)
		if err != nil {
			log.Error(err)
			return result, err
		}
		defer rows.Close()

		for rows.Next() {
			err := rows.Scan(
				&promotion.MinQty,
				&promotion.Type,
				&promotion.UOM,
				&promotion.BonusItemSKU,
				&promotion.BonusItemPrice,
				&promotion.Value,
			)
			if err != nil {
				tx.Rollback()
				log.Error(err)
				return result, err
			}
		}

		scannedItems[v.SKU] = v.Qty
		subtotal := float64(v.Qty) * inv.Price
		result.FinalPrice = result.FinalPrice + subtotal

		switch promotion.Type {
		case "DISCOUNT":
			result.FinalPrice = result.FinalPrice - (subtotal * promotion.Value / 100)
		case "BOGO_SCHEME":
			disc := inv.Price * (math.Floor(float64(v.Qty / promotion.MinQty)))
			result.FinalPrice = math.Round((result.FinalPrice-disc)*100) / 100
		case "FREE_ITEM":
			if val, ok := scannedItems[promotion.BonusItemSKU.String]; ok {
				disc := promotion.BonusItemPrice.Float64 * math.Min(float64(val), float64(v.Qty))
				result.FinalPrice = result.FinalPrice - disc
			} else {
				possiblePromoItems[promotion.BonusItemSKU.String] = models.Item{
					SKU: promotion.BonusItemSKU.String,
					Qty: v.Qty,
				}
			}
		}

		if val, ok := possiblePromoItems[v.SKU]; ok {
			result.FinalPrice = result.FinalPrice - (math.Min(float64(val.Qty), float64(v.Qty)) * inv.Price)
		}

		_, err = tx.Exec(updateInventorySQLQuery, v.Qty, v.SKU)
		if err != nil {
			tx.Rollback()
			log.Error(err)
			return result, err
		}
	}

	tx.Commit()

	return result, nil

}
