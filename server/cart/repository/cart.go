package repository

import (
	"github.com/NetSinx/yconnect-shop/server/cart/model"
	"gorm.io/gorm"
)

type cartRepository struct {
	db *gorm.DB
}

func CartRepository(db *gorm.DB) cartRepository {
	return cartRepository{
		db: db,
	}
}

func (c cartRepository) ListCart(cart []model.Cart) ([]model.Cart, error) {
	if err := c.db.Select("id", "product_id", "item", "user_id").Find(&cart).Error; err != nil {
		return nil, err
	}

	return cart, nil
}

func (c cartRepository) AddToCart(cart model.Cart) (model.Cart, error) {
	if err := c.db.Create(&cart).Error; err != nil {
		return cart, err
	}

	return cart, nil
}

func (c cartRepository) UpdateCart(cart model.Cart, id uint) (model.Cart, error) {
	err := c.db.Where("id = ?", id).Updates(&cart).Error
	if err != nil {
		return cart, err
	}

	if err := c.db.First(&cart, "id = ?", id).Error; err != nil {
		return cart, err
	}

	return cart, nil
}

func (c cartRepository) DeleteProductInCart(cart model.Cart, id uint) error {
	if err := c.db.First(&cart, "id = ?", id).Error; err != nil {
		return err
	}

	if err := c.db.Delete(&cart, "id = ?", id).Error; err != nil {
		return err
	}

	return nil
}

func (c cartRepository) GetCart(cart model.Cart, id uint) (model.Cart, error) {
	if err := c.db.First(&cart, "id = ?", id).Error; err != nil {
		return cart, err
	}

	return cart, nil
}

func (c cartRepository) GetCartByUser(cart []model.Cart, id uint) ([]model.Cart, error) {
	if err := c.db.Where("user_id = ?", id).Find(&cart).Error; err != nil {
		return nil, err
	}

	return cart, nil
}
