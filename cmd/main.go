package main

import (
	"fmt"
	"go.mod/internal"
	"os"
	"strings"
)

func main() {
	personaCollection := internal.PersonaCollection{}

	for {
		fmt.Println("Menú de opciones:\n")
		fmt.Println("1. Añadir internal")
		fmt.Println("2. Listar personas")
		fmt.Println("3. Buscar internal por DNI")
		fmt.Println("4. Listar personas por sexo")
		fmt.Println("5. Calcular media de edad")
		fmt.Println("6. Borrar internal por DNI")
		fmt.Println("0. Salir\n")

		var opcion int
		fmt.Print("Ingrese el número de la opción deseada: ")
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			// Añadir internal
			person := internal.ReadPersonData()
			personaCollection.AddPersona(person)
			fmt.Printf("%s %s, Edad: %d, Sexo: %s, DNI: %s\n", person.Nombre, person.Apellido, person.Edad, person.Sexo, person.DNI)
		case 2:
			// Listar todas las personas
			personaCollection.ListAllPeople()
		case 3:
			// Buscar internal por DNI/NIE
			if personaCollection.CheckDatos() == true {
				fmt.Print("Ingrese el DNI/NIE de la internal a buscar: ")
				var buscarDNI string
				fmt.Scan(&buscarDNI)

				foundPerson := personaCollection.FindPerson(buscarDNI)
				if foundPerson != nil {
					fmt.Printf("Persona encontrada: %s %s, Edad: %d, Sexo: %s\n", foundPerson.Nombre, foundPerson.Apellido, foundPerson.Edad, foundPerson.Sexo)
				} else {
					fmt.Println("Persona no encontrada.")
				}
			}
		case 4:
			// Listar personas por sexo
			fmt.Print("Ingrese el sexo a listar (H/M): ")
			var listarsSexo string
			fmt.Scan(&listarsSexo)

			personaCollection.ListPeopleBySex(strings.ToUpper(listarsSexo))
		case 5:
			// Calcular media de edad
			mediaEdad, mediaEdadSexo := personaCollection.CalculateAverageAge()
			fmt.Printf("Media de edad general: %.2f\n", mediaEdad)
			for sexo, mediaEdad := range mediaEdadSexo {
				fmt.Printf("Media de edad para el sexo %s: %.2f\n", sexo, mediaEdad)
			}
		case 6:
			// Borrar internal por DNI
			fmt.Print("Ingrese el DNI de la internal a borrar: ")
			var borrarDNI string
			fmt.Scan(&borrarDNI)
			personaCollection.DeletePerson(borrarDNI)
		case 0:
			// Salir
			fmt.Println("¡Adiós!")
			os.Exit(0)
		default:
			fmt.Println("Opción no válida. Inténtalo de nuevo.")
		}
	}
}
