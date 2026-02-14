package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/seu-usuario/ibge-api/internal/service"
	"github.com/seu-usuario/ibge-api/pkg/response"
)

// EstadoHandler gerencia as requisições HTTP relacionadas a estados
// Esta é a camada de apresentação (controllers em outros frameworks)
type EstadoHandler struct {
	service *service.EstadoService
}

// NewEstadoHandler cria um novo handler
// Recebe o service como dependência
func NewEstadoHandler(service *service.EstadoService) *EstadoHandler {
	return &EstadoHandler{
		service: service,
	}
}

// GetEstados manipula GET /api/v1/estados
// Lista todos os estados brasileiros
func (h *EstadoHandler) GetEstados(w http.ResponseWriter, r *http.Request) {
	log.Println("GET /api/v1/estados - Listando todos os estados")
	
	// Chama o service para buscar os dados
	estados, err := h.service.GetAllEstados()
	if err != nil {
		log.Printf("Erro ao buscar estados: %v", err)
		response.Error(w, http.StatusInternalServerError, "Erro ao buscar estados")
		return
	}
	
	// Retorna resposta de sucesso
	response.JSON(w, http.StatusOK, estados)
}

// GetEstadoByUF manipula GET /api/v1/estados/{uf}
// Busca um estado específico pela sigla
func (h *EstadoHandler) GetEstadoByUF(w http.ResponseWriter, r *http.Request) {
	// Extrai o parâmetro da URL usando gorilla/mux
	vars := mux.Vars(r)
	uf := vars["uf"]
	
	log.Printf("GET /api/v1/estados/%s - Buscando estado", uf)
	
	estado, err := h.service.GetEstadoByUF(uf)
	if err != nil {
		log.Printf("Erro ao buscar estado %s: %v", uf, err)
		response.Error(w, http.StatusNotFound, err.Error())
		return
	}
	
	response.JSON(w, http.StatusOK, estado)
}

// GetMunicipios manipula GET /api/v1/estados/{uf}/municipios
// Lista todos os municípios de um estado
func (h *EstadoHandler) GetMunicipios(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uf := vars["uf"]
	
	log.Printf("GET /api/v1/estados/%s/municipios - Listando municípios", uf)
	
	municipios, err := h.service.GetMunicipiosByUF(uf)
	if err != nil {
		log.Printf("Erro ao buscar municípios de %s: %v", uf, err)
		response.Error(w, http.StatusNotFound, err.Error())
		return
	}
	
	response.JSON(w, http.StatusOK, municipios)
}
