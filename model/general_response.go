package model

type GeneralResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type WebResponse struct {
	IsSuccessful bool   `json:"isSuccessful"`
	StatusCode   string `json:"statusCode"`
	Message      string `json:"message"`
}
