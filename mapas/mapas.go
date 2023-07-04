package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)
	fmt.Println(paises)


	paises["Mexico"] = "D.F"
	paises["Argentina"] = "Buenos Aires"
	paises["Colombia"] = "Bogota"

	fmt.Println(paises)
	fmt.Println(paises["Colombia"])

	campeonato := map[string]int{
		"Barcelona":31,
		"Real Madrid":36,
		"Chivas":21,
		"Boca Juniors":31,
	}

	fmt.Println(campeonato)
	fmt.Println(" ")
	for equipo, puntaje := range campeonato {
		fmt.Println(equipo, puntaje)
	}
	

}