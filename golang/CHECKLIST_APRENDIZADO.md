# ✅ Checklist de Aprendizado - Go Profissional

Use este checklist para acompanhar seu progresso no aprendizado de Go e arquitetura de software.

## 📚 Fase 1: Fundamentos de Go (Iniciante)

### Conceitos Básicos
- [ ] Entendo o que são packages em Go
- [ ] Sei a diferença entre `package main` e outros packages
- [ ] Entendo visibilidade (maiúscula = público, minúscula = privado)
- [ ] Sei usar `import` corretamente
- [ ] Entendo o que é `go.mod` e `go.sum`

### Tipos de Dados
- [ ] Sei criar e usar structs
- [ ] Entendo JSON tags (`json:"nome"`)
- [ ] Sei trabalhar com slices e arrays
- [ ] Entendo maps
- [ ] Sei usar ponteiros (`*` e `&`)

### Funções e Métodos
- [ ] Sei criar funções
- [ ] Entendo múltiplos retornos
- [ ] Sei criar métodos (receivers)
- [ ] Entendo a diferença entre receiver por valor e ponteiro
- [ ] Sei usar defer

### Error Handling
- [ ] Sempre verifico erros
- [ ] Sei criar erros customizados
- [ ] Entendo error wrapping (`%w`)
- [ ] Não uso panic em código de biblioteca

### Interfaces
- [ ] Entendo o que são interfaces
- [ ] Sei que implementação é implícita
- [ ] Prefiro interfaces pequenas
- [ ] Defino interfaces no consumidor, não no produtor

**Teste prático:** Crie uma struct com métodos e uma interface

---

## 🏗️ Fase 2: Arquitetura (Intermediário)

### Clean Architecture
- [ ] Entendo separação de camadas
- [ ] Sei o que faz cada camada (Handler, Service, Client)
- [ ] Entendo fluxo de dados entre camadas
- [ ] Sei por que não misturar responsabilidades

### Padrões de Design
- [ ] Entendo Dependency Injection
- [ ] Sei usar o padrão Repository
- [ ] Uso Constructor Pattern (`NewXxx`)
- [ ] Entendo Options Pattern

### Estrutura de Projeto
- [ ] Sei organizar diretórios (`cmd/`, `internal/`, `pkg/`)
- [ ] Entendo quando usar `internal/` vs `pkg/`
- [ ] Sei estruturar packages por funcionalidade
- [ ] Evito dependências circulares

**Teste prático:** Explique o fluxo de uma requisição HTTP no projeto

---

## 🧪 Fase 3: Testes (Intermediário)

### Testes Unitários
- [ ] Sei criar arquivos `*_test.go`
- [ ] Uso `testing.T` corretamente
- [ ] Escrevo testes com Arrange-Act-Assert
- [ ] Uso table-driven tests
- [ ] Sei executar testes (`go test`)

### Mocking
- [ ] Entendo por que usar mocks
- [ ] Sei criar mocks de interfaces
- [ ] Testo cada camada isoladamente
- [ ] Não testo implementação, testo comportamento

### Cobertura
- [ ] Sei gerar relatório de cobertura
- [ ] Entendo o que é cobertura de código
- [ ] Não busco 100% de cobertura cegamente

**Teste prático:** Escreva testes para um service com mock do client

---

## 🚀 Fase 4: Boas Práticas (Avançado)

### Código Limpo
- [ ] Uso nomenclatura consistente
- [ ] Escrevo código idiomático em Go
- [ ] Documento funções exportadas
- [ ] Mantenho funções pequenas e focadas
- [ ] Uso early returns

### Performance
- [ ] Entendo quando usar ponteiros
- [ ] Sei evitar alocações desnecessárias
- [ ] Uso profiling quando necessário
- [ ] Entendo goroutines e channels

### Segurança
- [ ] Valido entrada do usuário
- [ ] Trato erros adequadamente
- [ ] Não exponho informações sensíveis
- [ ] Uso timeouts em requisições HTTP

**Teste prático:** Revise código e identifique melhorias

---

## 🌐 Fase 5: APIs REST (Avançado)

### HTTP
- [ ] Entendo métodos HTTP (GET, POST, PUT, DELETE)
- [ ] Sei usar status codes corretamente
- [ ] Entendo headers HTTP
- [ ] Sei trabalhar com JSON

### Roteamento
- [ ] Sei usar gorilla/mux
- [ ] Entendo path parameters
- [ ] Sei usar query parameters
- [ ] Implemento versionamento de API

### Middleware
- [ ] Entendo o que são middlewares
- [ ] Sei criar middlewares customizados
- [ ] Uso middleware para logging
- [ ] Entendo ordem de execução

**Teste prático:** Adicione um novo endpoint à API

---

## 🔧 Fase 6: Ferramentas (Avançado)

### Go Tools
- [ ] Uso `go fmt` regularmente
- [ ] Executo `go vet` antes de commit
- [ ] Uso `go mod tidy` para limpar dependências
- [ ] Sei fazer cross-compilation

### Debug
- [ ] Sei usar logs efetivamente
- [ ] Uso debugger quando necessário
- [ ] Entendo stack traces
- [ ] Sei fazer profiling

### Git
- [ ] Faço commits atômicos
- [ ] Escrevo mensagens de commit claras
- [ ] Uso .gitignore adequadamente

**Teste prático:** Configure um projeto Go do zero

---

## 🎯 Projetos Práticos

### Projeto 1: Modificar API Existente
- [ ] Adicionar novo endpoint
- [ ] Adicionar validação de entrada
- [ ] Escrever testes para novo código
- [ ] Documentar mudanças

### Projeto 2: Criar API do Zero
- [ ] Estruturar projeto profissionalmente
- [ ] Implementar CRUD completo
- [ ] Adicionar testes
- [ ] Adicionar documentação

### Projeto 3: Adicionar Funcionalidades
- [ ] Implementar cache
- [ ] Adicionar autenticação
- [ ] Implementar rate limiting
- [ ] Adicionar métricas

---

## 📊 Avaliação de Conhecimento

### Nível Iniciante (0-30%)
- Entende sintaxe básica de Go
- Consegue ler código simples
- Precisa de ajuda para estruturar projetos

**Próximo passo:** Foque em CONCEITOS_GO.md e pratique

### Nível Intermediário (30-60%)
- Escreve código Go idiomático
- Entende arquitetura básica
- Consegue criar APIs simples

**Próximo passo:** Estude ARQUITETURA.md e BOAS_PRATICAS.md

### Nível Avançado (60-80%)
- Estrutura projetos profissionalmente
- Aplica padrões de design
- Escreve testes efetivos

**Próximo passo:** Contribua para projetos open source

### Nível Expert (80-100%)
- Domina Go e suas ferramentas
- Arquiteta sistemas complexos
- Mentora outros desenvolvedores

**Próximo passo:** Crie bibliotecas e frameworks

---

## 🎓 Certificação Informal

Você pode se considerar proficiente em Go quando conseguir:

1. ✅ Criar uma API REST do zero
2. ✅ Estruturar projeto seguindo Clean Architecture
3. ✅ Escrever testes unitários com mocks
4. ✅ Aplicar padrões de design apropriados
5. ✅ Debugar e resolver problemas
6. ✅ Revisar código de outros
7. ✅ Explicar decisões arquiteturais
8. ✅ Otimizar performance quando necessário

---

## 📝 Registro de Progresso

### Data de Início: ___/___/___

### Marcos Importantes
- [ ] Executei a API pela primeira vez
- [ ] Entendi a arquitetura do projeto
- [ ] Escrevi meu primeiro teste em Go
- [ ] Adicionei um novo endpoint
- [ ] Criei um projeto Go do zero
- [ ] Contribuí para um projeto open source

### Tempo Investido
- Leitura de documentação: ___ horas
- Prática com código: ___ horas
- Projetos pessoais: ___ horas
- Total: ___ horas

### Próximas Metas
1. ___________________________________
2. ___________________________________
3. ___________________________________

---

## 💡 Dicas de Estudo

### Para Aprender Melhor
1. **Pratique diariamente** - Mesmo que 30 minutos
2. **Leia código de outros** - Projetos open source
3. **Escreva código** - Não apenas leia
4. **Explique para outros** - Ensinar é aprender
5. **Faça projetos pessoais** - Aplique o conhecimento

### Recursos Recomendados
- [ ] Tour of Go (https://go.dev/tour/)
- [ ] Effective Go (https://go.dev/doc/effective_go)
- [ ] Go by Example (https://gobyexample.com/)
- [ ] Este projeto completo
- [ ] Projetos open source em Go

### Comunidade
- [ ] Participe do Reddit r/golang
- [ ] Siga desenvolvedores Go no Twitter
- [ ] Contribua para projetos open source
- [ ] Participe de meetups locais

---

## 🎯 Desafios Práticos

### Desafio 1: Modificação Básica
Adicione um endpoint que retorna apenas os nomes dos estados.

### Desafio 2: Validação
Adicione validação para garantir que UF tem exatamente 2 caracteres.

### Desafio 3: Cache
Implemente cache em memória para estados.

### Desafio 4: Testes
Escreva testes para todos os handlers.

### Desafio 5: Middleware
Crie um middleware de autenticação básica.

### Desafio 6: Documentação
Adicione documentação Swagger à API.

### Desafio 7: Docker
Dockerize a aplicação.

### Desafio 8: CI/CD
Configure GitHub Actions para testes automáticos.

---

## 🏆 Conquistas

Marque conforme for completando:

- [ ] 🥉 Bronze: Executei a API
- [ ] 🥈 Prata: Entendi a arquitetura
- [ ] 🥇 Ouro: Adicionei um endpoint
- [ ] 💎 Diamante: Criei um projeto do zero
- [ ] 👑 Mestre: Contribuí para open source

---

**Lembre-se:** Aprender programação é uma jornada, não um destino. Seja paciente consigo mesmo e celebre cada conquista! 🚀
