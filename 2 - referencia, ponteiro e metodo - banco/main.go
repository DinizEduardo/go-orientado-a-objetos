package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	/* contaDoEduardo := ContaCorrente{
		titular:       "Eduardo Diniz",
		numeroAgencia: 589,
		numeroConta:   123456,
		saldo:         125.50,
	} */

	/* contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200} */

	/* fmt.Println(contaDoEduardo == contaDoEduardo2) */
	/* fmt.Println(contaDaBruna == contaDaBruna2) */

	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"
	contaDaCris.saldo = 500

	var contaDaCris2 *ContaCorrente
	contaDaCris2 = new(ContaCorrente)
	contaDaCris2.titular = "Cris"
	contaDaCris2.saldo = 500

	fmt.Println(*contaDaCris == *contaDaCris2)
}
