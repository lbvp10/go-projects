package logger

import (
	"net/http"
)

type CauseList []interface{}

type ErrorApi struct {
	Message string    `json:"message"`
	Cause   CauseList `json:"cause"`
	Code    string    `json:"code"`
	Status  int       `json:"status"`
}

func NewApiError(message string, error string, status int, cause CauseList, logger Logger) ErrorApi {
	errorApi := ErrorApi{Message: message, Status: status, Cause: cause}
	logger.ErrorF(errorApi.Message, error)
	return errorApi
}

func NewNotFoundApiError(message string) ErrorApi {
	errorApi := ErrorApi{Message: message, Status: http.StatusNotFound, Code: "not_found"}
	return errorApi
}

func NewBadRequestApiError(message string) ErrorApi {
	errorApi := ErrorApi{Message: message, Status: http.StatusBadRequest, Code: "bad_request"}
	return errorApi
}

func NewValidationApiError(message string, error string, cause CauseList, logger Logger) ErrorApi {
	errorApi := ErrorApi{Message: message, Status: http.StatusConflict, Code: "conflict"}
	logger.ErrorF(errorApi.Message, error)
	return errorApi
}

func NewInternalServerApiError(err error, code string, logger Logger) ErrorApi {
	cause := CauseList{}
	if err != nil {
		cause = append(cause, err.Error())
	}
	errorApi := ErrorApi{Message: "internal_server_error", Status: http.StatusInternalServerError, Cause: cause}
	logger.ErrorF(errorApi.Message, errorApi)
	return errorApi
}
