package service_errors

const (
	InsufficientBalanceErr = "insufficient balance"
	AccountNotFoundErr     = "sheba not found"
	InvalidStatusErr       = "invalid status"
)

const (
	BadRequestErrCode               int = 4000
	InsufficientBalanceErrCode      int = 4001
	AccountNotfoundErCode           int = 4002
	InvalidTransactionStatusErrCode int = 4003
)
