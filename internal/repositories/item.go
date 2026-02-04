package repositories

import (
	"myapi/internal/config"
	"myapi/internal/models"
)

type ItemRepository struct{}

func NewItemRepository() *ItemRepository {
	return &ItemRepository{}
}

func (r *ItemRepository) ListAll() ([]models.Item, error) {
	var itens []models.Item
	if err := config.DB.Find(&itens).Error; err != nil {
		return nil, err
	}
	return itens, nil
}

func (r *ItemRepository) FindByID(id int) (*models.Item, error) {
	var item models.Item
	if err := config.DB.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *ItemRepository) FindByCodigo(cod string) (*models.Item, error) {
	var item models.Item
	// Busca o item onde o campo "codigo" é igual ao valor fornecido
	if err := config.DB.Where("codigo = ?", cod).First(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *ItemRepository) Create(item *models.Item) (*models.Item, error) {
	if err := config.DB.Create(&item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func (r *ItemRepository) Update(item *models.Item) error {
	if err := config.DB.Save(&item).Error; err != nil {
		return err
	}
	return nil
}

func (r *ItemRepository) Delete(id int) error {
	if err := config.DB.Delete(&models.Item{}, id).Error; err != nil {
		return err
	}
	return nil
}
