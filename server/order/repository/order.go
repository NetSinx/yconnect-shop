package repository

import (
	"github.com/NetSinx/yconnect-shop/server/order/model/entity"
	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func OrderRepo(db *gorm.DB) *orderRepository {
	return &orderRepository{
		db: db,
	}
}

func (or *orderRepository) ListOrder(order []entity.Order) []entity.Order {
	or.db.Preload("Products").Find(&order)

	return order
}

func (or *orderRepository) AddOrder(order entity.Order) error {
	if err := or.db.Create(&order).Error; err != nil {
		return err
	}

	return nil
}

func (or *orderRepository) GetOrder(order entity.Order, id uint) (entity.Order, error) {
	if err := or.db.First(&order, "id = ?", id).Error; err != nil {
		return entity.Order{}, err
	}

	return order, nil
}

func (or *orderRepository) DeleteOrder(order entity.Order, id uint) error {
	if err := or.db.First(&order, "id = ?", id).Error; err != nil {
		return err
	}

	if err := or.db.Delete(&order, "id = ?", id).Error; err != nil {
		return err
	}

	return nil
}