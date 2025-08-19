package dto

type RespData struct {
	Data   any `json:"data"`
}

type MessageResp struct {
	Message string `json:"message"`
}

type ResponseCSRF struct {
	CSRFToken string `json:"csrf_token"`
}