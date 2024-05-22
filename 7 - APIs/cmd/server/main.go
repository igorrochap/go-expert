package main

import "github.com/igorrochap/goexpert/7-APIs/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
