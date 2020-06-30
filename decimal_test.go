package dec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToFixed(t *testing.T) {
	f1 := toFixed(1, 0)
	if f1 != 1 {
		t.Error("Mal convertido to fixed",
			"\nEsperado: ", 1,
			"\nObtenido: ", f1)
	}

	f2 := toFixed(10000, 0)
	if f2 != 10000 {
		t.Error("Mal convertido to fixed",
			"\nEsperado: ", 10000,
			"\nObtenido: ", f2)
	}

	f3 := toFixed(10.63, 2)
	if f3 != 10.63 {
		t.Error("Mal convertido to fixed",
			"\nEsperado: ", 1063,
			"\nObtenido: ", f3)
	}

	f4 := toFixed(10.6399, 2)
	if f4 != 10.64 {
		t.Error("Mal convertido to fixed",
			"\nEsperado: ", 10.64,
			"\nObtenido: ", f4)
	}

	f5 := toFixed(51.64239, 2)
	if f5 != 51.64 {
		t.Error("Mal convertido to fixed",
			"\nEsperado: ", 51.64,
			"\nObtenido: ", f5)
	}
}

func TestDescribirDecena(t *testing.T) {
	var letras string

	letras = describirDecena(15)
	if letras != "QUINCE " {
		t.Error("Dio: " + letras)
	}

	letras = describirDecena(118)
	if letras != "DIECIOCHO " {
		t.Error("Dio: " + letras)
	}

	letras = describirDecena(90)
	if letras != "NOVENTA " {
		t.Error("Dio: " + letras)
	}

	letras = describirDecena(44)
	if letras != "CUARENTA Y CUATRO " {
		t.Error("Dio: " + letras)
	}
}

func TestALetras(t *testing.T) {

	type datos struct {
		N        float64
		Expected string
	}

	testData := []datos{
		{
			N:        3,
			Expected: "PESOS TRES.-",
		},
		{
			N:        12,
			Expected: "PESOS DOCE.-",
		},
		{
			N:        122,
			Expected: "PESOS CIENTO VEINTIDOS.-",
		},
		{
			N:        824.32,
			Expected: "PESOS OCHOCIENTOS VEINTICUATRO C/32/100.-",
		},
		{
			N:        1023,
			Expected: "PESOS MIL VEINTITRES.-",
		},
		{
			N:        9003,
			Expected: "PESOS NUEVE MIL TRES.-",
		},
		{
			N:        14889,
			Expected: "PESOS CATORCE MIL OCHOCIENTOS OCHENTA Y NUEVE.-",
		},
		{
			N:        18777.91,
			Expected: "PESOS DIECIOCHO MIL SETECIENTOS SETENTA Y SIETE C/91/100.-",
		},
		{
			N:        738111.99,
			Expected: "PESOS SETECIENTOS TREINTA Y OCHO MIL CIENTO ONCE C/99/100.-",
		},
		{
			N:        421738111.99,
			Expected: "PESOS CUATROCIENTOS VEINTIUN MILLONES SETECIENTOS TREINTA Y OCHO MIL CIENTO ONCE C/99/100.-",
		},
		{
			N:        999999999.99,
			Expected: "PESOS NOVECIENTOS NOVENTA Y NUEVE MILLONES NOVECIENTOS NOVENTA Y NUEVE MIL NOVECIENTOS NOVENTA Y NUEVE C/99/100.-",
		},
	}

	for _, v := range testData {
		actual := enPalabras(v.N, "PESOS")
		assert.Equal(t, v.Expected, actual)
	}
}

func TestCorrerComa(t *testing.T) {
	out := correrComa(543, 2)
	assert.Equal(t, 5, out)

	{
		out := correrComa(543, -2)
		assert.Equal(t, 54300, out)
	}
}
