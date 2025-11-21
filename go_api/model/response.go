package model

type Response struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}
