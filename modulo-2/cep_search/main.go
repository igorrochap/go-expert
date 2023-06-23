package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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
	for _, cep := range os.Args[1:] {
		request, error := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
		if error != nil {
			fmt.Fprintf(os.Stderr, "Request error: %v\n", error)
		}
		defer request.Body.Close()

		response := readRequest(*request)
		address := parseJson(response)

		writeFile(address)
	}
}

func readRequest(request http.Response) []byte {
	parsedRequest, error := io.ReadAll(request.Body)
	if error != nil {
		fmt.Fprintf(os.Stderr, "Error reading the request body: %v\n", error)
	}
	return parsedRequest
}

func parseJson(content []byte) Address {
	var address Address
	error := json.Unmarshal(content, &address)
	if error != nil {
		fmt.Fprintf(os.Stderr, "Error parsing request body: %v\n", error)
	}

	return address
}

func writeFile(address Address) {
	file, error := os.Create("address.txt")
	if error != nil {
		fmt.Fprintf(os.Stderr, "Error creating the file: %v\n", error)
	}
	defer file.Close()

	_, error = file.WriteString(fmt.Sprintf("CEP: %s,\nLocalidade: %s,\nUF: %s\n", address.Cep, address.Localidade, address.Uf))
	if error != nil {
		fmt.Fprintf(os.Stderr, "Error writing on the file: %v\n", error)
	}

	fmt.Println("File created successfully!")
	fmt.Printf("City: %s\n", address.Localidade)
}
