package main

import "fmt"

func exibeMap(M map[int]string) {
	for i, n := range M{
		fmt.Printf("Indice: %d, Nome: %s\n",i ,n)
	}
}
