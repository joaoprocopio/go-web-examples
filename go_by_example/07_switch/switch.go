package main

import (
	"fmt"
	"time"
)

// declarações switch expressam condicionais através de diversos ramos

func main() {
	i := 2

	fmt.Println("escreva", i, " como ")

	// aqui está um switch básico
	switch i {
	case 1:
		fmt.Println("um")
	case 2:
		fmt.Println("dois")
	case 3:
		fmt.Println("três")
	}

	// você pode usar vírgula para separar múltiplas expressões na mesma cláusula case
	// nesse exemplo também tem a cláusula default, que atribui um valor padrão para casos não cobertos
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("é fim de semana")
	default:
		fmt.Println("é dia de semana")
	}

	t := time.Now()

	// switch sem uma expressão é uma maneira de expressar lógica if/else.
	// expressões de caso de switch não precisam ser constantes
	switch {
	case t.Hour() < 12:
		fmt.Println("é antes de meio dia")
	default:
		fmt.Println("é depois de meio dia")
	}

	// um switch de tipo pode comparar tipos ao invés de valores
	// você pode utilizar isto para descobrir o tipo de uma interface
	// neste exemplo a variável t vai ter o tipo correspondente a sua clausula.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("sou bool")
		case int:
			fmt.Println("sou int")
		default:
			fmt.Printf("não conheço o tipo %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
