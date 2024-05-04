package main

import "fmt"

// trabalhar com if e else no Go é bem direto ao ponto

func main() {
	// exemplo básico
	if 7%2 == 0 {
		fmt.Println("7 é par")
	} else {
		fmt.Println("7 é ímpar")
	}

	// você pode ter um if sem um else
	if 8%4 == 0 {
		fmt.Println("8 é divisível por 4")
	}

	// operadores lógicos como && (AND) e || (OR) são de vez em quando úteis em condições
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("ou 8 ou 7 é par")
	}

	// uma declaração pode preceder condicionais, qualquer variável declarada dentro desta declaração
	// estará disponível em ramo atual e nos ramos subsequentes
	if num := 9; num < 0 {
		fmt.Println(num, "é negativo")
	} else if num < 10 {
		fmt.Println(num, "tem 1 dígito")
	} else {
		fmt.Println(num, "tem múltiplos dígitos")
	}
}

// é importante notar que você não precisa usar parenteses nas condições em GO, mas as chaves são obrigatórias
// não existe if ternário em Go, logo você tem que usar um if completo até para condições básicas
