package main

import "fmt"

// caracteriscas do banco de dados
type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
	senha         string
}

// metedo para sacar dinheiro
func (c *ContaCorrente) Sacar(valorDoSaque float64, senha string) string {
	if senha != c.senha {
		return "Assinatura digital incorreta"
	}
	if valorDoSaque > 0 && valorDoSaque <= c.saldo {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	} else {
		return "Saldo insuficiente ou valor de saque inválido."
	}
}

// metodo para deposito
func (c *ContaCorrente) Depositar(valorDoDeposito float64) string {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito relizado com sucesso !"
	} else {
		return "Valor de deposito invalido "
	}
}

// metodo de tranferir dinheiro para outra conta
func (c *ContaCorrente) tranferir(valor float64, outraConta *ContaCorrente, senha string) string {
	if senha != c.senha {
		return "Assinatura digital incorreto"
	}
	if valor > 0 && valor <= c.saldo {
		c.saldo -= valor
		outraConta.saldo += valor
		return "Trasferencia relizada com sucesso!"
	} else {
		return "Saldo insuficiente ou valor incorreto"
	}
}

func main() {
	contaDoMatheus := ContaCorrente{
		titular:       "Matheus",
		numeroAgencia: 589,
		numeroConta:   5312,
		saldo:         125.5,
		senha:         "senha75",
	}
	//como nao presisa de cc cv só nessecita escrever os valores observe

	contaDaRosi := ContaCorrente{
		titular:       "Rosi",
		numeroAgencia: 415,
		numeroConta:   1745,
		saldo:         1324.5,
		senha:         "senha76",
	}
	fmt.Println(contaDoMatheus)
	fmt.Println(contaDaRosi)

	// saque com senha correta
	fmt.Println("\nsaque com senha correta:")
	fmt.Println(contaDoMatheus.Sacar(25, "senha75"))
	fmt.Println("saldo de Matheus ", contaDoMatheus.saldo)

	//transferencia com senha correta
	fmt.Println("\n tranferencia com senha correta:")
	fmt.Println(contaDoMatheus.tranferir(50, &contaDaRosi, "senha76"))
	fmt.Println("saldo de Matheus:", contaDoMatheus.saldo)
	fmt.Println("saldo de Rosi:", contaDaRosi.saldo)

}
