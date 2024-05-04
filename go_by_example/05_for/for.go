package main

import "fmt"

// for é a *única* palavra chave para realizar laços de repetição em Go
// aqui estão alguns tipos básicos de laços de repetição

func main() {
	// o tipo mais básico, com uma única condição
	i := 0

	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	// o clássico laço do C, também visto em outras linguagens como JavaScript
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// outro jeito de realizar o comportamento básico de "faça isso N vezes", é usando a cláusula range
	for i := range 3 {
		fmt.Println("range", i)
	}

	// for sem condição declarada vai realizar seu ciclo repetidamente até você explicitamente utilizar
	// a cláusula break para quebrar a execução do laço de repetição, ou a cláusula return para encerrar a função
	for {
		fmt.Println("loop")
		break
	}

	// você também pode utilizar da cláusula continue para ir à próxima iteração do laço
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

// break e continue servem para controle de fluxo do laço de repetição
// return dentro de um laço de repetição faz com que a função seja encerrada com um retorno
