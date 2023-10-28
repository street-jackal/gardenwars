package responses

type BaseResponse struct {
	Status  int    `json:"status"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}
