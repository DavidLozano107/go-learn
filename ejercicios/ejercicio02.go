package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Ejercicio02Solucion() string {
	var numero int
	var err error
	var texto string

	scanner := bufio.NewScanner(os.Stdin)

	
	for {
		fmt.Println("Ingresa un numero")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err == nil {
				break
			}
			fmt.Println("El dato que ingresaste es incorrecto.")
		}
	}

	fmt.Println("El numero digitado es:", numero)
	
	for i := 0; i <= 10; i++ {
		texto += fmt.Sprintf("%d x %d = %d \n", numero, i, (numero * i))
	}
		
	return texto
}

