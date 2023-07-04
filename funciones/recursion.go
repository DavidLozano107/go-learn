package funciones

import "fmt"

func Exponencia(valor, exponente int ) {
	i:= 0
	if i > exponente {
		return
	}
	fmt.Println(valor)
	Exponencia(valor * valor, exponente)
}