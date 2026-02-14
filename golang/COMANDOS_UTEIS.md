# Comandos Úteis - Go

## 📦 Gerenciamento de Dependências

```bash
# Inicializar um novo módulo
go mod init github.com/seu-usuario/projeto

# Baixar dependências do go.mod
go mod download

# Adicionar uma dependência
go get github.com/gorilla/mux

# Adicionar versão específica
go get github.com/gorilla/mux@v1.8.1

# Remover dependências não usadas
go mod tidy

# Verificar dependências
go list -m all

# Atualizar todas as dependências
go get -u ./...

# Criar vendor (cópia local das dependências)
go mod vendor
```

## 🏃 Executar e Compilar

```bash
# Executar sem compilar
go run cmd/api/main.go

# Executar com variáveis de ambiente
SERVER_PORT=3000 go run cmd/api/main.go

# Compilar para o sistema atual
go build -o api.exe cmd/api/main.go

# Compilar para Linux (cross-compile do Windows)
set GOOS=linux
set GOARCH=amd64
go build -o api cmd/api/main.go

# Compilar para Windows
set GOOS=windows
set GOARCH=amd64
go build -o api.exe cmd/api/main.go

# Compilar para Mac
set GOOS=darwin
set GOARCH=amd64
go build -o api cmd/api/main.go

# Compilar com otimizações (binário menor)
go build -ldflags="-s -w" -o api.exe cmd/api/main.go

# Ver tamanho do binário
dir api.exe  # Windows
ls -lh api   # Linux/Mac
```

## 🧪 Testes

```bash
# Executar todos os testes
go test ./...

# Executar com detalhes
go test -v ./...

# Executar testes de um package específico
go test ./internal/service

# Executar teste específico
go test -run TestGetEstados ./internal/service

# Executar com cobertura
go test -cover ./...

# Gerar relatório de cobertura
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Executar testes em paralelo
go test -parallel 4 ./...

# Executar com timeout
go test -timeout 30s ./...

# Executar testes de benchmark
go test -bench=. ./...

# Ver quais testes seriam executados
go test -v -run=XXX ./...
```

## 🔍 Análise de Código

```bash
# Formatar código (padrão Go)
go fmt ./...

# Verificar erros comuns
go vet ./...

# Instalar golint
go install golang.org/x/lint/golint@latest

# Executar linter
golint ./...

# Instalar golangci-lint (linter completo)
# Windows: baixar de https://github.com/golangci/golangci-lint/releases
golangci-lint run

# Ver imports não usados
go list -f '{{.ImportPath}} {{.Imports}}' ./...

# Ver dependências de um package
go list -f '{{.Deps}}' ./internal/handler
```

## 📊 Informações do Projeto

```bash
# Ver versão do Go
go version

# Ver variáveis de ambiente do Go
go env

# Ver GOPATH
go env GOPATH

# Ver GOROOT
go env GOROOT

# Listar todos os packages do projeto
go list ./...

# Ver informações de um package
go list -json ./internal/handler

# Ver dependências diretas
go list -m all

# Ver dependências de uma dependência
go mod graph

# Ver por que uma dependência é necessária
go mod why github.com/gorilla/mux
```

## 🐛 Debug e Profiling

```bash
# Executar com race detector (detecta race conditions)
go run -race cmd/api/main.go

# Compilar com race detector
go build -race -o api.exe cmd/api/main.go

# Gerar profile de CPU
go test -cpuprofile=cpu.prof ./...
go tool pprof cpu.prof

# Gerar profile de memória
go test -memprofile=mem.prof ./...
go tool pprof mem.prof

# Ver goroutines em execução
# Adicione no código:
import _ "net/http/pprof"
# Acesse: http://localhost:8080/debug/pprof/
```

## 📝 Documentação

```bash
# Gerar documentação local
go doc fmt.Println

# Ver documentação de um package
go doc ./internal/handler

# Iniciar servidor de documentação
godoc -http=:6060
# Acesse: http://localhost:6060

# Ver documentação de uma função
go doc internal/handler.EstadoHandler.GetEstados
```

## 🔧 Ferramentas Úteis

```bash
# Instalar ferramentas úteis
go install golang.org/x/tools/cmd/goimports@latest  # Organiza imports
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest  # Linter
go install github.com/swaggo/swag/cmd/swag@latest  # Gera documentação Swagger

# Usar goimports (melhor que go fmt)
goimports -w .

# Gerar mocks (mockgen)
go install github.com/golang/mock/mockgen@latest
mockgen -source=internal/client/ibge_client.go -destination=internal/client/mock_client.go
```

## 🚀 Performance

```bash
# Benchmark
go test -bench=. -benchmem ./...

# Benchmark específico
go test -bench=BenchmarkGetEstados ./internal/service

# Executar benchmark N vezes
go test -bench=. -benchtime=10s ./...

# Comparar benchmarks
go test -bench=. ./... > old.txt
# Faça mudanças no código
go test -bench=. ./... > new.txt
# Instale benchcmp: go install golang.org/x/tools/cmd/benchcmp@latest
benchcmp old.txt new.txt
```

## 🐳 Docker (Bonus)

```dockerfile
# Criar Dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o api cmd/api/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/api .
EXPOSE 8080
CMD ["./api"]
```

```bash
# Construir imagem
docker build -t ibge-api .

# Executar container
docker run -p 8080:8080 ibge-api

# Executar com variáveis de ambiente
docker run -p 3000:3000 -e SERVER_PORT=3000 ibge-api
```

## 📋 Checklist de Qualidade

Antes de fazer commit, execute:

```bash
# 1. Formatar código
go fmt ./...

# 2. Organizar imports
goimports -w .

# 3. Verificar erros
go vet ./...

# 4. Executar linter
golangci-lint run

# 5. Executar testes
go test ./...

# 6. Verificar cobertura
go test -cover ./...

# 7. Limpar dependências
go mod tidy
```

## 🎯 Atalhos Úteis

```bash
# Criar alias no PowerShell (adicione ao $PROFILE)
function gorun { go run cmd/api/main.go }
function gotest { go test -v ./... }
function gofmt { go fmt ./... }

# Usar
gorun
gotest
gofmt
```

## 📚 Recursos Online

```bash
# Go Playground - testar código online
# https://go.dev/play/

# Go by Example
# https://gobyexample.com/

# Effective Go
# https://go.dev/doc/effective_go

# Go Tour
# https://go.dev/tour/

# Awesome Go - lista de bibliotecas
# https://awesome-go.com/
```

## 🔥 Dicas Rápidas

```bash
# Ver todas as funções exportadas de um package
go doc -all fmt

# Encontrar onde um símbolo é definido
go doc fmt.Println

# Limpar cache de build
go clean -cache

# Limpar cache de testes
go clean -testcache

# Ver tempo de compilação
go build -x cmd/api/main.go

# Compilar mais rápido (sem otimizações)
go build -gcflags=-N cmd/api/main.go
```
