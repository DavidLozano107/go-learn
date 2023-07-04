package arregos_slices

import "fmt"

var tabla [10]int = [10]int{10, 0, 59}
var matriz [20][30]int 

func MuestroArreglos() {
	tabla[7] = 33
	tabla[2] = 54

	tabla2 := [10]string{"Pablo", "Juan"}

	fmt.Println(tabla)
	fmt.Println(tabla2)

	for i:=0; i<len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	fmt.Println("Nombre de tablas ")
	for i:=0; i<len(tabla2); i++ {
		fmt.Println((i+1),tabla2[i])
	}

}