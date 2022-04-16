package controllers

var (
	getInventorySQLQuery = `
	SELECT 
		sku,
		name,
		price,
		qty
	FROM inventories WHERE sku = ? FOR UPDATE;
	`

	getPromotionDiscountSQLQuery = `
	SELECT 
		min_qty,
		typ,
		uom, 
		inventories.sku AS bonus_item_sku,
		inventories.price AS bonus_item_price,
		value
	FROM
		promotions
        LEFT JOIN inventories ON uom = sku
	WHERE
		item_sku = ?
		AND min_qty <= ?
		AND is_active = 1
	ORDER BY promotions.created_at DESC
	LIMIT 1;
	`

	updateInventorySQLQuery = `
	UPDATE inventories
		SET qty = qty - ?,
		updated_at = CURRENT_TIMESTAMP(6)
	WHERE
		sku = ?
	`
)
