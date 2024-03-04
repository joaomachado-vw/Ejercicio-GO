package internal

import (
	"fmt"
	"sort"
)

type Persona struct {
	Nombre   string
	Apellido string
	Sexo     string
	Edad     int
	DNI      string
}

// PersonCollection es una colección de personas.
type PersonaCollection struct {
	Personas []Persona
}

func (pc *PersonaCollection) AddPersona(persona Persona) {
	pc.Personas = append(pc.Personas, persona)
}

// ListAllPeople lista todas las personas ordenadas por edad.
func (pc *PersonaCollection) ListAllPeople() {
	//Checkear si hay personas cadastradas y Ordenar personas por edad
	if pc.CheckDatos() == true {
		sort.Slice(pc.Personas, func(i, j int) bool {
			return pc.Personas[i].Edad < pc.Personas[j].Edad
		})

		//Mostrar personas
		fmt.Println("Listado de personas:")
		for _, p := range pc.Personas {
			fmt.Printf("%s %s, Edad: %d, Sexo: %s, DNI: %s\n", p.Nombre, p.Apellido, p.Edad, p.Sexo, p.DNI)
		}
	}
}

// FindPerson busca a una internal por su DNI.
func (pc *PersonaCollection) FindPerson(dni string) *Persona {

	for _, p := range pc.Personas {
		if p.DNI == dni {
			return &p
		}
	}

	return nil
}

// ListPeopleBySex lista todas las personas de un sexo específico.
func (pc *PersonaCollection) ListPeopleBySex(sexo string) {
	if pc.CheckDatos() == true {
		fmt.Printf("Listado de personas de sexo %s:\n", sexo)
		for _, p := range pc.Personas {
			if p.Sexo == sexo {
				fmt.Printf("%s %s, Edad: %d, DNI: %s\n", p.Nombre, p.Apellido, p.Edad, p.DNI)
			}
		}
	}
}

// CalculateAverageAge calcula la edad media general y por sexos.
func (pc *PersonaCollection) CalculateAverageAge() (float64, map[string]float64) {
	pc.CheckDatos()
	totalEdad := 0
	contador := 0
	edadPorSexo := make(map[string]int)
	contadorPorSexo := make(map[string]int)

	for _, p := range pc.Personas {
		totalEdad += p.Edad
		contador++
		edadPorSexo[p.Sexo] += p.Edad
		contadorPorSexo[p.Sexo]++
	}

	edadMedia := float64(totalEdad) / float64(contador)
	edadMediaSexo := make(map[string]float64)

	for sexo, edad := range edadPorSexo {
		edadMediaSexo[sexo] = float64(edad) / float64(contadorPorSexo[sexo])
	}

	return edadMedia, edadMediaSexo
}

// DeletePerson elimina a una internal por su DNI.
func (pc *PersonaCollection) DeletePerson(dni string) {
	if pc.CheckDatos() == true {
		for i, p := range pc.Personas {
			if p.DNI == dni {
				// Eliminar internal encontrada
				pc.Personas = append(pc.Personas[:i], pc.Personas[i+1:]...)
				fmt.Printf("Persona con DNI %s eliminada.\n", dni)
				return
			}
		}
		fmt.Printf("Persona con DNI %s no encontrada.\n", dni)
	}
}
