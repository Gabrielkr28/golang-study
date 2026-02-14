# 📂 Estrutura Completa do Projeto

```
golang/
│
├── 📁 cmd/                          # Entry points da aplicação
│   └── 📁 api/
│       └── 📄 main.go              # Função main() - inicia o servidor
│
├── 📁 internal/                     # Código privado da aplicação
│   │
│   ├── 📁 config/                  # Configurações
│   │   └── 📄 config.go           # Carrega variáveis de ambiente
│   │
│   ├── 📁 handler/                 # HTTP Handlers (Controllers)
│   │   └── 📄 estado_handler.go   # Handlers de estados
│   │
│   ├── 📁 service/                 # Lógica de Negócio
│   │   ├── 📄 estado_service.go   # Service de estados
│   │   └── 📄 estado_service_test.go  # Testes do service
│   │
│   ├── 📁 client/                  # Clientes para APIs externas
│   │   └── 📄 ibge_client.go      # Cliente da API do IBGE
│   │
│   └── 📁 model/                   # Modelos de Dados
│       └── 📄 estado.go           # Structs de Estado e Município
│
├── 📁 pkg/                          # Código público/reutilizável
│   └── 📁 response/
│       └── 📄 response.go          # Padronização de respostas HTTP
│
├── 📄 go.mod                        # Definição do módulo e dependências
├── 📄 go.sum                        # Checksums das dependências
├── 📄 Makefile                      # Comandos úteis (make run, make test)
├── 📄 .gitignore                    # Arquivos ignorados pelo Git
│
├── 📄 README.md                     # Documentação principal
├── 📄 INICIO_RAPIDO.md             # Guia de 5 minutos
├── 📄 ARQUITETURA.md               # Explicação da arquitetura
├── 📄 CONCEITOS_GO.md              # Conceitos fundamentais de Go
├── 📄 EXEMPLOS_USO.md              # Exemplos de uso da API
├── 📄 COMANDOS_UTEIS.md            # Comandos úteis do Go
└── 📄 ESTRUTURA_PROJETO.md         # Este arquivo
```

## 🎯 Responsabilidade de Cada Arquivo

### 📄 cmd/api/main.go
**O que faz:**
- Inicializa a aplicação
- Carrega configurações
- Cria instâncias (client, service, handler)
- Configura rotas HTTP
- Inicia o servidor

**Quando modificar:**
- Adicionar novos endpoints
- Adicionar middlewares
- Mudar porta do servidor
- Adicionar novos handlers

---

### 📄 internal/config/config.go
**O que faz:**
- Define struct de configuração
- Carrega variáveis de ambiente
- Define valores padrão

**Quando modificar:**
- Adicionar novas configurações
- Mudar valores padrão
- Adicionar validações de config

---

### 📄 internal/handler/estado_handler.go
**O que faz:**
- Recebe requisições HTTP
- Extrai parâmetros (query, path, body)
- Chama o service
- Retorna respostas HTTP

**Quando modificar:**
- Adicionar novos endpoints
- Mudar validações de entrada
- Mudar formato de resposta

**NÃO deve:**
- Ter lógica de negócio
- Acessar banco/APIs diretamente
- Fazer cálculos complexos

---

### 📄 internal/service/estado_service.go
**O que faz:**
- Implementa lógica de negócio
- Valida regras de negócio
- Orquestra chamadas a clients
- Transforma dados

**Quando modificar:**
- Adicionar regras de negócio
- Adicionar validações complexas
- Implementar cache
- Adicionar retry logic

**NÃO deve:**
- Saber sobre HTTP (status codes, headers)
- Formatar respostas HTTP
- Fazer requisições HTTP diretamente

---

### 📄 internal/client/ibge_client.go
**O que faz:**
- Faz requisições HTTP para API externa
- Trata erros de rede
- Serializa/deserializa JSON
- Implementa timeouts

**Quando modificar:**
- Adicionar novos endpoints da API
- Mudar timeout
- Adicionar retry logic
- Mudar URL base

**NÃO deve:**
- Ter lógica de negócio
- Validar regras de negócio
- Saber sobre handlers HTTP

---

### 📄 internal/model/estado.go
**O que faz:**
- Define estruturas de dados
- Define tags JSON
- Representa entidades do domínio

**Quando modificar:**
- Adicionar novos campos
- Criar novos modelos
- Mudar serialização JSON

---

### 📄 pkg/response/response.go
**O que faz:**
- Padroniza respostas HTTP
- Define formato JSON padrão
- Facilita retorno de erros

**Quando modificar:**
- Mudar formato de resposta
- Adicionar campos padrão
- Adicionar helpers de resposta

---

## 🔄 Fluxo de Dados

```
┌─────────────┐
│   Cliente   │
│  (Browser)  │
└──────┬──────┘
       │ HTTP Request
       ▼
┌─────────────────────────────────┐
│  cmd/api/main.go                │
│  - Router recebe requisição     │
│  - Roteia para handler correto  │
└──────┬──────────────────────────┘
       │
       ▼
┌─────────────────────────────────┐
│  internal/handler/              │
│  - Extrai parâmetros            │
│  - Valida entrada               │
│  - Chama service                │
└──────┬──────────────────────────┘
       │
       ▼
┌─────────────────────────────────┐
│  internal/service/              │
│  - Aplica regras de negócio     │
│  - Valida dados                 │
│  - Chama client                 │
└──────┬──────────────────────────┘
       │
       ▼
┌─────────────────────────────────┐
│  internal/client/               │
│  - Faz requisição HTTP          │
│  - Deserializa JSON             │
│  - Retorna dados                │
└──────┬──────────────────────────┘
       │
       ▼
┌─────────────────────────────────┐
│  API Externa (IBGE)             │
│  - Retorna dados                │
└──────┬──────────────────────────┘
       │
       │ (Resposta volta pelo mesmo caminho)
       ▼
┌─────────────────────────────────┐
│  pkg/response/                  │
│  - Formata resposta JSON        │
│  - Retorna ao cliente           │
└─────────────────────────────────┘
```

## 📊 Dependências Entre Camadas

```
main.go
  ├─→ config
  ├─→ handler
  │    └─→ service
  │         └─→ client
  │              └─→ model
  └─→ response
```

**Regra de Ouro:** Camadas superiores dependem de inferiores, nunca o contrário!

## 🎨 Padrão de Cores (Responsabilidades)

- 🟦 **Azul (Handler)**: Entrada/Saída HTTP
- 🟩 **Verde (Service)**: Lógica de Negócio
- 🟨 **Amarelo (Client)**: Acesso a Dados Externos
- 🟪 **Roxo (Model)**: Estruturas de Dados
- 🟧 **Laranja (Config)**: Configurações
- 🟥 **Vermelho (Response)**: Utilitários

## 📝 Checklist para Adicionar Funcionalidade

### Exemplo: Adicionar endpoint de regiões

1. ✅ Criar modelo em `internal/model/regiao.go`
2. ✅ Adicionar método no client `internal/client/ibge_client.go`
3. ✅ Adicionar método no service `internal/service/regiao_service.go`
4. ✅ Criar handler `internal/handler/regiao_handler.go`
5. ✅ Registrar rota em `cmd/api/main.go`
6. ✅ Criar testes `internal/service/regiao_service_test.go`

## 🧪 Onde Adicionar Testes

```
internal/
├── handler/
│   ├── estado_handler.go
│   └── estado_handler_test.go      ← Testes de handler
├── service/
│   ├── estado_service.go
│   └── estado_service_test.go      ← Testes de service
└── client/
    ├── ibge_client.go
    └── ibge_client_test.go         ← Testes de client
```

## 🚀 Arquivos de Documentação

| Arquivo | Quando Ler |
|---------|-----------|
| **INICIO_RAPIDO.md** | Primeiro! Para começar em 5 minutos |
| **README.md** | Visão geral do projeto |
| **ARQUITETURA.md** | Entender a estrutura e padrões |
| **CONCEITOS_GO.md** | Aprender conceitos de Go |
| **EXEMPLOS_USO.md** | Ver exemplos práticos |
| **COMANDOS_UTEIS.md** | Referência de comandos |
| **ESTRUTURA_PROJETO.md** | Este arquivo - mapa do projeto |

## 💡 Dicas

1. **Sempre comece pelo modelo**: Defina suas structs primeiro
2. **Client depois**: Implemente acesso aos dados
3. **Service em seguida**: Adicione lógica de negócio
4. **Handler por último**: Exponha via HTTP
5. **Testes sempre**: Teste cada camada isoladamente

## 🎓 Para Aprender

1. Leia `cmd/api/main.go` para ver como tudo se conecta
2. Siga uma requisição do handler até o client
3. Veja como os testes usam mocks
4. Tente adicionar um novo endpoint
5. Experimente modificar a lógica de negócio

---

**Lembre-se:** Esta estrutura é usada por empresas como Mercado Livre, Uber e Google. Dominar isso é dominar Go profissional! 🚀
