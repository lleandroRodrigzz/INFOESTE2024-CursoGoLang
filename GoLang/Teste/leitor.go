package main

import (
	"os"
)

func LerJSON() []byte {
	ba, err := os.ReadFile("/mnt/c/Users/leand/OneDrive/Área De Trabalho/GoLang/Teste/example.json")
	if err != nil{
		return nil
	}
	return ba
}