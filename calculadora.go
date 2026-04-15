package main

import "fmt"

func main() {

	var a, b float64
	var op string

	fmt.Println("==== CALCULADORA ====")

	// Ingresar los datos
	fmt.Print("Ingrese el primer numero: ")
	fmt.Scan(&a)

	fmt.Print("Ingrese el segundo numero: ")
	fmt.Scan(&b)

	fmt.Print("Ingrese la operacion a ejecutar(+, -, *, /, ^, !): ")
	fmt.Scan(&op)

	// OPCIONES DE OPERACION
	switch op {

	case "+":
		fmt.Println("Resultado:", a+b)

	case "-":
		fmt.Println("Resultado:", a-b)

	case "*":
		fmt.Println("Resultado:", a*b)

	case "/":
		if b == 0 {
			fmt.Println("No se puede dividir para cero")
		} else {
			fmt.Println("Resultado:", a/b)
		}

	case "^":
		resultado := 1.0
		for i := 0; i < int(b); i++ {
			resultado = resultado * a
		}
		fmt.Println("Resultado:", resultado)

	case "!":
		if a < 0 {
			fmt.Println("No hay factorial de negativos")
		} else {
			resultado := 1.0
			for i := 1; i <= int(a); i++ {
				resultado = resultado * float64(i)
			}
			fmt.Println("Resultado:", resultado)
		}

	default:
		fmt.Println("Error ----operacion no valida----")
	}
}