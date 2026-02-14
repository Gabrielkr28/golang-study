# Exemplos de Uso da API

## 🚀 Iniciando a Aplicação

### Opção 1: Usando Go diretamente
```bash
# Baixar dependências
go mod download

# Executar
go run cmd/api/main.go
```

### Opção 2: Usando Makefile
```bash
# Baixar dependências
make deps

# Executar
make run
```

### Opção 3: Compilar e executar
```bash
# Compilar
go build -o api.exe cmd/api/main.go

# Executar o binário
./api.exe
```

## 📡 Testando os Endpoints

### 1. Health Check

Verifica se a API está funcionando:

```bash
curl http://localhost:8080/health
```

**Resposta:**
```json
{
  "success": true,
  "data": {
    "status": "healthy",
    "service": "ibge-api"
  }
}
```

### 2. Listar Todos os Estados

```bash
curl http://localhost:8080/api/v1/estados
```

**Resposta:**
```json
{
  "success": true,
  "data": [
    {
      "id": 11,
      "sigla": "RO",
      "nome": "Rondônia"
    },
    {
      "id": 12,
      "sigla": "AC",
      "nome": "Acre"
    },
    {
      "id": 35,
      "sigla": "SP",
      "nome": "São Paulo"
    }
    // ... todos os 27 estados
  ]
}
```

### 3. Buscar Estado Específico

```bash
# São Paulo
curl http://localhost:8080/api/v1/estados/SP

# Rio de Janeiro
curl http://localhost:8080/api/v1/estados/RJ

# Minas Gerais
curl http://localhost:8080/api/v1/estados/MG
```

**Resposta (SP):**
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

**Erro (estado não existe):**
```bash
curl http://localhost:8080/api/v1/estados/XX
```

```json
{
  "success": false,
  "error": "estado não encontrado"
}
```

### 4. Listar Municípios de um Estado

```bash
# Municípios de São Paulo
curl http://localhost:8080/api/v1/estados/SP/municipios

# Municípios do Rio de Janeiro
curl http://localhost:8080/api/v1/estados/RJ/municipios
```

**Resposta (primeiros municípios de SP):**
```json
{
  "success": true,
  "data": [
    {
      "id": 3500105,
      "nome": "Adamantina"
    },
    {
      "id": 3500204,
      "nome": "Adolfo"
    },
    {
      "id": 3550308,
      "nome": "São Paulo"
    }
    // ... 645 municípios de SP
  ]
}
```

## 🧪 Testando com PowerShell (Windows)

Se você estiver no Windows sem curl, use PowerShell:

```powershell
# Health Check
Invoke-RestMethod -Uri "http://localhost:8080/health" -Method Get

# Listar estados
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/estados" -Method Get

# Buscar estado
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/estados/SP" -Method Get

# Listar municípios
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/estados/SP/municipios" -Method Get
```

## 🌐 Testando no Navegador

Você também pode testar diretamente no navegador:

1. Inicie a aplicação: `go run cmd/api/main.go`
2. Abra o navegador e acesse:
   - http://localhost:8080/health
   - http://localhost:8080/api/v1/estados
   - http://localhost:8080/api/v1/estados/SP
   - http://localhost:8080/api/v1/estados/SP/municipios

## 🔧 Configurando Variáveis de Ambiente

Você pode customizar a porta e a URL da API do IBGE:

### Windows (CMD)
```cmd
set SERVER_PORT=3000
set IBGE_API_URL=https://servicodados.ibge.gov.br/api/v1
go run cmd/api/main.go
```

### Windows (PowerShell)
```powershell
$env:SERVER_PORT="3000"
$env:IBGE_API_URL="https://servicodados.ibge.gov.br/api/v1"
go run cmd/api/main.go
```

### Linux/Mac
```bash
export SERVER_PORT=3000
export IBGE_API_URL=https://servicodados.ibge.gov.br/api/v1
go run cmd/api/main.go
```

## 📊 Exemplos de Integração

### JavaScript (Fetch API)

```javascript
// Buscar todos os estados
fetch('http://localhost:8080/api/v1/estados')
  .then(response => response.json())
  .then(data => {
    console.log('Estados:', data.data);
  })
  .catch(error => console.error('Erro:', error));

// Buscar estado específico
async function buscarEstado(uf) {
  try {
    const response = await fetch(`http://localhost:8080/api/v1/estados/${uf}`);
    const data = await response.json();
    
    if (data.success) {
      console.log('Estado:', data.data);
    } else {
      console.error('Erro:', data.error);
    }
  } catch (error) {
    console.error('Erro de rede:', error);
  }
}

buscarEstado('SP');
```

### Python (requests)

```python
import requests

# Buscar todos os estados
response = requests.get('http://localhost:8080/api/v1/estados')
data = response.json()

if data['success']:
    for estado in data['data']:
        print(f"{estado['sigla']} - {estado['nome']}")

# Buscar municípios
response = requests.get('http://localhost:8080/api/v1/estados/SP/municipios')
municipios = response.json()['data']
print(f"São Paulo tem {len(municipios)} municípios")
```

### C# (.NET)

```csharp
using System.Net.Http;
using System.Text.Json;

var client = new HttpClient();

// Buscar estados
var response = await client.GetAsync("http://localhost:8080/api/v1/estados");
var json = await response.Content.ReadAsStringAsync();
var data = JsonSerializer.Deserialize<ApiResponse>(json);

Console.WriteLine($"Encontrados {data.Data.Count} estados");

public class ApiResponse
{
    public bool Success { get; set; }
    public List<Estado> Data { get; set; }
}

public class Estado
{
    public int Id { get; set; }
    public string Sigla { get; set; }
    public string Nome { get; set; }
}
```

## 🧪 Executando Testes

```bash
# Executar todos os testes
go test ./...

# Executar testes com detalhes
go test -v ./...

# Executar testes de um package específico
go test -v ./internal/service

# Executar com cobertura
go test -cover ./...

# Gerar relatório de cobertura HTML
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## 📈 Monitoramento

Os logs da aplicação mostram todas as requisições:

```
2024/01/15 10:30:45 Iniciando servidor na porta 8080
2024/01/15 10:30:45 Servidor rodando em http://localhost:8080
2024/01/15 10:31:12 GET /api/v1/estados 127.0.0.1:54321
2024/01/15 10:31:12 GET /api/v1/estados - Listando todos os estados
2024/01/15 10:31:15 GET /api/v1/estados/SP 127.0.0.1:54322
2024/01/15 10:31:15 GET /api/v1/estados/SP - Buscando estado
```

## 🎯 Casos de Uso Reais

### 1. Criar um Dropdown de Estados

```javascript
async function carregarEstados() {
  const response = await fetch('http://localhost:8080/api/v1/estados');
  const { data } = await response.json();
  
  const select = document.getElementById('estado');
  data.forEach(estado => {
    const option = document.createElement('option');
    option.value = estado.sigla;
    option.textContent = `${estado.sigla} - ${estado.nome}`;
    select.appendChild(option);
  });
}
```

### 2. Buscar Municípios ao Selecionar Estado

```javascript
async function carregarMunicipios(uf) {
  const response = await fetch(`http://localhost:8080/api/v1/estados/${uf}/municipios`);
  const { data } = await response.json();
  
  const select = document.getElementById('municipio');
  select.innerHTML = '<option>Selecione...</option>';
  
  data.forEach(municipio => {
    const option = document.createElement('option');
    option.value = municipio.id;
    option.textContent = municipio.nome;
    select.appendChild(option);
  });
}
```

## 🐛 Troubleshooting

### Erro: "bind: address already in use"
A porta 8080 já está em uso. Mude a porta:
```bash
set SERVER_PORT=3000
go run cmd/api/main.go
```

### Erro: "no required module provides package"
Baixe as dependências:
```bash
go mod download
```

### Erro: "cannot find package"
Execute do diretório raiz do projeto onde está o go.mod
