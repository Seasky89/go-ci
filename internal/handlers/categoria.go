package handlers

import (
	"encoding/json"
	"myapi/internal/models"
	"myapi/internal/repositories"
	"myapi/internal/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ==================== HANDLERS PARA CATEGORIAS ====================

// Listar todas as categorias
func ListCategorias(w http.ResponseWriter, r *http.Request) {
	repository := repositories.NewCategoriaRepository()
	categorias, err := repository.ListAll()
	if err != nil {
		utils.RespondWithError(w, "Erro ao buscar categorias", http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode(categorias); err != nil {
		utils.RespondWithError(w, "Erro ao codificar as categorias", http.StatusInternalServerError)
		return
	}
}

// Buscar uma única categoria pelo id (via query string: ?id=1)
func GetCategoria(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		utils.RespondWithError(w, "ID não fornecido", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(w, "ID inválido", http.StatusBadRequest)
		return
	}
	repository := repositories.NewCategoriaRepository()
	categoria, err := repository.FindByID(id)
	if err != nil {
		utils.RespondWithError(w, "Categoria não encontrada", http.StatusNotFound)
		return
	}
	if err = json.NewEncoder(w).Encode(categoria); err != nil {
		utils.RespondWithError(w, "Erro ao codificar as categorias", http.StatusInternalServerError)
		return
	}

}

// Criar uma nova categoria (envie JSON via POST)
func CreateCategoria(w http.ResponseWriter, r *http.Request) {
	var categoria models.Categoria
	if err := json.NewDecoder(r.Body).Decode(&categoria); err != nil {
		utils.RespondWithError(w, "Erro ao decodificar a categoria", http.StatusBadRequest)
		return
	}
	repository := repositories.NewCategoriaRepository()
	if err := repository.Create(categoria); err != nil {
		utils.RespondWithError(w, "Erro ao criar a categoria", http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(categoria); err != nil {
		utils.RespondWithError(w, "Erro ao codificar a categoria criada", http.StatusInternalServerError)
		return
	}
}

// Atualizar uma categoria (envie JSON via PUT, com o campo id preenchido)
func UpdateCategoria(w http.ResponseWriter, r *http.Request) {
	var categoria models.Categoria
	if err := json.NewDecoder(r.Body).Decode(&categoria); err != nil {
		utils.RespondWithError(w, "Erro ao decodificar a categoria", http.StatusBadRequest)
		return
	}
	repository := repositories.NewCategoriaRepository()
	if err := repository.Update(categoria); err != nil {
		utils.RespondWithError(w, "Erro ao atualizar a categoria", http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(categoria); err != nil {
		utils.RespondWithError(w, "Erro ao codificar a categoria atualizada", http.StatusInternalServerError)
		return
	}
}

// Deletar uma categoria (via query string: ?id=1)
func DeleteCategoria(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		utils.RespondWithError(w, "ID não fornecido", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(w, "ID inválido", http.StatusBadRequest)
		return
	}
	repository := repositories.NewCategoriaRepository()
	if err := repository.Delete(id); err != nil {
		utils.RespondWithError(w, "Erro ao deletar a categoria", http.StatusInternalServerError)
		return
	}
	if _, err := w.Write([]byte("Categoria deletada com sucesso")); err != nil {
		utils.RespondWithError(w, "Erro ao escrever a resposta: "+err.Error(), http.StatusInternalServerError)
	}
}
