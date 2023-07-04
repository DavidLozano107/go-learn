package arregos_slices

import "fmt"

var tablaS []int = []int{2,5,6}
var arreglo [10]int = [10]int{6,1,2,3,4,56,6,7,8,10}

func MuestroSlice() {
	fmt.Println(tablaS)

	// Slice creado desde un vector
	porcion := arreglo[3:] 
	porcion2 := arreglo[:5]
	porcion3 := arreglo[3:6]

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)

}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Println(elementos)
	

	nums := make([]int, 0)

	for i:= 0; i<100; i++ {
		nums = append(nums, (i+1))
	}

	for i:= 0; i<len(nums); i++ {
		fmt.Println(nums[i])
	}

	fmt.Printf("Largo %d, Capacidad %d \n", len(nums), cap(nums))

}