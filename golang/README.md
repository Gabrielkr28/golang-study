# API IBGE - Projeto Profissional em Go

> 📚 **[Ver Índice Completo da Documentação](INDICE.md)** - Guia de navegação por todos os documentos

## 📋 Sobre o Projeto

API REST profissional em Go que consulta dados públicos do IBGE (Instituto Brasileiro de Geografia e Estatística).
Este projeto demonstra a estrutura de uma aplicação Go de nível empresarial, seguindo padrões usados por empresas como Mercado Livre, Uber e Google.

## 🏗️ Arquitetura do Projeto

### Estrutura de Diretórios

```
.
├── cmd/api/              # Entry point da aplicação
├── internal/             # Código privado da aplicação
│   ├── config/          # Configurações e variáveis de ambiente
│   ├── handler/         # HTTP handlers (camada de apresentação)
│   ├── service/         # Lógica de negócio
│   ├── client/          # Clientes para APIs externas
│   └── model/           # Estruturas de dados
├── pkg/                 # Código reutilizável/público
│   └── response/        # Padronização de respostas HTTP
└── go.mod               # Dependências do projeto
```

### Padrões Arquiteturais Utilizados

1. **Clean Architecture**: Separação clara entre camadas
   - `handler`: Recebe requisições HTTP
   - `service`: Processa lógica de negócio
   - `client`: Comunica com APIs externas
   - `model`: Define estruturas de dados

2. **Dependency Injection**: Injeção de dependências para facilitar testes

3. **Repository Pattern**: Abstração do acesso a dados externos

## 🚀 Como Executar

### Pré-requisitos
- Go 1.21 ou superior
- Conexão com internet (para consultar API do IBGE)

### Instalação

```bash
# Baixar dependências
go mod download

# Executar a aplicação
go run cmd/api/main.go
```

A API estará disponível em `http://localhost:8080`

## 📡 Endpoints Disponíveis

### 1. Listar Estados
```http
GET /api/v1/estados
```

**Resposta:**
```json
{
  "success": true,
  "data": [
    {
      "id": 35,
      "sigla": "SP",
      "nome": "São Paulo"
    }
  ]
}
```

### 2. Buscar Estado por UF
```http
GET /api/v1/estados/{uf}
```

**Exemplo:** `GET /api/v1/estados/SP`

**Resposta:**
```json
{
  "success": true,
  "data": {
    "id": 35,
    "sigla": "SP",
    "nome": "São Paulo"
  }
}
```

### 3. Listar Municípios de um Estado
```http
GET /api/v1/estados/{uf}/municipios
```

**Exemplo:** `GET /api/v1/estados/SP/municipios`

### 4. Health Check
```http
GET /health
```

## 🔧 Tecnologias Utilizadas

- **Go 1.21**: Linguagem de programação
- **Gorilla Mux**: Roteamento HTTP robusto
- **net/http**: Cliente HTTP nativo do Go
- **encoding/json**: Serialização/deserialização JSON

## 📚 Conceitos Importantes em Go

### 1. Packages
- `internal/`: Código privado, não pode ser importado por outros projetos
- `pkg/`: Código público, pode ser reutilizado
- `cmd/`: Entry points da aplicação

### 2. Interfaces
Usadas para abstrair comportamentos e facilitar testes:
```go
type IBGEClient interface {
    GetEstados() ([]model.Estado, error)
}
```

### 3. Error Handling
Go não usa exceções, mas retorna erros explicitamente:
```go
data, err := client.GetEstados()
if err != nil {
    // Tratar erro
}
```

### 4. Structs e JSON Tags
```go
type Estado struct {
    ID    int    `json:"id"`
    Sigla string `json:"sigla"`
    Nome  string `json:"nome"`
}
```

## 🧪 Boas Práticas Implementadas

1. ✅ Separação de responsabilidades (handlers, services, clients)
2. ✅ Tratamento de erros consistente
3. ✅ Padronização de respostas HTTP
4. ✅ Configuração centralizada
5. ✅ Logging estruturado
6. ✅ Código idiomático em Go
7. ✅ Injeção de dependências
8. ✅ Versionamento de API (v1)

## 📖 Explicação das Camadas

### Handler (Controlador)
- Recebe requisições HTTP
- Valida entrada
- Chama o service
- Retorna resposta HTTP

### Service (Lógica de Negócio)
- Processa regras de negócio
- Orquestra chamadas a clientes externos
- Transforma dados

### Client (Integração Externa)
- Comunica com APIs externas
- Trata erros de rede
- Serializa/deserializa dados

### Model (Modelo de Dados)
- Define estruturas de dados
- Representa entidades do domínio

## 🎯 Por que essa estrutura?

Esta estrutura é usada em produção por grandes empresas porque:

1. **Escalabilidade**: Fácil adicionar novos endpoints e funcionalidades
2. **Testabilidade**: Cada camada pode ser testada isoladamente
3. **Manutenibilidade**: Código organizado e fácil de entender
4. **Reutilização**: Componentes podem ser reutilizados
5. **Colaboração**: Múltiplos desenvolvedores podem trabalhar simultaneamente

## 📚 Documentação Completa

Este projeto inclui documentação detalhada para ajudá-lo a entender e trabalhar com Go:

1. **[INICIO_RAPIDO.md](INICIO_RAPIDO.md)** - Comece aqui! Guia de 5 minutos
2. **[ARQUITETURA.md](ARQUITETURA.md)** - Entenda a estrutura do projeto em detalhes
3. **[CONCEITOS_GO.md](CONCEITOS_GO.md)** - Aprenda conceitos fundamentais de Go
4. **[EXEMPLOS_USO.md](EXEMPLOS_USO.md)** - Exemplos práticos de uso da API
5. **[COMANDOS_UTEIS.md](COMANDOS_UTEIS.md)** - Referência rápida de comandos Go
6. **[ESTRUTURA_PROJETO.md](ESTRUTURA_PROJETO.md)** - Mapa visual do projeto
7. **[DIAGRAMA_VISUAL.md](DIAGRAMA_VISUAL.md)** - Diagramas da arquitetura
8. **[BOAS_PRATICAS.md](BOAS_PRATICAS.md)** - Boas práticas em Go

## 📝 Próximos Passos

- [ ] Adicionar testes unitários
- [ ] Implementar cache (Redis)
- [ ] Adicionar middleware de autenticação
- [ ] Implementar rate limiting
- [ ] Adicionar métricas (Prometheus)
- [ ] Dockerizar a aplicação
- [ ] Adicionar CI/CD

## 🎓 Aprendizado

Este projeto foi criado como exemplo educacional de uma API REST profissional em Go. Cada arquivo está extensivamente comentado em português para facilitar o aprendizado.

**Ordem recomendada de estudo:**
1. Leia o INICIO_RAPIDO.md e execute a aplicação
2. Explore o código começando por cmd/api/main.go
3. Leia ARQUITETURA.md para entender o design
4. Estude CONCEITOS_GO.md para aprender Go
5. Pratique modificando o código

## 📄 Licença

MIT
