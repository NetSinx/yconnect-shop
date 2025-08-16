package domain

type RespData struct {
	Data   interface{} `json:"data"`
}

type MessageResp struct {
	Message string `json:"message"`
}

type ResponseCSRF struct {
	CSRFToken string `json:"csrf_token"`
}