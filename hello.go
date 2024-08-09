package main

import "fmt"

func main() {
	//colocando  varivel (string)
	var nome string = "Nery"
	fmt.Println("Ola sr.", nome)

	//colocando a varievel int(nummero inteiro)
	var idade int = 20
	fmt.Println("sua idade é :", idade)

	//função scanf é como um ponteiro para a variavel que guardara a entrada do usuario

	fmt.Println("1- iniciar monitoramento ")
	fmt.Println("2- exibir ")
	fmt.Println("0- exit ")
	var comando int
	fmt.Scanf("%d", &comando)

	fmt.Println("o comando escolido foi..", comando)

	

}
