# Conceitos Fundamentais de Go

## 📦 Packages e Módulos

### O que é um Package?
Um package é uma coleção de arquivos Go no mesmo diretório. Todo arquivo Go começa com `package nome`.

```go
package main  // Package executável
package handler  // Package de biblioteca
```

### Tipos de Packages

1. **package main**: Cria um executável
   - Deve ter uma função `main()`
   - É o entry point da aplicação

2. **Outros packages**: Criam bibliotecas reutilizáveis
   - Podem ser importados por outros packages

### Módulos (go.mod)
Define o nome do projeto e suas dependências:

```go
module github.com/seu-usuario/ibge-api

go 1.21

require (
    github.com/gorilla/mux v1.8.1
)
```

## 🔒 Visibilidade (Public vs Private)

Go usa capitalização para definir visibilidade:

```go
// Exportado (público) - começa com maiúscula
type Estado struct {
    Nome string  // Campo público
}

// Não exportado (privado) - começa com minúscula
type estado struct {
    nome string  // Campo privado
}
```

**Regra**: Se começa com letra MAIÚSCULA, é público. Minúscula = privado.

## 🏗️ Structs

Structs são tipos compostos que agrupam dados:

```go
type Estado struct {
    ID    int    `json:"id"`     // Tag JSON para serialização
    Sigla string `json:"sigla"`
    Nome  string `json:"nome"`
}

// Criando uma instância
estado := Estado{
    ID:    35,
    Sigla: "SP",
    Nome:  "São Paulo",
}
```

### JSON Tags
As tags `json:"nome"` definem como o campo será serializado:

```go
type Pessoa struct {
    Nome  string `json:"name"`           // Será "name" no JSON
    Idade int    `json:"age,omitempty"`  // Omite se for zero
    Senha string `json:"-"`              // Nunca aparece no JSON
}
```

## 🎭 Interfaces

Interfaces definem comportamentos (métodos) sem implementação:

```go
// Definição da interface
type IBGEClient interface {
    GetEstados() ([]Estado, error)
    GetEstadoByUF(uf string) (*Estado, error)
}

// Implementação (implícita em Go!)
type ibgeClient struct {
    baseURL string
}

// Se implementar todos os métodos, automaticamente implementa a interface
func (c *ibgeClient) GetEstados() ([]Estado, error) {
    // implementação
}
```

**Importante**: Em Go, não precisa declarar que implementa uma interface. Se tem os métodos, implementa automaticamente!

## 🔧 Métodos

Métodos são funções associadas a um tipo:

```go
type EstadoService struct {
    client IBGEClient
}

// Método com receiver (c *EstadoService)
func (s *EstadoService) GetAllEstados() ([]Estado, error) {
    return s.client.GetEstados()
}
```

### Receiver por Valor vs Ponteiro

```go
// Por valor - cria uma cópia
func (s EstadoService) Metodo1() {}

// Por ponteiro - modifica o original
func (s *EstadoService) Metodo2() {}
```

**Regra geral**: Use ponteiro (*) quando:
- Precisa modificar o receiver
- O struct é grande (evita cópia)
- Quer consistência (se um método usa *, todos devem usar)

## ⚠️ Error Handling

Go não tem exceções. Erros são valores retornados:

```go
func GetEstado(uf string) (*Estado, error) {
    if uf == "" {
        return nil, fmt.Errorf("UF não pode ser vazia")
    }
    
    estado, err := client.Get(uf)
    if err != nil {
        return nil, fmt.Errorf("erro ao buscar: %w", err)  // %w preserva o erro original
    }
    
    return estado, nil
}

// Uso
estado, err := GetEstado("SP")
if err != nil {
    log.Printf("Erro: %v", err)
    return
}
// Usa estado...
```

### Criando Erros

```go
import "errors"

// Erro simples
err := errors.New("algo deu errado")

// Erro formatado
err := fmt.Errorf("usuário %s não encontrado", nome)

// Wrapping de erro (Go 1.13+)
err := fmt.Errorf("falha ao processar: %w", originalErr)
```

## 🎯 Ponteiros

Ponteiros armazenam endereços de memória:

```go
// Declaração
var p *int  // Ponteiro para int

// Obtendo endereço
x := 42
p = &x  // & obtém o endereço

// Dereferenciando
fmt.Println(*p)  // * acessa o valor (42)

// Com structs
estado := &Estado{Nome: "SP"}  // Cria ponteiro diretamente
fmt.Println(estado.Nome)       // Go permite acessar sem *
```

**Quando usar ponteiros?**
- Para modificar valores em funções
- Para evitar cópias de structs grandes
- Para permitir valores nil

## 🔄 Defer

`defer` adia a execução de uma função até o fim da função atual:

```go
func ProcessarArquivo() error {
    file, err := os.Open("arquivo.txt")
    if err != nil {
        return err
    }
    defer file.Close()  // Será executado no final, mesmo com erro
    
    // Processa arquivo...
    // Se houver return ou panic, file.Close() será chamado
}
```

**Uso comum**: Fechar recursos (arquivos, conexões, etc.)

## 🔀 Goroutines e Concorrência

Goroutines são threads leves do Go:

```go
// Função normal
processarDados()

// Goroutine - executa concorrentemente
go processarDados()

// Exemplo prático
func BuscarDados() {
    go func() {
        // Executa em paralelo
        dados := api.Get()
        fmt.Println(dados)
    }()
}
```

### Channels

Channels permitem comunicação entre goroutines:

```go
// Criar channel
ch := make(chan string)

// Enviar valor
ch <- "mensagem"

// Receber valor
msg := <-ch

// Exemplo completo
func main() {
    ch := make(chan string)
    
    go func() {
        ch <- "Olá do goroutine!"
    }()
    
    msg := <-ch
    fmt.Println(msg)
}
```

## 📝 Convenções de Nomenclatura

```go
// Variáveis e funções: camelCase
var nomeCompleto string
func calcularTotal() {}

// Tipos e constantes: PascalCase
type EstadoService struct {}
const MaxRetries = 3

// Acrônimos: todas maiúsculas ou minúsculas
var userID int      // ✅
var userId int      // ❌
type HTTPClient     // ✅
type HttpClient     // ❌

// Packages: minúsculas, sem underscore
package httputil    // ✅
package http_util   // ❌
```

## 🎨 Padrões Comuns

### Constructor Pattern

```go
// NewXxx é o padrão para construtores
func NewEstadoService(client IBGEClient) *EstadoService {
    return &EstadoService{
        client: client,
    }
}
```

### Options Pattern

```go
type Config struct {
    Timeout time.Duration
    Retries int
}

type Option func(*Config)

func WithTimeout(d time.Duration) Option {
    return func(c *Config) {
        c.Timeout = d
    }
}

func NewClient(opts ...Option) *Client {
    cfg := &Config{
        Timeout: 10 * time.Second,
        Retries: 3,
    }
    
    for _, opt := range opts {
        opt(cfg)
    }
    
    return &Client{config: cfg}
}

// Uso
client := NewClient(
    WithTimeout(5 * time.Second),
    WithRetries(5),
)
```

## 🧪 Testing

```go
// Arquivo: xxx_test.go
package service

import "testing"

// Função de teste começa com Test
func TestGetEstados(t *testing.T) {
    // Arrange
    service := NewEstadoService(mockClient)
    
    // Act
    result, err := service.GetEstados()
    
    // Assert
    if err != nil {
        t.Errorf("Esperava nil, recebeu: %v", err)
    }
    
    if len(result) != 2 {
        t.Errorf("Esperava 2, recebeu: %d", len(result))
    }
}

// Executar: go test ./...
```

## 📚 Recursos Úteis

- [Tour of Go](https://go.dev/tour/) - Tutorial interativo oficial
- [Effective Go](https://go.dev/doc/effective_go) - Guia de boas práticas
- [Go by Example](https://gobyexample.com/) - Exemplos práticos
- [Go Playground](https://go.dev/play/) - Testar código online
