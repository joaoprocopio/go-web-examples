package main

import (
	"fmt"
	"math"
)

// go suporta constants de character, string, boolean e valores numeric
// Go supports

// const declara um valor constante
// constantes podem aparecer em qualquer lugar que a declaração var pode aparecer

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500_000_000

	// expressões constantes executam aritmética com precisão arbitrária
	// uma constante numérica não tem tipo até que seja atribuído um, como uma conversão explícita
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))

	// um número pode receber um tipo quando é usado em um contexto que requira um
	// como uma atribuição de variável ou chamada de função
	// por exemplo, math.Sin requer um float64
	fmt.Println(math.Sin(n))
}
