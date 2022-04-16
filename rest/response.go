package rest

import (
	"fita-assignment/models"
)

type Meta struct {
	Path      string `json:"path"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

type HTTPErrResp struct {
	Meta Meta `json:"metadata"`
}

type HTTPCheckoutResultResponse struct {
	Data HTTPCheckoutResultDataResponse `json:"data"`
	Meta Meta                           `json:"metadata"`
}

type HTTPCheckoutResultDataResponse struct {
	Checkout models.CheckoutResult `json:"checkout"`
}
