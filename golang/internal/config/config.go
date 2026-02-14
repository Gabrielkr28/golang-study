package config

import (
	"os"
)

// Config armazena todas as configurações da aplicação
// Em produção, essas configurações viriam de variáveis de ambiente
type Config struct {
	ServerPort string // Porta onde o servidor HTTP vai rodar
	IBGEAPIUrl string // URL base da API do IBGE
}

// Load carrega as configurações da aplicação
// Usa variáveis de ambiente com fallback para valores padrão
func Load() *Config {
	return &Config{
		ServerPort: getEnv("SERVER_PORT", "8080"),
		IBGEAPIUrl: getEnv("IBGE_API_URL", "https://servicodados.ibge.gov.br/api/v1"),
	}
}

// getEnv busca uma variável de ambiente ou retorna um valor padrão
// Isso é um padrão comum em aplicações Go para configuração
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
