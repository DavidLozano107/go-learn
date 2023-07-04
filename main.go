package main

import (
	"github.com/davidlozano107/go-desde-0/webserver"
)

func main() {
	// if os := runtime.GOOS; os == "Linux." {
	// 	fmt.Println("Si es linux")
	// } else {
	// 	fmt.Println("No es Linux es", os)
	// }

	// switch os := runtime.GOOS; os {
	// case "linux":
	// 	fmt.Println("es linux")
	// case "windows":
	// 	fmt.Println("es windows")
	// default:
	// 	fmt.Printf("%s \n", os)
	// }

	// numbero, texto := ejercicios.Ejercicio01Solucion("50")
	// fmt.Println(numbero)
	// fmt.Println(texto)

	// teclado.IngresoNumeros()
	
	// fmt.Println()

	// file.GrabaTabla()
	// file.SumaTabla()
	// file.LeoArchivo()
	// funciones.Calculos()

	// funciones.LlamarClosure()
	// funciones.Exponencia(2, 8)
	// arregosslices.MuestroSlice()
	//arregosslices.Capacidad()
	// mapas.MostrarMapas()

	// Pedro := new(modelos.Hombre)
	// Jimena := new(modelos.Mujer)

	// ejer_interfaces.HumanosRespirando(Pedro)
	// ejer_interfaces.HumanosRespirando(Jimena)

	// defer_panic.EjemploPanic()

	// canal1 :=  make(chan bool)

	// go goroutines.MiNombreLento("David Lozano", canal1)
	// <- canal1

	webserver.MiWebServer()

}