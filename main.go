package main

import "fmt"

// caracteriscas do banco de dados
type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

// metedo para sacar dinheiro
func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	if valorDoSaque > 0 && valorDoSaque <= c.saldo {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	} else {
		return "Saldo insuficiente ou valor de saque inválido."
	}
}

// metodo para desito
func (c *ContaCorrente) Depositar(valorDoDeposito float64) string {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito relizado com sucesso !"
	} else {
		return "Valor de deposito invalido "
	}
}

func main() {
	contaDoMatheus := ContaCorrente{titular: "Matheus",
		numeroAgencia: 589, numeroConta: 5312, saldo: 125.5}
	//como nao presisa de cc cv só nessecita escrever os valores observe

	contaDaRosi := ContaCorrente{"Rosi", 415, 1745, 1234.5}
	fmt.Println(contaDoMatheus)
	fmt.Println(contaDaRosi)
	// sacar
	fmt.Println(contaDoMatheus.Sacar(50))
	fmt.Println(contaDoMatheus.saldo)

	// depositar
	fmt.Println(contaDaRosi.Depositar(300))
	fmt.Println(contaDaRosi.saldo)

}
