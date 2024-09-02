package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/NetSinx/yconnect-shop/server/order/model/domain"
	"github.com/NetSinx/yconnect-shop/server/order/model/entity"
	"github.com/NetSinx/yconnect-shop/server/order/repository"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type orderService struct {
	orderRepo repository.OrderRepository
}

func OrderServ(orderRepo repository.OrderRepository) *orderService {
	return &orderService{
		orderRepo: orderRepo,
	}
}

func (os *orderService) ListOrder(order []entity.Order, username string) ([]entity.Order, error) {
	listOrder, err := os.orderRepo.ListOrder(order, username)
	if err != nil {
		return order, err
	}

	for _, o := range listOrder {
		respProduct, err := http.Get(fmt.Sprintf("http://localhost:8081/product/%d", o.ProductID))
		if err != nil || respProduct.StatusCode != 200 {
			return listOrder, nil
		}
	
		var respData domain.DataProduct
	
		json.NewDecoder(respProduct.Body).Decode(&respData)
	
		o.Product = respData.Data
	}
	
	return listOrder, nil
}

func (os *orderService) AddOrder(order entity.Order) error {
	if err := validator.New().Struct(order); err != nil {
		return err
	}

	if err := os.orderRepo.AddOrder(order); err != nil {
		return err
	}

	return nil
}

func (os *orderService) DeleteOrder(order entity.Order, username, id string) error {
	err := os.orderRepo.DeleteOrder(order, username, id)
	if err != nil && err == gorm.ErrRecordNotFound {
		return fmt.Errorf("pesanan tidak ditemukan")
	} else if err != nil {
		return err
	}

	return nil
}