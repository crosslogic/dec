package dec

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Crear un nuevo D4 desde un float me genere un D4
func TestD4CrearFloat(t *testing.T) {
	d4 := NewD4(51.64239)
	assert.Equal(t, D4(516424), d4)

	otrod4 := NewD4(10)
	assert.Equal(t, D4(100000), otrod4)
}

// Si tengo un D4 con los cuatro decimales, que me lo pase a un float redondeado.
func TestD4AHumano(t *testing.T) {
	hum := D4(516499).Float()

	if hum != 51.6499 {
		fmt.Printf("%.80f\n", hum)
		t.Error("Me lo normalizó mal")
	}
}

func TestD4AMaquina(t *testing.T) {
	hum := NewD4(51.6423)

	if hum != 516423 {
		fmt.Println(hum)
		t.Error("Lo llevó mal a formato base de datos.")
	}
}

func TestD4UnmarshalJSON(t *testing.T) {

	type prueba struct {
		D4    D4
		Texto string
	}

	pruebas := []prueba{
		{
			D4:    3524,
			Texto: fmt.Sprintf("%v\n", `{"Monto":0.3524}`),
		},
		{
			D4:    -1,
			Texto: fmt.Sprintf("%v\n", `{"Monto":-0.0001}`),
		},
		{
			D4:    9999999999,
			Texto: fmt.Sprintf("%v\n", `{"Monto":999999.9999}`),
		},
		{
			D4:    0,
			Texto: fmt.Sprintf("%v\n", `{"Monto":0}`),
		},
	}
	for i, v := range pruebas {
		res := struct{ Monto D4 }{}
		err := json.NewDecoder(strings.NewReader(v.Texto)).Decode(&res)
		assert.Nil(t, err, "on index %v", i)
		assert.Equal(t, res.Monto, v.D4, "on index %v", i)
	}
}

func TestD4MarshalJSON(t *testing.T) {
	type item struct {
		Monto D4
	}
	type prueba struct {
		Item  item
		Texto string
	}
	pruebas := []prueba{
		{
			Item:  item{D4(3524)},
			Texto: fmt.Sprint(`{"Monto":0.3524}`, "\n"),
		},
		{
			Item:  item{D4(-1)},
			Texto: fmt.Sprint(`{"Monto":-0.0001}`, "\n"),
		},
		{
			Item:  item{D4(9999999999)},
			Texto: fmt.Sprint(`{"Monto":999999.9999}`, "\n"),
		},
		{
			Item:  item{D4(0)},
			Texto: fmt.Sprint(`{"Monto":0}`, "\n"),
		},
	}
	for _, v := range pruebas {
		w := strings.Builder{}
		err := json.NewEncoder(&w).Encode(v.Item)
		texto := w.String()
		assert.Nil(t, err)
		assert.Equal(t, v.Texto, texto)
	}
}
