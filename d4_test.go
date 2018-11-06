package dec

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/pkg/errors"
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
		prueba{
			D4:    3524,
			Texto: fmt.Sprintf("%v\n", `{"Monto":0.3524}`),
		},
		prueba{
			D4:    -1,
			Texto: fmt.Sprintf("%v\n", `{"Monto":-0.0001}`),
		},
		prueba{
			D4:    9999999999,
			Texto: fmt.Sprintf("%v\n", `{"Monto":999999.9999}`),
		},
		prueba{
			D4:    0,
			Texto: fmt.Sprintf("%v\n", `{"Monto":0}`),
		},
	}
	for _, v := range pruebas {
		res := struct{ Monto D4 }{}
		err := json.NewDecoder(strings.NewReader(v.Texto)).Decode(&res)
		if err != nil {
			t.Error(errors.Wrapf(err, "al unmarshalizar %v", v.Texto))
		}

		if res.Monto != v.D4 {
			t.Fatal(t, fmt.Sprintf("No se leyó correctamente `%v`. Se esperaba: `%v`, pero se obtuvo `%v`.", v.Texto, v.D4, res.Monto))
		}
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
		prueba{
			Item:  item{D4(3524)},
			Texto: fmt.Sprint(`{"Monto":0.3524}`, "\n"),
		},
		prueba{
			Item:  item{D4(-1)},
			Texto: fmt.Sprint(`{"Monto":-0.0001}`, "\n"),
		},
		prueba{
			Item:  item{D4(9999999999)},
			Texto: fmt.Sprint(`{"Monto":999999.9999}`, "\n"),
		},
		prueba{
			Item:  item{D4(0)},
			Texto: fmt.Sprint(`{"Monto":0}`, "\n"),
		},
	}
	for _, v := range pruebas {
		w := strings.Builder{}
		err := json.NewEncoder(&w).Encode(v.Item)
		texto := w.String()
		if err != nil {
			t.Error(t, errors.Wrap(err, "no se pudo marsahlizar"))
		}
		if texto != v.Texto {
			t.Error(t, fmt.Sprintf("No se convirtió correctamente `%v`. Se esperaba: `%v`, pero se obtuvo `%v`.", v.Item, v.Texto, texto))
		}
	}
}
