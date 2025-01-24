package helper

import (
	"net/http"
	"simpleBank/src/pkg/service_errors"
)

var StatusCodeMapping = map[string]int{
	service_errors.AccountNotFoundErr:     service_errors.AccountNotfoundErCode,
	service_errors.InvalidStatusErr:       service_errors.InvalidTransactionStatusErrCode,
	service_errors.InsufficientBalanceErr: service_errors.InsufficientBalanceErrCode,
}

// StatusCodeToHttpMapping internal status to http
var StatusCodeToHttpMapping = map[int]int{
	//
	service_errors.AccountNotfoundErCode:           http.StatusNotFound,
	service_errors.InvalidTransactionStatusErrCode: http.StatusBadRequest,
	service_errors.InsufficientBalanceErrCode:      http.StatusBadRequest,
}

func TranslateErrorToStatusCode(err error) int {
	value, ok := StatusCodeMapping[err.Error()]
	if !ok {
		return service_errors.BadRequestErrCode
	}
	return value
}

func TranslateStatusCodeToErrorString(status int) int {
	status, ok := StatusCodeToHttpMapping[status]
	if !ok {
		return http.StatusBadRequest
	}
	return status
}
