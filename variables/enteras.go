package variables

import "fmt"

func MuestroEnteros() {
	muestro()
}

func muestro() {
	var intComun int
	intde32 := int32(10)
	intde64 := int64(100)

	fmt.Println("IntComun = ", intComun)
	fmt.Println("Intde32 = ", intde32)
	fmt.Println("Intde64 = ", intde64)
}
