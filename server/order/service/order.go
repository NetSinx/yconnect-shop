package service

import (
	"github.com/NetSinx/yconnect-shop/server/order/model/entity"
	"github.com/NetSinx/yconnect-shop/server/order/repository"
	"github.com/go-playground/validator/v10"
)

type orderService struct {
	orderRepo repository.OrderRepository
}

func OrderServ(orderRepo repository.OrderRepository) *orderService {
	return &orderService{
		orderRepo: orderRepo,
	}
}

func (os *orderService) ListOrder(order []entity.Order) []entity.Order {
	return os.orderRepo.ListOrder(order)
}

func (os *orderService) AddOrder(order entity.Order) error {
	if err := validator.New().Struct(&order); err != nil {
		return err
	}

	if err := os.orderRepo.AddOrder(order); err != nil {
		return err
	}

	return nil
}

func (os *orderService) GetOrder(order entity.Order, id uint) (entity.Order, error) {
	getOrder, err := os.orderRepo.GetOrder(order, id)
	if err != nil {
		return entity.Order{}, err
	}

	return getOrder, nil
}

func (os *orderService) DeleteOrder(order entity.Order, id uint) error {
	if err := os.orderRepo.DeleteOrder(order, id); err != nil {
		return err
	}

	return nil
}