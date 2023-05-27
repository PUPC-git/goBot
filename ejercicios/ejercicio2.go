package ejercicios

import (
	"fmt"

	"github.com/PUPC-git/goBot/teclado"
)

func TablaNumerica() string {
	var numero = teclado.Multiplicar()
	var texto string

	for i := 1; i < 10; i++ {
		fmt.Printf("%d x %d = %d \n", numero, i, i*numero)
		texto += fmt.Sprintf("%d x %d = %d \n", numero, i, i*numero)
	}
	return texto
}
