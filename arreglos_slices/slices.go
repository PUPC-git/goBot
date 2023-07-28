package arreglos_slices

import "fmt"

var tablaS []int = []int{2,5,4}
var arreglo [10]int = [10]int {6,99,33,556,11,4,71}

func MuestroSlices() {
	fmt.Println(tablaS)

	porcion := arreglo[3:] //slices creado desde vector de la posicion 3 haste el final
	porcion2 := arreglo[:5] //slice creado desde vector de la posicion 1 a la 5
	porcion3 := arreglo[2:6] //slice creado desde vector de la posicion 2 a la 6

    fmt.Println(arreglo)
	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {
	//los slices se pueden aumentar en tiempo de ejecucion
	//con un tamaño especifico
	elementos := make([]int, 5, 20)

	fmt.Printf("largo %d, Capacidad %d", len(elementos), cap(elementos))

    //sin indicarle un tamaño especifico
	nums := make([]int,0,0)
	for i:=0;i<100;i++ {
		nums = append(nums, i,)
	}
	fmt.Printf("largo %d, Capacidad %d", len(nums), cap(nums))
}