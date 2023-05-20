package variables

import (
	"fmt"
	"strconv"
	"time"
)

var variableLocal int
var VariableGlobal int
var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {
	Nombre = "Pedro"
	Estado = true
	Sueldo = 1577.76
	Fecha = time.Now()

	fmt.Println("Mostrando el resto")
	fmt.Println(" - Nombre: ", Nombre)
	fmt.Println(" - Estado: ", Estado)
	fmt.Println(" - Sueldo: ", Sueldo)
	fmt.Println(" - Fecha: ", Fecha)
}

func ConviertoTexto(numero int) (bool, string) {
	//var texto string
	texto := strconv.Itoa(numero)
	return true, texto
}
