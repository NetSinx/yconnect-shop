package converter

import (
	"github.com/NetSinx/yconnect-shop/server/product/internal/entity"
	"github.com/NetSinx/yconnect-shop/server/product/internal/model"
)

func ProductToResponse(product *entity.Product) *model.ProductResponse {
	gambarResponse := make([]model.GambarResponse, len(product.Gambar))
	for i, gambar := range product.Gambar {
		gambarResponse[i].ID = gambar.ID
		gambarResponse[i].Path = gambar.Path
		gambarResponse[i].CreatedAt = gambar.CreatedAt
		gambarResponse[i].UpdatedAt = gambar.UpdatedAt
	}

	return &model.ProductResponse{
		ID: product.ID,
		Nama: product.Nama,
		Slug: product.Slug,
		Gambar: gambarResponse,
		Deskripsi: product.Deskripsi,
		KategoriSlug: product.KategoriSlug,
		Harga: product.Harga,
		Stok: product.Stok,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
}

func ProductToEvent(product *entity.Product) *model.CategoryEvent {
	return &model.CategoryEvent{
		ID: product.ID,
		Nama: product.Nama,
		Slug: product.Slug,
	}
}
