# Makefile para facilitar comandos comuns

.PHONY: run build test clean help

# Executa a aplicação
run:
	@echo "Iniciando servidor..."
	go run cmd/api/main.go

# Compila a aplicação
build:
	@echo "Compilando..."
	go build -o bin/api cmd/api/main.go

# Baixa as dependências
deps:
	@echo "Baixando dependências..."
	go mod download
	go mod tidy

# Executa testes
test:
	@echo "Executando testes..."
	go test -v ./...

# Limpa arquivos compilados
clean:
	@echo "Limpando..."
	rm -rf bin/

# Mostra ajuda
help:
	@echo "Comandos disponíveis:"
	@echo "  make run    - Executa a aplicação"
	@echo "  make build  - Compila a aplicação"
	@echo "  make deps   - Baixa dependências"
	@echo "  make test   - Executa testes"
	@echo "  make clean  - Remove arquivos compilados"
