package repository

import (
	"github.com/NetSinx/yconnect-shop/server/order/model/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type orderRepository struct {
	db *gorm.DB
}

func OrderRepo(db *gorm.DB) *orderRepository {
	return &orderRepository{
		db: db,
	}
}

func (or *orderRepository) ListOrder(order []entity.Order, username string) ([]entity.Order, error) {
	if err := or.db.Find(&order, "username = ?", username).Error; err != nil {
		return order, err
	}

	return order, nil
}

func (or *orderRepository) AddOrder(order entity.Order) error {
	if err := or.db.Clauses(clause.OnConflict{DoNothing: false}).Create(&order).Error; err != nil {
		return err
	}

	return nil
}

func (or *orderRepository) DeleteOrder(order entity.Order, username, id string) error {
	if err := or.db.First(&order, "username = ? AND id = ?", username, id).Error; err != nil {
		return err
	}

	if err := or.db.Delete(&order).Error; err != nil {
		return err
	}

	return nil
}