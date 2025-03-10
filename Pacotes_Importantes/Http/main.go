package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type CEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {

	http.HandleFunc("/", BuscaCepHandle) // Criando a Rota
	http.ListenAndServe(":8080", nil)    // Subindo o servidor na porta 8080

}

func BuscaCepHandle(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" { // verificando rota
		res.WriteHeader(http.StatusNotFound) // Disparando um código de erro
		return
	}

	cepParam := req.URL.Query().Get("cep") // Setando o que será passado na Query String da URL
	if cepParam == "" {                    // Verificando se a query string está vazia
		res.WriteHeader(http.StatusNotFound) // Disparando código de erro
		return
	}

	cep, err := BuscaCep(cepParam)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(cep)

}

func BuscaCep(cep string) (*CEP, error) {
	req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/") // fazendo a requisição
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(req.Body) // Leitura do que foi retornado pela URL
	if err != nil {
		return nil, err
	}
	defer req.Body.Close() // USando o defer pra fechar a conexão

	var c CEP

	err = json.Unmarshal(body, &c) // convertendo a resposta recebida para um json e salvando na var c
	if err != nil {
		return nil, err
	}

	return &c, nil

}
