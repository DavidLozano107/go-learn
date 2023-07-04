package variables

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)


var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func MuestroRestoVariables() {
	Nombre = "Pedro"
	Estado = true
	Sueldo = 1577.88
	Fecha = time.Now()

	fmt.Println(Nombre)
	fmt.Println(reflect.TypeOf(Nombre))
	fmt.Println(Estado)
	fmt.Println(reflect.TypeOf(Estado))
	fmt.Println(Sueldo)
	fmt.Println(reflect.TypeOf(Sueldo))
	fmt.Println(Fecha)
	fmt.Println(reflect.TypeOf(Fecha))

}

func ConviertoATexto(numero int)(bool, string) {
	texto := strconv.Itoa(numero)
	return true, texto
}