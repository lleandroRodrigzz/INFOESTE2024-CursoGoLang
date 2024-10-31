package main

import "fmt"
import "flag"

/*-----------------------------------EXERCICIO 1--------------------------------------*/
/*func exec1(){
	mapeador := map[int]string{
		1:"Leandro",
		2:"Mateus",
		3:"Arcesti",
	}
	
	delete(mapeador,2)
	mapeador[4] = "João Vitor"
	exibeMap(mapeador)		//Funcao 'ExibeMap' está no 'funcoes.go'
}*/
/*------------------------------------------------------------------------------------*/

/*-----------------------------------EXERCICIO 2--------------------------------------*/
func main(){
	nome := flag.String("nome", "VAZIO", "Seu Nome")
	idade := flag.Int("idade", 0, "Sua Idade")
	flag.Parse()
	fmt.Printf("Nome: %s\n", *nome)
	fmt.Printf("Idade: %d\n", *idade)

	fmt.Printf("Onde vai inserir no map?")
}
/*------------------------------------------------------------------------------------*/