# 🔧 Troubleshooting - Soluções para Problemas Comuns

## 🚨 Problemas ao Executar

### Erro: "go: command not found"

**Problema:** Go não está instalado ou não está no PATH.

**Solução:**
```bash
# Verificar se Go está instalado
go version

# Se não estiver instalado, baixe de:
# https://go.dev/dl/

# Windows: Adicione ao PATH
# Painel de Controle → Sistema → Variáveis de Ambiente
# Adicione: C:\Go\bin
```

---

### Erro: "bind: address already in use"

**Problema:** A porta 8080 já está sendo usada por outro processo.

**Solução 1 - Mudar a porta:**
```bash
# Windows CMD
set SERVER_PORT=3000
go run cmd/api/main.go

# Windows PowerShell
$env:SERVER_PORT="3000"
go run cmd/api/main.go

# Linux/Mac
export SERVER_PORT=3000
go run cmd/api/main.go
```

**Solução 2 - Encontrar e matar o processo:**
```bash
# Windows
netstat -ano | findstr :8080
taskkill /PID [número_do_pid] /F

# Linux/Mac
lsof -i :8080
kill -9 [PID]
```

---

### Erro: "no required module provides package"

**Problema:** Dependências não foram baixadas.

**Solução:**
```bash
# Baixar dependências
go mod download

# Limpar e reorganizar
go mod tidy

# Se ainda não funcionar, deletar go.sum e tentar novamente
del go.sum  # Windows
rm go.sum   # Linux/Mac
go mod download
```

---

### Erro: "cannot find package"

**Problema:** Você está executando de um diretório errado.

**Solução:**
```bash
# Certifique-se de estar no diretório raiz do projeto
# (onde está o go.mod)

# Verificar diretório atual
cd  # Windows
pwd # Linux/Mac

# Navegar para o diretório correto
cd C:\project\golang
```

---

## 🐛 Problemas de Compilação

### Erro: "undefined: mux"

**Problema:** Import não foi feito corretamente.

**Solução:**
```go
// Adicione o import
import "github.com/gorilla/mux"

// Execute
go mod tidy
go mod download
```

---

### Erro: "imported and not used"

**Problema:** Você importou um package mas não está usando.

**Solução:**
```go
// Remova o import não usado
// OU use _ para imports de efeito colateral
import _ "github.com/lib/pq"
```

---

### Erro: "declared and not used"

**Problema:** Variável declarada mas não usada.

**Solução:**
```go
// Remova a variável
// OU use _ para ignorar
_, err := someFunc()

// OU use a variável
result, err := someFunc()
fmt.Println(result)
```

---

## 🌐 Problemas de Rede

### Erro: "dial tcp: i/o timeout"

**Problema:** Não consegue conectar à API do IBGE.

**Solução:**
```bash
# 1. Verificar conexão com internet
ping google.com

# 2. Testar URL da API diretamente
curl https://servicodados.ibge.gov.br/api/v1/localidades/estados

# 3. Verificar firewall/proxy
# Pode estar bloqueando conexões HTTP

# 4. Aumentar timeout no código
// client/ibge_client.go
httpClient: &http.Client{
    Timeout: 30 * time.Second,  // Aumentar de 10s para 30s
}
```

---

### Erro: "x509: certificate signed by unknown authority"

**Problema:** Problema com certificados SSL.

**Solução:**
```bash
# Windows: Atualizar certificados
# Baixe e instale certificados raiz

# Ou desabilite verificação SSL (APENAS PARA DESENVOLVIMENTO!)
// NÃO USE EM PRODUÇÃO!
transport := &http.Transport{
    TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
}
client := &http.Client{Transport: transport}
```

---

## 🧪 Problemas com Testes

### Erro: "no tests to run"

**Problema:** Arquivo de teste não está nomeado corretamente.

**Solução:**
```bash
# Arquivo deve terminar com _test.go
# ✅ user_service_test.go
# ❌ user_service_tests.go
# ❌ test_user_service.go

# Função deve começar com Test
# ✅ func TestGetUser(t *testing.T)
# ❌ func testGetUser(t *testing.T)
```

---

### Erro: "testing: warning: no tests to run"

**Problema:** Nenhum teste encontrado no package.

**Solução:**
```bash
# Verificar se há arquivos *_test.go
dir *_test.go  # Windows
ls *_test.go   # Linux/Mac

# Criar um teste simples
# arquivo_test.go
package mypackage

import "testing"

func TestSomething(t *testing.T) {
    // teste aqui
}
```

---

## 💾 Problemas com Go Modules

### Erro: "go.mod file not found"

**Problema:** Não há arquivo go.mod no diretório.

**Solução:**
```bash
# Inicializar módulo
go mod init github.com/seu-usuario/projeto

# Baixar dependências
go mod download
```

---

### Erro: "module declares its path as X but was required as Y"

**Problema:** Nome do módulo no go.mod não corresponde ao import.

**Solução:**
```bash
# Opção 1: Corrigir go.mod
# Edite go.mod e mude o nome do módulo

# Opção 2: Corrigir imports
# Mude os imports no código para corresponder ao go.mod
```

---

## 🔍 Problemas de Diagnóstico

### Como ver logs detalhados?

```bash
# Executar com logs verbosos
go run -v cmd/api/main.go

# Ver requisições HTTP
# Adicione no código:
log.Printf("Request: %s %s", r.Method, r.URL.Path)
```

---

### Como debugar?

```bash
# Opção 1: Usar prints
log.Printf("Debug: valor = %v", valor)

# Opção 2: Usar Delve (debugger)
go install github.com/go-delve/delve/cmd/dlv@latest
dlv debug cmd/api/main.go

# Opção 3: VS Code
# Instale extensão Go
# F5 para iniciar debug
```

---

## 🔧 Problemas de Performance

### API está lenta

**Solução:**
```bash
# 1. Verificar timeout
# Aumentar timeout no client

# 2. Adicionar cache
// Implementar cache em memória ou Redis

# 3. Profiling
go test -cpuprofile=cpu.prof
go tool pprof cpu.prof
```

---

### Muita memória sendo usada

**Solução:**
```bash
# Memory profiling
go test -memprofile=mem.prof
go tool pprof mem.prof

# Verificar goroutines vazando
# Adicione no código:
import _ "net/http/pprof"
// Acesse: http://localhost:8080/debug/pprof/
```

---

## 🐳 Problemas com Docker (se usar)

### Erro: "docker: command not found"

**Solução:**
```bash
# Instale Docker Desktop
# https://www.docker.com/products/docker-desktop
```

---

### Container não inicia

**Solução:**
```bash
# Ver logs do container
docker logs [container_id]

# Executar interativamente
docker run -it ibge-api /bin/sh

# Verificar portas
docker ps
```

---

## 📝 Problemas com IDE

### VS Code não reconhece imports

**Solução:**
```bash
# 1. Instalar extensão Go
# 2. Recarregar window (Ctrl+Shift+P → Reload Window)
# 3. Executar Go: Install/Update Tools
# 4. Verificar GOPATH
go env GOPATH
```

---

### GoLand/IntelliJ não compila

**Solução:**
```bash
# 1. File → Invalidate Caches / Restart
# 2. Verificar Go SDK configurado
# 3. Rebuild project
```

---

## 🔐 Problemas de Permissão

### Erro: "permission denied"

**Solução Windows:**
```bash
# Executar como Administrador
# Ou verificar antivírus bloqueando
```

**Solução Linux/Mac:**
```bash
# Dar permissão de execução
chmod +x api
./api
```

---

## 📊 Verificação de Saúde

### Checklist de Diagnóstico

```bash
# 1. Go instalado?
go version

# 2. Dependências baixadas?
go mod download
go mod verify

# 3. Código compila?
go build cmd/api/main.go

# 4. Testes passam?
go test ./...

# 5. Linter OK?
go vet ./...

# 6. Formatação OK?
go fmt ./...

# 7. API responde?
curl http://localhost:8080/health
```

---

## 🆘 Ainda com Problemas?

### Passos para Resolver

1. **Leia a mensagem de erro completa**
   - Não ignore detalhes
   - Copie a mensagem exata

2. **Verifique o básico**
   - Go instalado?
   - Diretório correto?
   - Dependências baixadas?

3. **Isole o problema**
   - Funciona em outro computador?
   - Funciona com código mais simples?

4. **Busque ajuda**
   - Google a mensagem de erro
   - Stack Overflow
   - Go Forum
   - GitHub Issues

5. **Recrie do zero**
   - Às vezes é mais rápido começar de novo
   - Clone o repositório novamente

---

## 📚 Recursos Úteis

- [Go FAQ](https://go.dev/doc/faq)
- [Go Wiki](https://github.com/golang/go/wiki)
- [Stack Overflow - Go](https://stackoverflow.com/questions/tagged/go)
- [Reddit - r/golang](https://reddit.com/r/golang)
- [Go Forum](https://forum.golangbridge.org/)

---

**Dica:** Sempre leia a mensagem de erro completa. Go geralmente dá mensagens muito claras sobre o que está errado! 🔍
