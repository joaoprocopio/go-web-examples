package main

import (
	"fmt"
)

// em go, variáveis são explicitamente declaradas e
// usadas pelo compilador para verificar a correção de tipos de chamadas de funções

func main() {
	// var, declara uma ou mais variáveis
	// vc pode declara múltiplas variáveis de uma vez
	// go vai inferir o tipo das variáveis inicializadas (importante)
	// o fato de inferir o valor inicial faz com que a variável abaixo seja uma string válida
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	// variaveis declaradas sem inicialização correspondente, tem seu o valor zero
	// por exemplo, o valor zero de um int é 0
	var e int
	fmt.Println(e)

	// e também como exemplo, o valor zero de uma string é "", uma string vazia
	var f string
	fmt.Println(f, len(f))

	// a sintaxe :=, é uma forma abreviada de declarar e inicializar uma variável
	// por exemplo, esta abreviação é equivalente a: var f string = "apple"
	// esta sintaxe está disponível somente dentro de funções (importante)
	g := "apple"
	fmt.Println(g)
}
