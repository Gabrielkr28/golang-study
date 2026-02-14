package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/seu-usuario/ibge-api/internal/client"
	"github.com/seu-usuario/ibge-api/internal/config"
	"github.com/seu-usuario/ibge-api/internal/handler"
	"github.com/seu-usuario/ibge-api/internal/service"
	"github.com/seu-usuario/ibge-api/pkg/response"
)

func main() {
	// 1. Carrega as configurações
	cfg := config.Load()
	log.Printf("Iniciando servidor na porta %s", cfg.ServerPort)
	
	// 2. Inicializa as dependências (Dependency Injection)
	// A ordem importa: client -> service -> handler
	ibgeClient := client.NewIBGEClient(cfg.IBGEAPIUrl)
	estadoService := service.NewEstadoService(ibgeClient)
	estadoHandler := handler.NewEstadoHandler(estadoService)
	
	// 3. Configura o roteador (usando gorilla/mux para roteamento avançado)
	router := mux.NewRouter()
	
	// 4. Define as rotas da API
	// Versionamento de API (v1) é uma boa prática
	api := router.PathPrefix("/api/v1").Subrouter()
	
	// Rotas de estados
	api.HandleFunc("/estados", estadoHandler.GetEstados).Methods("GET")
	api.HandleFunc("/estados/{uf}", estadoHandler.GetEstadoByUF).Methods("GET")
	api.HandleFunc("/estados/{uf}/municipios", estadoHandler.GetMunicipios).Methods("GET")
	
	// 5. Rota de health check (importante para monitoramento)
	router.HandleFunc("/health", healthCheckHandler).Methods("GET")
	
	// 6. Middleware de logging (registra todas as requisições)
	router.Use(loggingMiddleware)
	
	// 7. Inicia o servidor HTTP
	log.Printf("Servidor rodando em http://localhost:%s", cfg.ServerPort)
	log.Println("Endpoints disponíveis:")
	log.Println("  GET /health")
	log.Println("  GET /api/v1/estados")
	log.Println("  GET /api/v1/estados/{uf}")
	log.Println("  GET /api/v1/estados/{uf}/municipios")
	
	if err := http.ListenAndServe(":"+cfg.ServerPort, router); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}

// healthCheckHandler verifica se a API está funcionando
// Usado por ferramentas de monitoramento (Kubernetes, Docker, etc.)
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, map[string]string{
		"status": "healthy",
		"service": "ibge-api",
	})
}

// loggingMiddleware é um middleware que registra todas as requisições
// Middlewares são funções que interceptam requisições antes dos handlers
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.Method, r.RequestURI, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
