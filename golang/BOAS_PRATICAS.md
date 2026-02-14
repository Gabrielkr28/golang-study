# 🌟 Boas Práticas em Go - Guia Completo

## 📝 Nomenclatura

### ✅ Bom
```go
// Variáveis e funções: camelCase
var userName string
func getUserByID(id int) {}

// Tipos: PascalCase
type UserService struct {}

// Constantes: PascalCase ou UPPER_CASE
const MaxRetries = 3
const API_TIMEOUT = 10

// Acrônimos: todos maiúsculos ou minúsculos
var userID int        // ✅
var httpClient *http.Client  // ✅
type APIResponse struct {}   // ✅
```

### ❌ Ruim
```go
var user_name string      // ❌ não use snake_case
var userId int            // ❌ acrônimo parcial
func GetUserById() {}     // ❌ acrônimo parcial
type Api_Response struct {} // ❌ snake_case
```

## 🎯 Estrutura de Código

### ✅ Ordem de Declarações

```go
package handler

// 1. Imports
import (
    "net/http"
    "github.com/seu-usuario/projeto/internal/service"
)

// 2. Constantes
const (
    MaxPageSize = 100
    DefaultPage = 1
)

// 3. Variáveis globais (evite quando possível)
var (
    ErrNotFound = errors.New("não encontrado")
)

// 4. Tipos
type UserHandler struct {
    service *service.UserService
}

// 5. Construtores
func NewUserHandler(service *service.UserService) *UserHandler {
    return &UserHandler{service: service}
}

// 6. Métodos públicos
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
    // implementação
}

// 7. Métodos privados
func (h *UserHandler) validateInput(input string) error {
    // implementação
}

// 8. Funções auxiliares
func parseID(s string) (int, error) {
    // implementação
}
```

## 🔧 Error Handling

### ✅ Bom - Sempre verifique erros

```go
// Verificação imediata
data, err := fetchData()
if err != nil {
    return fmt.Errorf("erro ao buscar dados: %w", err)
}

// Wrapping de erros (Go 1.13+)
if err := processData(); err != nil {
    return fmt.Errorf("falha no processamento: %w", err)
}

// Erros customizados
type ValidationError struct {
    Field string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Message)
}
```

### ❌ Ruim

```go
// Ignorar erros
data, _ := fetchData()  // ❌ NUNCA faça isso!

// Panic em código de biblioteca
if err != nil {
    panic(err)  // ❌ Use panic apenas em main ou casos extremos
}

// Erro genérico
return errors.New("erro")  // ❌ Seja específico!
```

## 🎨 Interfaces

### ✅ Bom - Interfaces pequenas e focadas

```go
// Interface pequena (melhor)
type Reader interface {
    Read(p []byte) (n int, err error)
}

// Interface no consumidor, não no produtor
// Defina a interface onde você USA, não onde implementa

// client/client.go (implementação)
type httpClient struct {
    baseURL string
}

func (c *httpClient) Get(url string) (*Response, error) {
    // implementação
}

// service/service.go (consumidor)
type HTTPClient interface {  // ← Interface definida aqui!
    Get(url string) (*Response, error)
}

type Service struct {
    client HTTPClient  // Usa a interface
}
```

### ❌ Ruim

```go
// Interface muito grande
type DataStore interface {
    Create()
    Read()
    Update()
    Delete()
    List()
    Search()
    Count()
    // ... muitos métodos
}

// Prefixo "I" desnecessário
type IUserService interface {}  // ❌ Não use "I"
```

## 🏗️ Structs

### ✅ Bom

```go
// Campos exportados (públicos) começam com maiúscula
type User struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    CreatedAt time.Time `json:"created_at"`
}

// Constructor com validação
func NewUser(name, email string) (*User, error) {
    if name == "" {
        return nil, errors.New("nome não pode ser vazio")
    }
    
    return &User{
        Name:      name,
        Email:     email,
        CreatedAt: time.Now(),
    }, nil
}

// Métodos com receiver de ponteiro quando modifica
func (u *User) UpdateEmail(email string) {
    u.Email = email
}

// Métodos com receiver de valor quando não modifica
func (u User) IsValid() bool {
    return u.Name != "" && u.Email != ""
}
```

### ❌ Ruim

```go
// Campos privados sem necessidade
type User struct {
    id    int     // ❌ Não pode ser serializado
    name  string  // ❌ Não pode ser acessado
}

// Sem validação
func NewUser(name string) *User {
    return &User{Name: name}  // ❌ E se name for vazio?
}
```

## 🔄 Concorrência

### ✅ Bom

```go
// Usar goroutines com WaitGroup
func ProcessItems(items []Item) {
    var wg sync.WaitGroup
    
    for _, item := range items {
        wg.Add(1)
        go func(i Item) {
            defer wg.Done()
            process(i)
        }(item)  // ← Passa item como parâmetro!
    }
    
    wg.Wait()
}

// Usar channels para comunicação
func FetchData() <-chan Result {
    ch := make(chan Result)
    
    go func() {
        defer close(ch)
        // busca dados
        ch <- result
    }()
    
    return ch
}

// Context para cancelamento
func DoWork(ctx context.Context) error {
    select {
    case <-ctx.Done():
        return ctx.Err()
    case result := <-work():
        return process(result)
    }
}
```

### ❌ Ruim

```go
// Closure sem parâmetro (race condition!)
for _, item := range items {
    go func() {
        process(item)  // ❌ item muda no loop!
    }()
}

// Goroutine sem controle
go doSomething()  // ❌ Como saber quando terminou?

// Channel sem close
ch := make(chan int)
go func() {
    ch <- 42
    // ❌ Esqueceu de fechar!
}()
```

## 📦 Packages

### ✅ Bom

```go
// Nome do package = nome do diretório
// internal/handler/user_handler.go
package handler  // ✅

// Imports organizados
import (
    // 1. Standard library
    "fmt"
    "net/http"
    
    // 2. Bibliotecas externas
    "github.com/gorilla/mux"
    
    // 3. Packages internos
    "github.com/seu-usuario/projeto/internal/service"
)

// Evite import circular
// ✅ handler → service → client
// ❌ handler → service → handler (circular!)
```

### ❌ Ruim

```go
// Nome diferente do diretório
// internal/handler/user_handler.go
package user  // ❌ Deveria ser "handler"

// Imports desorganizados
import (
    "github.com/gorilla/mux"
    "fmt"
    "github.com/seu-usuario/projeto/internal/service"
    "net/http"
)
```

## 🧪 Testes

### ✅ Bom

```go
// Nome do arquivo: xxx_test.go
// package_test para testes de caixa preta
package service_test

import (
    "testing"
    "github.com/seu-usuario/projeto/internal/service"
)

// Nome do teste: TestFuncao ou TestFuncao_Cenario
func TestGetUser_Success(t *testing.T) {
    // Arrange (preparar)
    mockClient := &MockClient{
        user: &User{ID: 1, Name: "João"},
    }
    svc := service.NewUserService(mockClient)
    
    // Act (executar)
    user, err := svc.GetUser(1)
    
    // Assert (verificar)
    if err != nil {
        t.Fatalf("esperava nil, recebeu: %v", err)
    }
    
    if user.Name != "João" {
        t.Errorf("esperava João, recebeu: %s", user.Name)
    }
}

// Table-driven tests
func TestValidateEmail(t *testing.T) {
    tests := []struct {
        name    string
        email   string
        wantErr bool
    }{
        {"válido", "user@example.com", false},
        {"inválido", "invalid", true},
        {"vazio", "", true},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := ValidateEmail(tt.email)
            if (err != nil) != tt.wantErr {
                t.Errorf("ValidateEmail() erro = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}
```

### ❌ Ruim

```go
// Teste sem nome descritivo
func TestFunc(t *testing.T) {}  // ❌ Teste o quê?

// Sem verificação de erro
func TestGetUser(t *testing.T) {
    user, _ := service.GetUser(1)  // ❌ Ignora erro
    // ...
}

// Teste dependente de ordem
func TestA(t *testing.T) {
    globalVar = 10
}

func TestB(t *testing.T) {
    // ❌ Depende de TestA ter executado antes
    if globalVar != 10 {
        t.Error("falhou")
    }
}
```

## 💡 Dicas Gerais

### ✅ Bom

```go
// 1. Use defer para cleanup
func ProcessFile(filename string) error {
    f, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer f.Close()  // ✅ Sempre fecha
    
    // processa arquivo
    return nil
}

// 2. Retorne cedo (early return)
func Validate(user *User) error {
    if user == nil {
        return errors.New("user é nil")
    }
    
    if user.Name == "" {
        return errors.New("nome vazio")
    }
    
    // lógica principal aqui
    return nil
}

// 3. Use zero values
var (
    count int        // 0
    name  string     // ""
    valid bool       // false
    ptr   *User      // nil
)

// 4. Prefira composição sobre herança
type Logger interface {
    Log(msg string)
}

type Service struct {
    logger Logger  // ✅ Composição
}

// 5. Documente funções exportadas
// GetUser busca um usuário pelo ID.
// Retorna erro se o usuário não for encontrado.
func GetUser(id int) (*User, error) {
    // implementação
}
```

### ❌ Ruim

```go
// 1. Não usar defer
func ProcessFile(filename string) error {
    f, _ := os.Open(filename)
    // processa
    f.Close()  // ❌ E se houver return antes?
    return nil
}

// 2. Indentação excessiva
func Validate(user *User) error {
    if user != nil {
        if user.Name != "" {
            if user.Email != "" {
                // ❌ Muita indentação!
                return nil
            }
        }
    }
    return errors.New("inválido")
}

// 3. Inicialização desnecessária
var count int = 0  // ❌ Redundante, já é 0

// 4. Função sem documentação
func GetUser(id int) (*User, error) {  // ❌ Sem doc
    // implementação
}
```

## 🎯 Padrões Comuns

### Constructor Pattern
```go
func NewService(dep Dependency) *Service {
    return &Service{dep: dep}
}
```

### Options Pattern
```go
type Option func(*Config)

func WithTimeout(d time.Duration) Option {
    return func(c *Config) {
        c.Timeout = d
    }
}

func NewClient(opts ...Option) *Client {
    cfg := &Config{Timeout: 10 * time.Second}
    for _, opt := range opts {
        opt(cfg)
    }
    return &Client{config: cfg}
}
```

### Functional Options
```go
client := NewClient(
    WithTimeout(5 * time.Second),
    WithRetries(3),
)
```

## 📚 Recursos

- [Effective Go](https://go.dev/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)
- [Go Proverbs](https://go-proverbs.github.io/)

## 🎓 Princípios

1. **Simplicidade**: Código simples é melhor que código "inteligente"
2. **Clareza**: Código claro é melhor que código conciso
3. **Composição**: Prefira composição sobre herança
4. **Interfaces**: Pequenas e focadas
5. **Erros**: Sempre trate erros explicitamente
6. **Concorrência**: Use quando necessário, não por padrão
7. **Testes**: Teste comportamento, não implementação

---

**Lembre-se:** Go valoriza simplicidade e clareza acima de tudo! 🚀
