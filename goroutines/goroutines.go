package goroutines

import (
	"fmt"
	"strings"
	"time"
)

func MiNombreLentooo(nombre string) {
	letras := strings.Split(nombre, "")
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Printf("La letra es %s \n", letra)
	}
}