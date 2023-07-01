package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Address struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	http.HandleFunc("/", searchByCepHandler)
	http.ListenAndServe(":8080", nil)
}

func searchByCepHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	cep := request.URL.Query().Get("cep")
	if cep == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	address, error := searchCEP(cep)
	if error != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	json.NewEncoder(writer).Encode(address)
}

func searchCEP(cep string) (*Address, error) {
	response, error := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if error != nil {
		return nil, error
	}
	defer response.Body.Close()
	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		return nil, error
	}
	var address Address
	error = json.Unmarshal(body, &address)
	if error != nil {
		return nil, error
	}
	return &address, nil
}
