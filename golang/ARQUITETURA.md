# Arquitetura do Projeto - Explicação Detalhada

## 🎯 Visão Geral

Este projeto segue a **Clean Architecture** (Arquitetura Limpa), um padrão usado por empresas como:
- Mercado Livre
- Uber
- Google
- Netflix
- Amazon

## 📐 Princípios Fundamentais

### 1. Separação de Responsabilidades (SoC)
Cada camada tem uma responsabilidade específica e não deve fazer o trabalho de outra.

### 2. Dependency Injection (DI)
As dependências são injetadas, não criadas internamente. Isso facilita testes e manutenção.

### 3. Inversão de Dependência
Camadas superiores não dependem de implementações concretas, mas de interfaces.

## 🏗️ Estrutura de Camadas

```
┌─────────────────────────────────────┐
│         HTTP Request                │
└──────────────┬──────────────────────┘
               │
               ▼
┌─────────────────────────────────────┐
│  HANDLER (Presentation Layer)       │  ← Recebe requisições HTTP
│  - Valida entrada                   │
│  - Chama service                    │
│  - Retorna resposta HTTP            │
└──────────────┬──────────────────────┘
               │
               ▼
┌─────────────────────────────────────┐
│  SERVICE (Business Logic Layer)     │  ← Lógica de negócio
│  - Processa regras                  │
│  - Orquestra operações              │
│  - Transforma dados                 │
└──────────────┬──────────────────────┘
               │
               ▼
┌─────────────────────────────────────┐
│  CLIENT (Data Access Layer)         │  ← Acesso a dados externos
│  - Faz requisições HTTP             │
│  - Trata erros de rede              │
│  - Serializa/deserializa            │
└──────────────┬──────────────────────┘
               │
               ▼
┌─────────────────────────────────────┐
│         External API (IBGE)         │
└─────────────────────────────────────┘
```

## 📁 Estrutura de Diretórios Explicada

### `/cmd/api/`
**Propósito**: Entry point da aplicação

```go
cmd/
└── api/
    └── main.go  // Função main() que inicia tudo
```

**Por que separar?**
- Permite múltiplos entry points (API, CLI, worker)
- Mantém a lógica de inicialização isolada
- Facilita criar diferentes binários

**Exemplo de múltiplos entry points:**
```
cmd/
├── api/main.go      // Servidor HTTP
├── worker/main.go   // Background jobs
└── cli/main.go      // Ferramenta de linha de comando
```

### `/internal/`
**Propósito**: Código privado da aplicação

```go
internal/
├── config/      // Configurações
├── handler/     // HTTP handlers
├── service/     // Lógica de negócio
├── client/      // Clientes externos
└── model/       // Estruturas de dados
```

**Por que `internal`?**
- Go impede que outros projetos importem código de `internal/`
- Garante que o código é privado da aplicação
- Evita dependências indesejadas

### `/pkg/`
**Propósito**: Código reutilizável/público

```go
pkg/
└── response/    // Utilitários de resposta HTTP
```

**Diferença entre `internal/` e `pkg/`:**
- `internal/`: Privado, só este projeto usa
- `pkg/`: Público, outros projetos podem importar

### `/internal/handler/`
**Camada de Apresentação (Controllers)**

**Responsabilidades:**
- ✅ Receber requisições HTTP
- ✅ Validar entrada (query params, path params)
- ✅ Chamar o service apropriado
- ✅ Retornar resposta HTTP formatada
- ❌ NÃO deve ter lógica de negócio
- ❌ NÃO deve acessar banco/APIs diretamente

**Exemplo:**
```go
func (h *EstadoHandler) GetEstados(w http.ResponseWriter, r *http.Request) {
    // 1. Validação (se necessário)
    // 2. Chama service
    estados, err := h.service.GetAllEstados()
    
    // 3. Trata erro
    if err != nil {
        response.Error(w, http.StatusInternalServerError, "Erro ao buscar estados")
        return
    }
    
    // 4. Retorna resposta
    response.JSON(w, http.StatusOK, estados)
}
```

### `/internal/service/`
**Camada de Lógica de Negócio**

**Responsabilidades:**
- ✅ Implementar regras de negócio
- ✅ Validar dados (regras complexas)
- ✅ Orquestrar chamadas a múltiplos clients
- ✅ Transformar dados
- ✅ Aplicar políticas (cache, retry, etc.)
- ❌ NÃO deve saber sobre HTTP
- ❌ NÃO deve formatar respostas HTTP

**Exemplo:**
```go
func (s *EstadoService) GetEstadoByUF(uf string) (*model.Estado, error) {
    // 1. Validação de negócio
    if len(uf) != 2 {
        return nil, fmt.Errorf("UF inválida")
    }
    
    // 2. Normalização
    uf = strings.ToUpper(uf)
    
    // 3. Busca dados
    estado, err := s.ibgeClient.GetEstadoByUF(uf)
    if err != nil {
        return nil, err
    }
    
    // 4. Poderia aplicar transformações aqui
    // Por exemplo: adicionar dados calculados, cache, etc.
    
    return estado, nil
}
```

### `/internal/client/`
**Camada de Acesso a Dados**

**Responsabilidades:**
- ✅ Fazer requisições HTTP
- ✅ Tratar timeouts e erros de rede
- ✅ Serializar/deserializar JSON
- ✅ Implementar retry logic (se necessário)
- ❌ NÃO deve ter lógica de negócio
- ❌ NÃO deve validar regras de negócio

**Por que usar Interface?**
```go
// Interface permite trocar implementação facilmente
type IBGEClient interface {
    GetEstados() ([]model.Estado, error)
}

// Implementação real
type ibgeClient struct { ... }

// Implementação mock para testes
type mockIBGEClient struct { ... }
```

**Benefícios:**
- Facilita testes (usa mock em vez de API real)
- Permite trocar implementação sem mudar código
- Desacopla camadas

### `/internal/model/`
**Modelos de Dados**

**Responsabilidades:**
- ✅ Definir estruturas de dados
- ✅ Representar entidades do domínio
- ✅ Definir tags de serialização

**Exemplo:**
```go
type Estado struct {
    ID    int    `json:"id"`
    Sigla string `json:"sigla"`
    Nome  string `json:"nome"`
}
```

### `/internal/config/`
**Configurações**

**Responsabilidades:**
- ✅ Carregar variáveis de ambiente
- ✅ Definir valores padrão
- ✅ Validar configurações

**Por que centralizar?**
- Um único lugar para todas as configs
- Fácil de testar
- Fácil de documentar

## 🔄 Fluxo de uma Requisição

Vamos seguir uma requisição `GET /api/v1/estados/SP`:

### 1. Router (main.go)
```go
// Recebe a requisição e roteia para o handler correto
router.HandleFunc("/api/v1/estados/{uf}", estadoHandler.GetEstadoByUF)
```

### 2. Handler (handler/estado_handler.go)
```go
func (h *EstadoHandler) GetEstadoByUF(w http.ResponseWriter, r *http.Request) {
    // Extrai parâmetro da URL
    vars := mux.Vars(r)
    uf := vars["uf"]  // "SP"
    
    // Chama o service
    estado, err := h.service.GetEstadoByUF(uf)
    
    // Retorna resposta
    if err != nil {
        response.Error(w, 404, err.Error())
        return
    }
    response.JSON(w, 200, estado)
}
```

### 3. Service (service/estado_service.go)
```go
func (s *EstadoService) GetEstadoByUF(uf string) (*model.Estado, error) {
    // Valida
    if len(uf) != 2 {
        return nil, fmt.Errorf("UF inválida")
    }
    
    // Normaliza
    uf = strings.ToUpper(uf)
    
    // Busca dados via client
    return s.ibgeClient.GetEstadoByUF(uf)
}
```

### 4. Client (client/ibge_client.go)
```go
func (c *ibgeClient) GetEstadoByUF(uf string) (*model.Estado, error) {
    // Monta URL
    url := fmt.Sprintf("%s/localidades/estados/%s", c.baseURL, uf)
    
    // Faz requisição HTTP
    resp, err := c.httpClient.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    
    // Deserializa JSON
    var estado model.Estado
    json.NewDecoder(resp.Body).Decode(&estado)
    
    return &estado, nil
}
```

### 5. API Externa (IBGE)
```
GET https://servicodados.ibge.gov.br/api/v1/localidades/estados/SP
```

### 6. Resposta volta pelo mesmo caminho
```
Client → Service → Handler → HTTP Response
```

## 🎨 Padrões de Design Utilizados

### 1. Dependency Injection
```go
// main.go
ibgeClient := client.NewIBGEClient(cfg.IBGEAPIUrl)
estadoService := service.NewEstadoService(ibgeClient)  // Injeta client
estadoHandler := handler.NewEstadoHandler(estadoService)  // Injeta service
```

**Benefícios:**
- Facilita testes (injeta mocks)
- Reduz acoplamento
- Torna dependências explícitas

### 2. Repository Pattern
O `IBGEClient` é um repository que abstrai o acesso a dados:

```go
type IBGEClient interface {
    GetEstados() ([]model.Estado, error)
}
```

**Benefícios:**
- Pode trocar fonte de dados (API, banco, cache)
- Facilita testes
- Isola lógica de acesso a dados

### 3. Constructor Pattern
```go
func NewEstadoService(client IBGEClient) *EstadoService {
    return &EstadoService{
        client: client,
    }
}
```

**Benefícios:**
- Inicialização consistente
- Validação de dependências
- Fácil de encontrar (sempre New*)

## 🧪 Testabilidade

A arquitetura facilita testes em cada camada:

### Testar Handler
```go
// Usa mock do service
mockService := &MockEstadoService{}
handler := NewEstadoHandler(mockService)
```

### Testar Service
```go
// Usa mock do client
mockClient := &MockIBGEClient{}
service := NewEstadoService(mockClient)
```

### Testar Client
```go
// Usa servidor HTTP de teste
server := httptest.NewServer(...)
client := NewIBGEClient(server.URL)
```

## 📊 Comparação com Outras Arquiteturas

### MVC Tradicional
```
Controller → Model → View
```
**Problema**: Model faz muita coisa (negócio + dados)

### Nossa Arquitetura
```
Handler → Service → Client
```
**Vantagem**: Responsabilidades bem definidas

## 🚀 Escalabilidade

Esta arquitetura permite crescer facilmente:

### Adicionar Cache
```go
type CachedEstadoService struct {
    service *EstadoService
    cache   Cache
}

func (s *CachedEstadoService) GetEstados() ([]Estado, error) {
    // Tenta cache primeiro
    if cached := s.cache.Get("estados"); cached != nil {
        return cached, nil
    }
    
    // Busca do service
    estados, err := s.service.GetEstados()
    if err == nil {
        s.cache.Set("estados", estados)
    }
    
    return estados, err
}
```

### Adicionar Banco de Dados
```go
// Novo client para banco
type DatabaseClient interface {
    GetEstados() ([]Estado, error)
}

// Service usa ambos
type EstadoService struct {
    ibgeClient IBGEClient
    dbClient   DatabaseClient
}
```

### Adicionar Autenticação
```go
// Middleware no router
router.Use(authMiddleware)

func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Valida token
        if !isAuthenticated(r) {
            response.Error(w, 401, "Não autorizado")
            return
        }
        next.ServeHTTP(w, r)
    })
}
```

## 💡 Boas Práticas Implementadas

1. ✅ Cada arquivo tem uma responsabilidade clara
2. ✅ Interfaces para desacoplamento
3. ✅ Erros são tratados em cada camada
4. ✅ Configuração centralizada
5. ✅ Logging estruturado
6. ✅ Código testável
7. ✅ Nomenclatura consistente
8. ✅ Documentação inline

## 🎓 Conclusão

Esta arquitetura é profissional porque:

- **Manutenível**: Fácil encontrar e modificar código
- **Testável**: Cada camada pode ser testada isoladamente
- **Escalável**: Fácil adicionar funcionalidades
- **Legível**: Estrutura clara e previsível
- **Profissional**: Usado por empresas de tecnologia líderes

É assim que projetos reais são estruturados em Go!
