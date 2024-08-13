package domain

type MessageResp struct {
	Message string `json:"message"`
}

type DataResp struct {
	Data interface{} `json:"data"`
}