package models

//Cliente representação da tabela pessoa
type Cliente struct {
	Idpessoa  int    `json:"idpessoa"`
	Razao     string `json:"razao"`
	Fantasia  string `json:"fantasia"`
	Bloqueado string `json:"bloqueado"`
	Cpfcnpj   string `json:"cpfcnpj"`
	Rgie      string `json:"rgie"`
}
