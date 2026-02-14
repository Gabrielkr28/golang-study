package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/seu-usuario/ibge-api/internal/model"
)

// IBGEClient é a interface que define os métodos para interagir com a API do IBGE
// Usar interface facilita testes (podemos criar mocks) e desacopla o código
type IBGEClient interface {
	GetEstados() ([]model.Estado, error)
	GetEstadoByUF(uf string) (*model.Estado, error)
	GetMunicipiosByUF(uf string) ([]model.Municipio, error)
}

// ibgeClient é a implementação concreta do IBGEClient
// O nome em minúsculo indica que é privado (não exportado)
type ibgeClient struct {
	baseURL    string
	httpClient *http.Client
}

// NewIBGEClient cria uma nova instância do cliente IBGE
// Este é o padrão "Constructor" em Go
func NewIBGEClient(baseURL string) IBGEClient {
	return &ibgeClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 10 * time.Second, // Timeout de 10 segundos para evitar requests travados
		},
	}
}

// GetEstados busca todos os estados brasileiros
func (c *ibgeClient) GetEstados() ([]model.Estado, error) {
	url := fmt.Sprintf("%s/localidades/estados", c.baseURL)
	
	// Faz a requisição HTTP GET
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer requisição: %w", err)
	}
	defer resp.Body.Close() // Importante: sempre fechar o body
	
	// Verifica se a resposta foi bem-sucedida
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API retornou status %d", resp.StatusCode)
	}
	
	// Lê o corpo da resposta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler resposta: %w", err)
	}
	
	// Deserializa o JSON para a struct
	var estados []model.Estado
	if err := json.Unmarshal(body, &estados); err != nil {
		return nil, fmt.Errorf("erro ao decodificar JSON: %w", err)
	}
	
	return estados, nil
}

// GetEstadoByUF busca um estado específico pela sigla (UF)
func (c *ibgeClient) GetEstadoByUF(uf string) (*model.Estado, error) {
	url := fmt.Sprintf("%s/localidades/estados/%s", c.baseURL, uf)
	
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer requisição: %w", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("estado não encontrado")
	}
	
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API retornou status %d", resp.StatusCode)
	}
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler resposta: %w", err)
	}
	
	var estado model.Estado
	if err := json.Unmarshal(body, &estado); err != nil {
		return nil, fmt.Errorf("erro ao decodificar JSON: %w", err)
	}
	
	return &estado, nil
}

// GetMunicipiosByUF busca todos os municípios de um estado
func (c *ibgeClient) GetMunicipiosByUF(uf string) ([]model.Municipio, error) {
	url := fmt.Sprintf("%s/localidades/estados/%s/municipios", c.baseURL, uf)
	
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer requisição: %w", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API retornou status %d", resp.StatusCode)
	}
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler resposta: %w", err)
	}
	
	var municipios []model.Municipio
	if err := json.Unmarshal(body, &municipios); err != nil {
		return nil, fmt.Errorf("erro ao decodificar JSON: %w", err)
	}
	
	return municipios, nil
}
