package response

import (
	"encoding/json"
	"net/http"
)

// Response é a estrutura padrão para todas as respostas da API
// Isso garante consistência nas respostas para os clientes
type Response struct {
	Success bool        `json:"success"`          // Indica se a operação foi bem-sucedida
	Data    interface{} `json:"data,omitempty"`   // Dados retornados (omitido se vazio)
	Error   string      `json:"error,omitempty"`  // Mensagem de erro (omitido se vazio)
}

// JSON envia uma resposta JSON padronizada
// Esta função centraliza a lógica de resposta HTTP
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	
	response := Response{
		Success: statusCode >= 200 && statusCode < 300,
		Data:    data,
	}
	
	json.NewEncoder(w).Encode(response)
}

// Error envia uma resposta de erro padronizada
func Error(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	
	response := Response{
		Success: false,
		Error:   message,
	}
	
	json.NewEncoder(w).Encode(response)
}
