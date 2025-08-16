package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"github.com/NetSinx/yconnect-shop/server/order/model/domain"
	"github.com/NetSinx/yconnect-shop/server/order/model/entity"
	"github.com/NetSinx/yconnect-shop/server/order/rabbitmq"
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


func (os *orderService) GetOrder(order []entity.Order, username string) ([]entity.Order, error) {
	respUser, err := http.Get("http://user-service:8082/user/")

	getOrder, err := os.orderRepo.GetOrder(order, user_id)
	if err != nil {
		return order, err
	}

	for _, order := range getOrder {
		rabbitmq.GetProductByID(order.ProductID)
		product, _ := rabbitmq.ConsumeProduct()
		order.Product = product
	}

	return getOrder, nil
}

func (os *orderService) AddOrder(reqOrder domain.OrderRequest) error {
	var order entity.Order

	if err := validator.New().Struct(reqOrder); err != nil {
		return err
	}

	respProduct, err := http.Get(fmt.Sprintf("http://localhost:8081/product/%d", reqOrder.ProductID))
	if err != nil || respProduct.StatusCode != 200 {
		return errors.New("produk tidak ditemukan")
	}

	var respData domain.DataProduct

	json.NewDecoder(respProduct.Body).Decode(&respData)

	order.Product = respData.Data
	order.Kuantitas = reqOrder.Kuantitas
	order.Status = "Sedang Diproses"
	order.Estimasi = time.Now().Add(72 * time.Hour)

	if err := os.orderRepo.AddOrder(order); err != nil {
		return err
	}

	return nil
}

func (os *orderService) DeleteOrder(order entity.Order, username string) error {
	rabbitmq.GetUserID(username)

	user_id, err := rabbitmq.ConsumeUserID()
	if err != nil {
		return err
	}

	err = os.orderRepo.DeleteOrder(order, username, user_id)
	if err != nil && err == gorm.ErrRecordNotFound {
		return fmt.Errorf("pesanan tidak ditemukan")
	} else if err != nil {
		return err
	}

	return nil
}