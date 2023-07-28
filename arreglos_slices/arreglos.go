package arreglos_slices

import "fmt"

//esto es un vector, si no se le define tamaño es un slice
//con el = al final se estan definiendo las 3 primeras posiciones del vector
var tabla [10] int  = [10]int{10, 0, 22}
var matriz [20][30]int

func MuestroArreglos() {
	tabla[7] = 33
	tabla[2] = 43

	tabla2 := [10]string{"Pablo", "Juan"}

	fmt.Println(tabla)
	fmt.Println(tabla2)

	for i:=0; i<len(tabla);i++ {
		fmt.Println(tabla[i])
	}

	matriz[7][24] = 15

	fmt.Println(matriz)
}