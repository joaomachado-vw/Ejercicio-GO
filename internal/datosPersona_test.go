package internal

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadPersonData(t *testing.T) {
	tests := []struct {
		typeCheck    string
		input        string
		expected     string
		inputEdad    int
		expectedEdad int
	}{
		{"Nombre: ", "JOAO\n", "JOAO", 0, 0},
		{"Apellido: ", "MACHADO\n", "MACHADO", 0, 0},
		{"Sexo: ", "H\n", "H", 0, 0},
		{"Edad: ", "30\n", "30", 30, 30},
		{"DNI: ", "Y8388671X\n", "Y8388671X", 0, 0},

		// {"LUCAS\nMACHADO\nh\n25\n987654321\n", Persona{"LUCAS", "MACHADO", "H", 25, "987654321"}},
		// {"ISADORA\nTASCA\nM\n27\n456789123\n", Persona{"ISADORA", "TASCA", "M", 27, "456789123"}},
		// {"MARIA\nSilva\nM\n62\n978894631\n", Persona{"MARIA", "SILVA", "M", 62, "978894631"}},
	}
	for _, test := range tests {
		t.Run("ValidName", func(t *testing.T) {
			resetInput := setInput(test.input)
			defer resetInput()
			if test.typeCheck == "Edad: " {
				edad := readEdad(test.typeCheck)
				require.Equal(t, test.expectedEdad, edad, "esperado '%s' recebido '%s'", test.expectedEdad, test.inputEdad)
			} else {
				nombre := readNombre(test.typeCheck)
				require.Equal(t, test.expected, nombre, "esperado '%s' recebido '%s'", test.expected, test.input)
			}

		})
	}

}

// func TestPersonaCollection_CheckDatos(t *testing.T) {
// 	tests := []struct {
// 		personas       []Persona
// 		expected       bool
// 		expectedOutput string
// 	}{
// 		{[]Persona{}, false, "No hay datos cadastrados!\n"},
// 		{[]Persona{{Nombre: "Joao"}}, true, ""},
// 	}

// 	for _, test := range tests {
// 		t.Run("", func(t *testing.T) {
// 			pc := &PersonaCollection{Personas: test.personas}
// 			writer := &bytes.Buffer{}

// 			result := pc.CheckDatos()

// 			if result != test.expected {
// 				t.Errorf("Esperado %t, recibido %t", test.expected, result)
// 			}

// 			if writer.String() != test.expectedOutput {
// 				t.Errorf("Esperado: %s, recibido: %s", test.expectedOutput, writer.String())
// 			}
// 		})
// 	}
// }

func setInput(input string) func() {
	old := os.Stdin
	r, w, _ := os.Pipe()
	os.Stdin = r

	w.Write([]byte(input))
	w.Close()

	return func() { os.Stdin = old }
}
