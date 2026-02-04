package validators

import (
	"errors"
	"myapi/internal/models"
)

func ValidateCategoria(categoria *models.Categoria) error {
	if len(categoria.Codigo) != 4 {
		return errors.New("código precisa ter 4 caracteres")
	}
	return nil
}
