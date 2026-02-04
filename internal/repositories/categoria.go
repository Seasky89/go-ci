package repositories

import (
	"myapi/internal/config"
	"myapi/internal/models"
)

type CategoriaRepository struct{}

func NewCategoriaRepository() *CategoriaRepository {
	return &CategoriaRepository{}
}

func (r *CategoriaRepository) ListAll() ([]models.Categoria, error) {
	var categorias []models.Categoria
	if err := config.DB.Find(&categorias).Error; err != nil {
		return nil, err
	}
	return categorias, nil
}

func (r *CategoriaRepository) FindByID(id int) (*models.Categoria, error) {
	var categoria models.Categoria
	if err := config.DB.First(&categoria, id).Error; err != nil {
		return nil, err
	}
	return &categoria, nil
}

func (r *CategoriaRepository) Create(categoria *models.Categoria) (*models.Categoria, error) {
	if err := config.DB.Create(&categoria).Error; err != nil {
		return nil, err
	}
}

func (r *CategoriaRepository) Update(categoria *models.Categoria) error {
	if err := config.DB.Save(&categoria).Error; err != nil {
		return err
	}
	return nil
}

func (r *CategoriaRepository) Delete(id int) error {
	if err := config.DB.Delete(&models.Categoria{}, id).Error; err != nil {
		return err
	}
	return nil
}
