package repository

import (
	"github.com/NetSinx/yconnect-shop/cart/model"
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
	if err := c.db.Find(&cart).Error; err != nil {
		return nil, err
	}

	return cart, nil
}

func (c cartRepository) CreateCart(cart model.Cart) (model.Cart, error) {
	if err := c.db.Create(&cart).Error; err != nil {
		return cart, err
	}

	return cart, nil
}

func (c cartRepository) UpdateCart(cart model.Cart, id string) (model.Cart, error) {
	if err := c.db.Where("id = ?", id).Updates(&cart).Error; err != nil {
		return cart, err
	}
	
	if err := c.db.First(&cart, "id = ?", id).Error; err != nil {
		return cart, err
	}

	return cart, nil
}

func (c cartRepository) DeleteCart(cart model.Cart, id string) error {
	if err := c.db.Delete(&cart, "id = ?", id).Error; err != nil {
		return err
	}

	return nil
}

func (c cartRepository) GetCartById(cart model.Cart, id string) (model.Cart, error) {
	if err := c.db.First(&cart, "id = ?", id).Error; err != nil {
		return cart, err
	}

	return cart, nil
}

func (c cartRepository) GetCartBySlug(cart model.Cart, slug string) (model.Cart, error) {
	if err := c.db.First(&cart, "slug = ?", slug).Error; err != nil {
		return cart, err
	}

	return cart, nil
}