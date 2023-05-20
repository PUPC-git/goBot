package main

import (
	"fmt"

	"github.com/PUPC-git/goBot/variables"
)

func main() {
	variables.MuestroEnteros()
	variables.RestoVariables()
	estado, texto := variables.ConviertoTexto(32)
	fmt.Println("Pinto VariableGlobal ", variables.VariableGlobal)
	fmt.Println("Estado: ", estado)
	fmt.Println("Texto: ", texto)
}
