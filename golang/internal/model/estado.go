package model

// Estado representa um estado brasileiro retornado pela API do IBGE
// As tags `json` definem como os campos serão serializados/deserializados
type Estado struct {
	ID    int    `json:"id"`    // ID numérico do estado no IBGE
	Sigla string `json:"sigla"` // UF do estado (ex: SP, RJ)
	Nome  string `json:"nome"`  // Nome completo do estado
}

// Municipio representa um município brasileiro
type Municipio struct {
	ID     int    `json:"id"`
	Nome   string `json:"nome"`
	Estado Estado `json:"microrregiao"` // Relacionamento com o estado
}
