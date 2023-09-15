package main

import (
	"fmt"

	"rsc.io/quote"
)

const PI = 3.14

const (
	x = 100
	y = 0b1010 // binario
	z = 0o12   // Octal
	w = 0xFF   // Hexadecimal
)

const (
	Domingo = iota + 1
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func main() {
	fmt.Println("Hola Mundo")
	fmt.Println(quote.Hello())

	//Declaracion de variables
	//	var firstName, lastName string
	//	var age int

	var (
		firstName, lastName string
		age                 int
	)

	firstName = "Alex"
	lastName = "Rodriguez"
	age = 27

	var id, middleName = 1, "Santiago"

	tel := 44231123
	tel = 34231234

	fmt.Println(" id: ", id, " nombre: ", firstName, " segundo nombre: ", middleName, " apellido: ", lastName, " edad: ", age, " telefono: ", tel)
	fmt.Println(x, y, z, w)
	fmt.Println(Viernes)

}
