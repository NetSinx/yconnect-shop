package domain

import "github.com/NetSinx/yconnect-shop/server/user/model/entity"

type MessageResp struct {
	Message string `json:"message"`
}

type RespData struct {
	Data   entity.User `json:"data"`
}
