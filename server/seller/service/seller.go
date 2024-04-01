package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"github.com/NetSinx/yconnect-shop/server/seller/model/domain"
	"github.com/NetSinx/yconnect-shop/server/seller/model/entity"
	"github.com/NetSinx/yconnect-shop/server/seller/repository"
)

type sellerService struct {
	SellerRepository repository.SellerRepo
}

func SellerService(sr repository.SellerRepo) sellerService {
	return sellerService{
		SellerRepository: sr,
	}
}

func (ss sellerService) ListSeller() ([]entity.Seller, error) {
	var respListSeller domain.GetProductResponse

	listSeller, err := ss.SellerRepository.ListSeller()
	if err != nil {
		return []entity.Seller{}, errors.New("seller tidak ditemukan")
	}

	for i := range listSeller {
		resp, _ := http.Get(fmt.Sprintf("http://product-service:8081/product/seller/%d", listSeller[i].Id))

		json.NewDecoder(resp.Body).Decode(&respListSeller)

		listSeller[i].Product = respListSeller.Data
	}

	return listSeller, nil
}