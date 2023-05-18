package main

import (
	"estudo/banco/clientes"
	"estudo/banco/contas"
	"fmt"
)

func main() {
	clienteUm := clientes.Titular{Nome: "Titular 1", CPF: "147.258.369-22", Profissao: "Tester"}
	contaUm := contas.ContaPoupanca{Titular: clienteUm, NumeroAgencia: 1234, NumeroConta: 253614}
	contaUm.Depositar(100.0)
	PagarBoleto(&contaUm, 60.0)
	fmt.Println(contaUm)
}

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

/*
func main() {
	contaUm := contas.ContaCorrente{Titular: clientes.Titular{Nome: "Titular 1", CPF: "", Profissao: ""}, NumeroAgencia: 1234, NumeroConta: 253614, Saldo: 100.0}
	contaDois := contas.ContaCorrente{Titular: clientes.Titular{Nome: "Titular 2", CPF: "", Profissao: ""}, NumeroAgencia: 4321, NumeroConta: 852960, Saldo: 110.0}
	fmt.Println(contaUm)
	fmt.Println(contaDois)
	fmt.Println("##############################")

	status := contaUm.Transferir(20.0, &contaDois)

	fmt.Println(status)
	fmt.Println(contaUm)
	fmt.Println(contaDois)
	fmt.Println("##############################")

	if !status {
		test()
	}
}

func test() {
	contaUm := contas.ContaCorrente{Titular: clientes.Titular{Nome: "Titular 1", CPF: "", Profissao: ""}, NumeroAgencia: 1234, NumeroConta: 253614, Saldo: 10.0}

	for {

		fmt.Println("Escolha uma opção: ")
		fmt.Println("1- Verificar Saldo")
		fmt.Println("2- Saque")
		fmt.Println("3- Depósito")
		fmt.Println("0- Sair do Programa")

		var opcao int
		fmt.Scan(&opcao)

		switch opcao {
		case 1:
			fmt.Println("Saldo Atual: ", contaUm.Saldo)
		case 2:
			var valorDoSaque float64
			fmt.Println("Digite o valor do saque: ")
			fmt.Scan(&valorDoSaque)
			retorno := contaUm.Sacar(valorDoSaque)
			fmt.Println(retorno, "Saldo Atual: ", contaUm.Saldo)
		case 3:
			var valorDoDeposito float64

			fmt.Println("Digite o valor do depósito: ")
			fmt.Scan(&valorDoDeposito)
			retorno := contaUm.Depositar(valorDoDeposito)
			fmt.Println(retorno, "Saldo Atual: ", contaUm.Saldo)
		case 0:
			os.Exit(0)
		default:
			os.Exit(-1)
		}
	}
}
*/
