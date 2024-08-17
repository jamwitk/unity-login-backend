package models

type ErrorException struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
