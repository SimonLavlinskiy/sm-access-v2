package models

type ErrorResponse struct {
	ErrorCode    int     `json:"errorCode"`
	ErrorMessage *string `json:"errorMessage"`
}
