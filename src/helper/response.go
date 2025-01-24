package helper

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type SuccessResponse struct {
	Message string `json:"message"`
	Result  any    `json:"data"`
}

func SendErrorResponse(w http.ResponseWriter, err error) {
	code := TranslateErrorToStatusCode(err)
	httpStatusCode := TranslateStatusCodeToErrorString(code)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)

	response := ErrorResponse{
		Message: err.Error(),
		Code:    code,
	}

	json.NewEncoder(w).Encode(response)
}

func SendSuccessResponse(w http.ResponseWriter, r any, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := SuccessResponse{
		Message: msg,
		Result:  r,
	}

	json.NewEncoder(w).Encode(response)
}
