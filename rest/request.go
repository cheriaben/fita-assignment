package rest

import "fita-assignment/models"

type HTTPCheckoutRequest struct {
	Data HTTPCheckoutDataRequest `json:"data"`
}

type HTTPCheckoutDataRequest struct {
	CheckoutItems []models.Item `json:"checkout_items"`
}
