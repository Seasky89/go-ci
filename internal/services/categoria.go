package services

import (
	"encoding/json"
	"errors"
	"myapi/internal/models"
	"myapi/internal/validators"
	"net/http"
)

func DecodeAndValidateCategoria(r *http.Request) (*models.Categoria, error) {
	var categoria models.Categoria
	if err := json.NewDecoder(r.Body).Decode(&categoria); err != nil {
		return nil, errors.New("Erro ao decodificar a categoria")
	}
	err := validators.ValidateCategoria(&categoria)
	if err != nil {
		return nil, errors.New("Erro de validação: " + err.Error())
	}
	return &categoria, nil
}
