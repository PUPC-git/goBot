package ejercicios

import (
	"fmt"

	"github.com/PUPC-git/goBot/teclado"
)

func TablaNumerica() {
	var numero = teclado.Multiplicar()

	for i := 1; i < 10; i++ {
		fmt.Printf("%d x %d = %d \n", numero, i, i*numero)
	}
}
