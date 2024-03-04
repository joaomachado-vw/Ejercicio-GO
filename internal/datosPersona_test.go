package internal

import (
	"bytes"
	"testing"
)

func TestReadPersonData(t *testing.T) {
	tests := []struct {
		input    string
		expected Persona
	}{
		{"JOAO\nMACHADO\nH\n30\n123456789\n", Persona{"JOAO", "MACHADO", "H", 30, "123456789"}},
		{"LUCAS\nMACHADO\nh\n25\n987654321\n", Persona{"LUCAS", "MACHADO", "H", 25, "987654321"}},
		{"ISADORA\nTASCA\nM\n27\n456789123\n", Persona{"ISADORA", "TASCA", "M", 27, "456789123"}},
		{"MARIA\nSilva\nM\n62\n978894631\n", Persona{"MARIA", "SILVA", "M", 62, "978894631"}},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			persona := ReadPersonData()

			if persona != test.expected {
				t.Errorf("Esperado %+v, recibido %+v", test.expected, persona)
			}
		})
	}
}

func TestPersonaCollection_CheckDatos(t *testing.T) {
	tests := []struct {
		personas       []Persona
		expected       bool
		expectedOutput string
	}{
		{[]Persona{}, false, "No hay datos cadastrados!\n"},
		{[]Persona{{Nombre: "Joao"}}, true, ""},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			pc := &PersonaCollection{Personas: test.personas}
			writer := &bytes.Buffer{}

			result := pc.CheckDatos()

			if result != test.expected {
				t.Errorf("Esperado %t, recibido %t", test.expected, result)
			}

			if writer.String() != test.expectedOutput {
				t.Errorf("Esperado: %s, recibido: %s", test.expectedOutput, writer.String())
			}
		})
	}
}
