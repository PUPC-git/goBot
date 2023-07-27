package funciones

import (
	"fmt"
)

func Calculos() {
	var numero1 int = 32
	var numero2 int = 243

	calculo := func(numeroA int, numeroB int) int {
		return numeroA + numeroB + numero1 + numero2
	}

    fmt.Println(calculo(40,25))

    calculo = func(numeroA int, numeroB int) int {
		return numeroA + numeroB
	}

	fmt.Println(calculo(40,25))
}