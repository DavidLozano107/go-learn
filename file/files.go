package file

import (
	"bufio"
	"fmt"
	"os"

	"github.com/davidlozano107/go-desde-0/ejercicios"
)

var fileName string = "./file/txt/tabla.txt"

func GrabaTabla() {
	var texto string = ejercicios.Ejercicio02Solucion()
	archivo, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error al crear el archivo", err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla() {
	texto := ejercicios.Ejercicio02Solucion()
	if !Append(fileName, texto) {
		fmt.Println("Error al concatenar contenido")
	}
}

func Append(fileN string,texto string ) bool {
	arch, err := os.OpenFile(fileN, os.O_WRONLY | os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error durante el Appends", err.Error())
		return false
	}

	_, err = arch.WriteString(fmt.Sprintln(""))
	_, err = arch.WriteString(texto)

	if err != nil {
		fmt.Println("Error durante el WriteString", err.Error())
		return false
	}

	arch.Close()
	return true

}

func LeoArchivo() {

	// archivo, err := ioutil.ReadFile(fileName)
	archivo, err := os.Open(fileName)
	
	if err != nil {
		fmt.Println("Error al leer archivo", err.Error())
		return
	}

	// fmt.Println(string(archivo))
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println(">", registro)
	}

	archivo.Close()

}