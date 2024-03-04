package internal

import (
	"fmt"
	"github.com/go-playground/validator"
	"regexp"
	"strings"
)

func readNombre(texto string) string {
	var nom string
	for {
		fmt.Print(texto)

		fmt.Scan(&nom)
		strings.ToUpper(nom)
		if texto == "Sexo: " {
			validator := validator.New()
			err := validator.Var(nom, "oneof=H M h m")
			if err != nil {
				fmt.Println("Only H or M accepted")
			} else {
				break
			}
		} else if texto == "DNI: " && validarDocumento(nom) == false {
			fmt.Println("Insert a valid DNI")
		} else {
			break
		}
	}

	return nom
}

func readEdad(texto string) int {
	var edad int
	for {
		fmt.Print(texto)
		fmt.Scan(&edad)
		validator := validator.New()
		err := validator.Var(edad, "gte=0")
		if err == nil {
			break
		}
		fmt.Println("Not a valid age")
	}
	return edad
}

func ReadPersonData() Persona {

	fmt.Println("\nIngrese los datos de la internal:")

	nombre := readNombre("Nombre: ")
	apellido := readNombre("Apellido: ")
	sexo := readNombre("Sexo: ")
	edad := readEdad("Edad: ")
	dni := readNombre("DNI: ")

	return Persona{
		Nombre:   nombre,
		Apellido: apellido,
		Sexo:     strings.ToUpper(sexo),
		Edad:     edad,
		DNI:      dni,
	}
}

func (pc *PersonaCollection) CheckDatos() bool {
	if len(pc.Personas) == 0 {
		fmt.Println("No hay datos cadastrados!")
		return false
	}
	return true
}

func validarDocumento(documento string) bool {
	documento = strings.ToUpper(documento)
	var regex *regexp.Regexp

	if strings.HasPrefix(documento, "X") || strings.HasPrefix(documento, "Y") || strings.HasPrefix(documento, "Z") {
		regex = regexp.MustCompile(`^[XYZ][0-9]{7}[TRWAGMYFPDXBNJZSQVHLCKE]$`)
	} else {
		regex = regexp.MustCompile(`^[0-9]{8}[TRWAGMYFPDXBNJZSQVHLCKE]$`)
	}

	return regex.MatchString(documento)
}
