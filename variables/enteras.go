package variables

import (
	"fmt"
	"reflect"
)

func MuestroEnteros() {
	var intComun int
	intDe32 := int32(10)
	fmt.Println("intcomun  = ", intComun)
	fmt.Println(reflect.TypeOf(intComun))
	fmt.Println("intDe32  = ", intDe32)
	fmt.Println(reflect.TypeOf(intDe32))
}