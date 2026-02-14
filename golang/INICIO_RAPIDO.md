# 🚀 Início Rápido - 5 Minutos

## Passo 1: Baixar Dependências
```bash
go mod download
```

## Passo 2: Executar a API
```bash
go run cmd/api/main.go
```

Você verá:
```
2024/01/15 10:30:45 Iniciando servidor na porta 8080
2024/01/15 10:30:45 Servidor rodando em http://localhost:8080
2024/01/15 10:30:45 Endpoints disponíveis:
2024/01/15 10:30:45   GET /health
2024/01/15 10:30:45   GET /api/v1/estados
2024/01/15 10:30:45   GET /api/v1/estados/{uf}
2024/01/15 10:30:45   GET /api/v1/estados/{uf}/municipios
```

## Passo 3: Testar no Navegador

Abra seu navegador e acesse:

1. **Health Check**: http://localhost:8080/health
2. **Listar Estados**: http://localhost:8080/api/v1/estados
3. **Buscar SP**: http://localhost:8080/api/v1/estados/SP
4. **Municípios de SP**: http://localhost:8080/api/v1/estados/SP/municipios

## Passo 4: Testar com PowerShell

```powershell
# Listar todos os estados
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/estados" -Method Get

# Buscar São Paulo
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/estados/SP" -Method Get

# Listar municípios de SP
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/estados/SP/municipios" -Method Get
```

## 🎉 Pronto!

Sua API está funcionando! Agora você pode:

- Ler o **README.md** para entender o projeto completo
- Ler **ARQUITETURA.md** para entender a estrutura
- Ler **CONCEITOS_GO.md** para aprender Go
- Ler **EXEMPLOS_USO.md** para ver mais exemplos
- Ler **COMANDOS_UTEIS.md** para comandos úteis

## 📝 Próximos Passos

1. Modifique o código em `internal/handler/estado_handler.go`
2. Adicione novos endpoints
3. Execute os testes: `go test ./...`
4. Compile: `go build -o api.exe cmd/api/main.go`

## ❓ Problemas?

### Porta 8080 já está em uso?
```bash
set SERVER_PORT=3000
go run cmd/api/main.go
```

### Erro de dependências?
```bash
go mod tidy
go mod download
```

### Quer recompilar?
```bash
go build -o api.exe cmd/api/main.go
./api.exe
```

## 🎓 Aprendendo Go

Este projeto é um exemplo completo de uma aplicação Go profissional. Explore os arquivos e veja como tudo se conecta!

**Ordem recomendada de leitura:**
1. `cmd/api/main.go` - Entry point
2. `internal/handler/estado_handler.go` - Handlers HTTP
3. `internal/service/estado_service.go` - Lógica de negócio
4. `internal/client/ibge_client.go` - Cliente HTTP
5. `internal/model/estado.go` - Modelos de dados

Divirta-se! 🎉
