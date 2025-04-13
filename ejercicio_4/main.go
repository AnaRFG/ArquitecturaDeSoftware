package main

import (
	"fmt"
	"math/rand"
)

func main() {
	play()
}

func play() {
	numAleatorio := rand.Intn(100)
	var numIngresado int

	for cantIntentos := 0; cantIntentos < 10; cantIntentos++ {
		fmt.Printf("Ingrese un número (intentos restantes %d):", 10-cantIntentos)
		fmt.Scanln(&numIngresado)

		if numIngresado == numAleatorio {
			fmt.Println("Ganaste!")
			displayMenu()
			return
		} else if numIngresado > numAleatorio {
			fmt.Println("El número es menor")
		} else {
			fmt.Println("No acertaste, el número es mayor")
		}
	}
	fmt.Println("Se acabaron los intentos.El número es:", numAleatorio)
	displayMenu()
}

func displayMenu() {
	var option string

	fmt.Println("Desea jugar nuevamente? (s/n): ")
	fmt.Scanln(&option)

	switch option {
	case "s":
		play()
	case "n":
		fmt.Println("Gracias por jugar")
	default:
		fmt.Println("Opción incorrecta.Intente de nuevo")
		displayMenu()
	}
}
