package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var numero2 int
var leyenda string
var err error

func IngresosNumeros() {
	//lectura por teclado
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese numero 1: ")

	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("ERORR!!!! Dato incorrecto" + err.Error())
		}
	}

	fmt.Println("Ingrese numero 2: ")

	if scanner.Scan() {
		numero2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("ERORR!!!! Dato incorrecto" + err.Error())
		}
	}

	fmt.Println("Ingrese leyenda: ")

	if scanner.Scan() {
		leyenda = scanner.Text()
	}
	fmt.Println(leyenda, numero1*numero2)
}

func Multiplicar() int {
	//lectura por teclado
	scanner := bufio.NewScanner(os.Stdin)

	var numero int
	var err error

	for {
		fmt.Println("Ingrese numero 1: ")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}
	return numero
}
