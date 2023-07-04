package ejercicios

import (
	"strconv"
)

func Ejercicio01Solucion(text string) (int, string) {
	
	var info string
	numero, _ := strconv.Atoi(text)

	if numero >= 100 {
		info = "Es mayor o igual a 100"
	} else {
		info = "Es menor a 100"
	}
	return numero, info
}