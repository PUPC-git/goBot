package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)
	fmt.Println(paises)

	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"

	fmt.Println(paises)
	fmt.Println(paises["Argentina"])

	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 39,
	}
	fmt.Println(campeonato)

	for equipo, puntuacion := range campeonato {
		fmt.Printf("Equipo %s, tiene una puntuacion de %d \n", equipo, puntuacion)
	}

	delete(campeonato, "Chivas")
	fmt.Println(campeonato)

	puntuacion, existe := campeonato["Real Madrid"]
	fmt.Printf("La puntuacion es %d, y el equipo existe = %t \n", puntuacion, existe)
}
