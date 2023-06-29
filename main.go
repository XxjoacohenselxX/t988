package main

import "fmt"

func main() {
	var numeros [10]int

	for i := 0; i < 10; i++ {
		fmt.Print("Ingrese un número entero: ")
		fmt.Scanln(&numeros[i])
	}

	min := buscarMinimo(numeros)
	max := buscarMaximo(numeros)
	promedio := calcularPromedio(numeros)
	sumatoria := calcularSumatoria(numeros)

	fmt.Println("Mínimo: ", min)
	fmt.Println("Máximo: ", max)
	fmt.Println("Promedio: ", promedio)
	fmt.Println("Sumatoria: ", sumatoria)

}

func buscarMinimo(nums [10]int) (minimo int) {
	minimo = nums[0]

	for i := 1; i < 10; i++{
		if nums[i] < minimo {
			minimo = nums[i]
		}
	}	

	return
}

func buscarMaximo(nums [10]int) (maximo int) {
	maximo = nums[0]

	//TODO completar
	return
}

func calcularPromedio(nums [10]int) (promedio float64) {

	//TODO completar
	return
}

func calcularSumatoria(nums [10]int) (sumatoria float64) {
	//TODO completar
	return
}
