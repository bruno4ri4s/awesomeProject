package main

import (
	"fmt"

	"rsc.io/quote"
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

}
