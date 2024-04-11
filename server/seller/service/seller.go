package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/NetSinx/yconnect-shop/server/seller/model/domain"
	"github.com/NetSinx/yconnect-shop/server/seller/model/entity"
	"github.com/NetSinx/yconnect-shop/server/seller/repository"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
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
	var respGetProduct domain.GetProductResponse

	listSeller, err := ss.SellerRepository.ListSeller()
	if err != nil {
		return []entity.Seller{}, fmt.Errorf("seller tidak ditemukan")
	}

	for i := range listSeller {
		resp, err := http.Get(fmt.Sprintf("http://product-service:8081/product/seller/%d", listSeller[i].Id))
		if err != nil || resp.StatusCode != 200 {
			return listSeller, nil
		}

		json.NewDecoder(resp.Body).Decode(&respGetProduct)

		listSeller[i].Product = respGetProduct.Data
	}

	return listSeller, nil
}

func (ss sellerService) RegisterSeller(username string) (entity.Seller, error) {
	var sellerValidity entity.Seller
	var respUser domain.GetUserResponse

	if err := validator.New().Struct(&sellerValidity); err != nil {
		return entity.Seller{}, err
	}

	resp, err := http.Get(fmt.Sprintf("http://localhost:8082/user/%s", username))
	if err != nil || resp.StatusCode != 200 {
		return entity.Seller{}, fmt.Errorf("seller gagal registrasi. user tidak ditemukan")
	}

	json.NewDecoder(resp.Body).Decode(&respUser)

	seller := entity.Seller{
		Name: sellerValidity.Name,
		Avatar: respUser.Data.Avatar,
		Username: respUser.Data.Username,
		Email: respUser.Data.Email,
		Alamat: respUser.Data.Alamat,
		NoTelp: respUser.Data.NoTelp,
		Product: respUser.Data.Seller.Product,
		UserID: respUser.Data.Id,
	}

	regSeller, err := ss.SellerRepository.RegisterSeller(seller)
	if err != nil {
		return entity.Seller{}, fmt.Errorf("seller sudah terdaftar")
	}

	return regSeller, nil
}

func (ss sellerService) UpdateSeller(username string) (entity.Seller, error) {
	var userResp domain.GetUserResponse

	respUpdSeller, err := http.Get(fmt.Sprintf("http://localhost:8082/user/%s", username))
	if err != nil {
		return entity.Seller{}, fmt.Errorf("service user sedang bermasalah")
	} else if respUpdSeller.StatusCode != 200 {
		return entity.Seller{}, fmt.Errorf("seller tidak ditemukan")
	}

	json.NewDecoder(respUpdSeller.Body).Decode(&userResp)

	seller := entity.Seller{
		Name: userResp.Data.Name,
		Avatar: userResp.Data.Avatar,
		Username: userResp.Data.Username,
		Email: userResp.Data.Email,
		Alamat: userResp.Data.Alamat,
		NoTelp: userResp.Data.NoTelp,
		Product: userResp.Data.Seller.Product,
		UserID: userResp.Data.Id,
	}

	updSeller, err := ss.SellerRepository.UpdateSeller(username, seller)
	if err == gorm.ErrDuplicatedKey {
		return entity.Seller{}, fmt.Errorf("seller sudah terdaftar")
	} else if err == gorm.ErrRecordNotFound {
		return entity.Seller{}, fmt.Errorf("seller tidak ditemukan")
	}

	return updSeller, nil
}

func (ss sellerService) DeleteSeller(username string) error {
	if err := ss.SellerRepository.DeleteSeller(username); err != nil {
		return fmt.Errorf("seller tidak ditemukan")
	}

	return nil
}

func (ss sellerService) GetSeller(username string) (entity.Seller, error) {
	var respGetProduct domain.GetProductResponse

	getSeller, err := ss.SellerRepository.GetSeller(username)
	if err != nil {
		return entity.Seller{}, fmt.Errorf("seller tidak ditemukan")
	}

	resp, err := http.Get(fmt.Sprintf("http://product-service:8081/product/seller/%d", getSeller.Id))
	if err != nil || resp.StatusCode != 200 {
		return getSeller, nil
	}

	json.NewDecoder(resp.Body).Decode(&respGetProduct)

	getSeller.Product = respGetProduct.Data

	return getSeller, nil
}