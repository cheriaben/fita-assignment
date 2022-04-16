package rest

import (
	"encoding/json"
	c "fita-assignment/controllers"
	"io/ioutil"
	"net/http"
)

func CheckOut(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		HTTPRespError(w, r, http.StatusBadRequest, "Error on read body")
		return
	}

	var requestBody HTTPCheckoutRequest
	if err := json.Unmarshal(body, &requestBody); err != nil {
		HTTPRespError(w, r, http.StatusBadRequest, "Error on unmarshal")
		return
	}

	result, err := c.Checkout(r.Context(), requestBody.Data.CheckoutItems)
	if err != nil {
		HTTPRespError(w, r, http.StatusBadRequest, "Error checkout")
		return
	}

	HTTPRespSuccess(w, r, http.StatusOK, result)
}
