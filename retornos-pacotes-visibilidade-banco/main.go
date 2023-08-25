package main

import (
	"fmt"
	"golang-poo/retornos-pacotes-visibilidade-banco/contas"
)

func main() {
	contaDaSilvia := contas.ContaCorrente{Titular: "Silvia", Saldo: 300}
	contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100}

	status := contaDoGustavo.Transferir(-200, &contaDaSilvia)
	fmt.Println(contaDaSilvia)
	fmt.Println(status)
	fmt.Println(contaDoGustavo)

}
