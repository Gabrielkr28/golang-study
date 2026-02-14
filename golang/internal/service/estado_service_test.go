package service

import (
	"errors"
	"testing"

	"github.com/seu-usuario/ibge-api/internal/model"
)

// mockIBGEClient é um mock do IBGEClient para testes
// Isso permite testar o service sem fazer requisições HTTP reais
type mockIBGEClient struct {
	estados    []model.Estado
	estado     *model.Estado
	municipios []model.Municipio
	err        error
}

func (m *mockIBGEClient) GetEstados() ([]model.Estado, error) {
	return m.estados, m.err
}

func (m *mockIBGEClient) GetEstadoByUF(uf string) (*model.Estado, error) {
	return m.estado, m.err
}

func (m *mockIBGEClient) GetMunicipiosByUF(uf string) ([]model.Municipio, error) {
	return m.municipios, m.err
}

// TestGetAllEstados testa o método GetAllEstados
func TestGetAllEstados(t *testing.T) {
	// Arrange (preparar)
	mockClient := &mockIBGEClient{
		estados: []model.Estado{
			{ID: 35, Sigla: "SP", Nome: "São Paulo"},
			{ID: 33, Sigla: "RJ", Nome: "Rio de Janeiro"},
		},
	}
	service := NewEstadoService(mockClient)
	
	// Act (executar)
	estados, err := service.GetAllEstados()
	
	// Assert (verificar)
	if err != nil {
		t.Errorf("Esperava nil, recebeu erro: %v", err)
	}
	
	if len(estados) != 2 {
		t.Errorf("Esperava 2 estados, recebeu %d", len(estados))
	}
}

// TestGetEstadoByUF_Invalid testa validação de UF inválida
func TestGetEstadoByUF_Invalid(t *testing.T) {
	mockClient := &mockIBGEClient{}
	service := NewEstadoService(mockClient)
	
	// Testa UF com tamanho inválido
	_, err := service.GetEstadoByUF("ABC")
	
	if err == nil {
		t.Error("Esperava erro para UF inválida, mas não recebeu")
	}
}

// TestGetEstadoByUF_Success testa busca bem-sucedida
func TestGetEstadoByUF_Success(t *testing.T) {
	expectedEstado := &model.Estado{
		ID:    35,
		Sigla: "SP",
		Nome:  "São Paulo",
	}
	
	mockClient := &mockIBGEClient{
		estado: expectedEstado,
	}
	service := NewEstadoService(mockClient)
	
	estado, err := service.GetEstadoByUF("sp") // Testa normalização para maiúsculas
	
	if err != nil {
		t.Errorf("Não esperava erro: %v", err)
	}
	
	if estado.Sigla != "SP" {
		t.Errorf("Esperava SP, recebeu %s", estado.Sigla)
	}
}

// TestGetEstadoByUF_Error testa tratamento de erro
func TestGetEstadoByUF_Error(t *testing.T) {
	mockClient := &mockIBGEClient{
		err: errors.New("erro de rede"),
	}
	service := NewEstadoService(mockClient)
	
	_, err := service.GetEstadoByUF("SP")
	
	if err == nil {
		t.Error("Esperava erro, mas não recebeu")
	}
}
