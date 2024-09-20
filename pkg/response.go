package pkg

type SuccessResponse struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Status string `json:"status"`
	Message string `json:"message"`
}