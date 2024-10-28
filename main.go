package main

import "fmt"

/*----------------------------------------MAIN-----------------------------------------*/

func main() {
	var op string
	for {
		fmt.Println("\n\n\n****MENU****")
		fmt.Println("A - Opção A")
		fmt.Println("B - Opção B")
		fmt.Println("C - Opção C")
		fmt.Println("X - Sair")
		fmt.Print("Esolha uma Opcao: ")
		fmt.Scanln(&op)
		fmt.Println("\n\n\n")
		switch op {
			case "A", "a":
				TestaMaiorIdadePessoa()
			case "B", "b":
				AnimalOpcaoB()
			case "C", "c":
				//opcaoC()
			case "X", "x":
				return
			default:
				fmt.Println("Opção inválida, tente novamente.")
		}
	}
}

/*-------------------------------------OPCAO 'A'--------------------------------------*/
type Pessoa struct{
	Idade int
}

func (p Pessoa) EhMaiorIdade() bool {
	return p.Idade >= 18
}

func TestaMaiorIdadePessoa(){
	pes := Pessoa{}
	pes.Idade = 15
	if pes.EhMaiorIdade(){
		fmt.Println("Maior de idade")
	} else {
		fmt.Println("Menor de idade")
	}
}
/*----------------------------------------------------------------------------------*/

/*-------------------------------------OPCAO 'B'--------------------------------------*/
type Animal interface {	// Interface Animal que define métodos comuns para todos os animais
	FazUmSom()
	Dorme()
}

type Cobra struct{} // Cobra satisfaz a interface Animal, mesmo tendo mais metodos
func (c Cobra) FazUmSom() {} // Animal
func (c Cobra) Dorme() {} // Animal
func (c Cobra) TrocaPele() {}

type Pato struct{} // Pato satisfaz a interface Animal, mesmo tendo mais metodos
func (p Pato) FazUmSom() {} // Animal
func (p Pato) Dorme() {} // Animal
func (p Pato) Voa() {}
func (p Pato) Nada() {}

type Porco struct{} // Porco satisfaz a interface Animal, porem sem acresimos de metodos
func (prc Porco) FazUmSom() {}
func (prc Porco) Dorme() {}

func fazCobraPatoPorcoDormirem(c Cobra, p Pato, prc Porco) {
   c.Dorme()
   p.Dorme()
   prc.Dorme()
}

func fazAnimaisDormirem(animais []Animal){
   for _, a := range animais {
      a.Dorme()
   }
}

func AnimalOpcaoB() {
   // Ao invez de fazer uma funcao que recebe todos os tipos de animais que eu possuo, posso simplesmente
   // usar um slice de interface Animal
   c := Cobra{}
   p := Pato{}
   prc := Porco{}
   fazCobraPatoPorcoDormirem(c,p,prc)

   animais := []Animal{
      c,
      p,
      prc,
   }
   fazAnimaisDormirem(animais)
}
/*----------------------------------------------------------------------------------*/