package main

import (
	"fmt"
)

//função de introdução de menu e de comando

func main() {

	exibirIntroducao()
	exibirMenu()
	comando := LerComando()
	//automatizar o if e else

	switch comando {
	case 1:

		iniciarMonitoramento()
	case 2:
		fmt.Println("Exibir")
	case 3:
		fmt.Println("Exit")
	default: //else
		fmt.Printf("No is valid")

	}

}
func exibirIntroducao() {
	//colocando  varivel (string)
	var nome string = "Nery"
	fmt.Println("Ola sr.", nome)

	//colocando a varievel int(nummero inteiro)
	var idade int = 20
	fmt.Println("sua idade é :", idade)

}
func exibirMenu() {
	//função scanf é como um ponteiro para a variavel que guardara a entrada do usuario

	fmt.Println("1- iniciar monitoramento ")
	fmt.Println("2- exibir ")
	fmt.Println("0- exit ")

}
func LerComando() int {
	var comando int
	fmt.Scanf("%d", &comando)

	fmt.Println("o comando escolido foi..", comando)

	return comando

}
func iniciarMonitoramento() { // multiplos retornos
	fmt.Println("Iniciar monitoramento")
	site := "https://www.alura.com.br"
	//resp, err := http.Get(site)

}
