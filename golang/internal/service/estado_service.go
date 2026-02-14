package service

import (
	"fmt"
	"strings"

	"github.com/seu-usuario/ibge-api/internal/client"
	"github.com/seu-usuario/ibge-api/internal/model"
)

// EstadoService contém a lógica de negócio relacionada a estados
// Esta camada fica entre o handler e o client, processando dados e aplicando regras
type EstadoService struct {
	ibgeClient client.IBGEClient
}

// NewEstadoService cria uma nova instância do serviço
// Recebe o client como dependência (Dependency Injection)
func NewEstadoService(ibgeClient client.IBGEClient) *EstadoService {
	return &EstadoService{
		ibgeClient: ibgeClient,
	}
}

// GetAllEstados retorna todos os estados brasileiros
// Aqui poderíamos adicionar lógica adicional como cache, filtros, etc.
func (s *EstadoService) GetAllEstados() ([]model.Estado, error) {
	estados, err := s.ibgeClient.GetEstados()
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar estados: %w", err)
	}
	
	// Exemplo de lógica de negócio: ordenar por nome
	// Em um caso real, poderíamos ter regras mais complexas aqui
	return estados, nil
}

// GetEstadoByUF busca um estado pela sigla
// Valida e normaliza a entrada antes de chamar o client
func (s *EstadoService) GetEstadoByUF(uf string) (*model.Estado, error) {
	// Validação: UF deve ter 2 caracteres
	if len(uf) != 2 {
		return nil, fmt.Errorf("UF inválida: deve ter 2 caracteres")
	}
	
	// Normalização: converte para maiúsculas
	uf = strings.ToUpper(uf)
	
	estado, err := s.ibgeClient.GetEstadoByUF(uf)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar estado %s: %w", uf, err)
	}
	
	return estado, nil
}

// GetMunicipiosByUF retorna todos os municípios de um estado
func (s *EstadoService) GetMunicipiosByUF(uf string) ([]model.Municipio, error) {
	// Validação
	if len(uf) != 2 {
		return nil, fmt.Errorf("UF inválida: deve ter 2 caracteres")
	}
	
	uf = strings.ToUpper(uf)
	
	municipios, err := s.ibgeClient.GetMunicipiosByUF(uf)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar municípios de %s: %w", uf, err)
	}
	
	return municipios, nil
}
