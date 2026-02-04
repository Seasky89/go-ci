package handlers

import (
	"encoding/json"
	"myapi/internal/repositories"
	"myapi/internal/services"
	"myapi/internal/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ==================== HANDLERS PARA ITENS ====================

// Listar todos os itens
func ListItens(w http.ResponseWriter, r *http.Request) {
	repository := repositories.NewItemRepository()
	itens, err := repository.ListAll()
	if err != nil {
		utils.RespondWithError(w, "erro ao listar os itens", http.StatusNotFound)
		return
	}
	if err = json.NewEncoder(w).Encode(itens); err != nil {
		utils.RespondWithError(w, "Erro ao codificar os itens", http.StatusInternalServerError)
		return
	}
}

// Buscar um único item pelo id (via query string: ?id=1)
func GetItem(w http.ResponseWriter, r *http.Request) {
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
	repository := repositories.NewItemRepository()
	item, err := repository.FindByID(id)
	if err != nil {
		utils.RespondWithError(w, "Item não encontrado", http.StatusNotFound)
		return
	}
	if err = json.NewEncoder(w).Encode(item); err != nil {
		utils.RespondWithError(w, "Erro ao codificar o item", http.StatusInternalServerError)
		return
	}
}

// Buscar um item pelo campo "codigo"
func GetItemByCodigo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cod := vars["codigo"]
	if cod == "" {
		utils.RespondWithError(w, "Código não fornecido", http.StatusBadRequest)
		return
	}
	repository := repositories.NewItemRepository()
	item, err := repository.FindByCodigo(cod)
	if err != nil {
		utils.RespondWithError(w, "Item não encontrado", http.StatusNotFound)
		return
	}
	if err = json.NewEncoder(w).Encode(item); err != nil {
		utils.RespondWithError(w, "Erro ao codificar o item", http.StatusInternalServerError)
		return
	}
}

// Criar um novo item (envie JSON via POST)
func CreateItem(w http.ResponseWriter, r *http.Request) {
	item, err := services.DecodeAndValidateItem(r)
	if err != nil {
		utils.RespondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}

	repository := repositories.NewItemRepository()
	createdItem, err := repository.Create(item)
	if err != nil {
		utils.RespondWithError(w, "Erro ao criar o item", http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode(createdItem); err != nil {
		utils.RespondWithError(w, "Erro ao codificar o item criado", http.StatusInternalServerError)
		return
	}
}

// Atualizar um item (envie JSON via PUT, com o campo id preenchido)
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	item, err := services.DecodeAndValidateItem(r)
	if err != nil {
		utils.RespondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}

	repository := repositories.NewItemRepository()
	if err := repository.Update(item); err != nil {
		utils.RespondWithError(w, "Erro ao atualizar o item", http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(item); err != nil {
		utils.RespondWithError(w, "Erro ao codificar o item atualizado", http.StatusInternalServerError)
		return
	}
}

// Deletar um item (via query string: ?id=1)
func DeleteItem(w http.ResponseWriter, r *http.Request) {
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
	repository := repositories.NewItemRepository()
	if err := repository.Delete(id); err != nil {
		utils.RespondWithError(w, "Erro ao deletar o item", http.StatusInternalServerError)
		return
	}
	if _, err = w.Write([]byte("Item deletado com sucesso")); err != nil {
		utils.RespondWithError(w, "Erro ao escrever a resposta: "+err.Error(), http.StatusInternalServerError)
	}
}
