package main

import (
	"encoding/json"
	"os"
)

type Account struct {
	Number  int `json:"number"`
	Balance int `json:"balance" validate:"gt=0"`
}

func main() {
	account := Account{Number: 1, Balance: 100}
	result, error := json.Marshal(account)
	if error != nil {
		panic(error)
	}
	println(result) //return bytes
	println(string(result))

	error = json.NewEncoder(os.Stdout).Encode(account)
	if error != nil {
		panic(error)
	}

	pureJson := []byte(`{"number": 2, "balance": 200}`)
	var accountX Account
	error = json.Unmarshal(pureJson, &accountX)
	if error != nil {
		panic(error)
	}
	println(accountX.Balance)
}
