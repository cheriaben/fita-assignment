package rest

import (
	"encoding/json"
	"fita-assignment/models"
	"fmt"
	"net/http"
	"time"
)

func HTTPRespError(w http.ResponseWriter, r *http.Request, statusCode int, message string) {

	jsonErrResp := &HTTPErrResp{
		Meta: Meta{
			Path:      r.URL.String(),
			Message:   fmt.Sprintf("%s %s [%d] %s", r.Method, r.URL.RequestURI(), statusCode, message),
			Timestamp: time.Now().Format(time.RFC3339),
		},
	}

	raw, err := json.Marshal(jsonErrResp)
	if err != nil {
		statusCode = http.StatusInternalServerError
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, _ = w.Write(raw)
}

func HTTPRespSuccess(w http.ResponseWriter, r *http.Request, statusCode int, resp interface{}) {

	meta := Meta{
		Path:      r.URL.String(),
		Message:   fmt.Sprintf("%s %s [%d] %s", r.Method, r.URL.RequestURI(), statusCode, http.StatusText(statusCode)),
		Timestamp: time.Now().Format(time.RFC3339),
	}

	var (
		raw []byte
		err error
	)

	switch data := resp.(type) {
	case models.CheckoutResult:
		result := HTTPCheckoutResultResponse{
			Meta: meta,
			Data: HTTPCheckoutResultDataResponse{
				Checkout: data,
			},
		}
		raw, err = json.Marshal(&result)
	default:
		HTTPRespError(w, r, http.StatusInternalServerError, "Failed on cast data")
		return
	}

	if err != nil {
		HTTPRespError(w, r, http.StatusInternalServerError, "Failed on marshal")
		return
	}

	w.Header().Set("httpheader.ContentType", "httpheader.ContentJSON")
	w.WriteHeader(statusCode)
	_, _ = w.Write(raw)
}
