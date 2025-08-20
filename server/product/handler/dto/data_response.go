package dto

var (
	CreateResponse = "Produk berhasil ditambahkan."
	UpdateResponse = "Produk berhasil diubah."
	DeleteResponse = "Produk berhasil dihapus."
)

type RespData struct {
	Data   any `json:"data"`
}

type MessageResp struct {
	Message string `json:"message"`
}

type ResponseCSRF struct {
	CSRFToken string `json:"csrf_token"`
}