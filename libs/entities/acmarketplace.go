package entities

type ACUser struct {
	ID                int     `json:"id"`
	Documento         string  `json:"documento"`
	Nome              string  `json:"nome"`
	PrimeiroNome      string  `json:"primeiroNome"`
	Email             string  `json:"email"`
	Celular           string  `json:"celular"`
	TelefoneValidado  bool    `json:"telefoneValidado"`
	UUID              string  `json:"uuid"`
	DataNascimento    string  `json:"dataNascimento"`
	Sexo              string  `json:"sexo"`
	Registrado        bool    `json:"registrado"`
	Newsletter        bool    `json:"newsletter"`
	TermsOfUse        bool    `json:"termsOfUse"`
	TermsOfUseVersion int     `json:"termsOfUseVersion"`
	CutOffValue       float64 `json:"cutOffValue"`
	CustomerIDHash    string  `json:"customerIdHash"`
}

type ACGetCustomerResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    struct {
		User ACUser `json:"user"`
		Auth struct {
			AccessToken      string `json:"access_token"`
			ExpiresInSeconds int    `json:"expires_in_seconds"`
		} `json:"auth"`
	} `json:"data"`
}
