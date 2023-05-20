package main

import (
	"fmt"
	"runtime"

	"github.com/PUPC-git/goBot/ejercicios"
	"github.com/PUPC-git/goBot/teclado"
	"github.com/PUPC-git/goBot/variables"
)

func main() {
	variables.MuestroEnteros()
	variables.RestoVariables()
	estado, texto := variables.ConviertoTexto(32)
	fmt.Println("Pinto VariableGlobal ", variables.VariableGlobal)
	fmt.Println("Estado: ", estado)
	fmt.Println("Texto: ", texto)
	//os := runtime.GOOS
	//if os == "Linux." {
	if os := runtime.GOOS; os == "linux" {
		fmt.Println("Linux")
	} else if os == "OS X" {
		fmt.Println(os)
	} else {
		fmt.Println("Esto no es windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es linux")
	case "darwin":
		fmt.Println("Esto es darwin")
	default:
		fmt.Printf("%s \n", os)
	}

	numero, texto := ejercicios.Ejercicio1ConvNum("233")

	fmt.Println("numero: ", numero)
	fmt.Println("texto: ", texto)

	teclado.IngresosNumeros()
}
