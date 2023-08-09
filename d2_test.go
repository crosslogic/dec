package dec

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func TestD2CrearFloat(t *testing.T) {
	d2 := NewD2(51.64239)
	//hum := d4.M()

	if d2 != D2(5164) {
		t.Error(fmt.Sprint("Mal creado el D2\n",
			"Esperado: ", 5164, "\n",
			"Obtenido: ", int(d2)))
	}
}

// Si tengo un D4 con los cuatro decimales, que me lo pase a un float redondeando
func TestD2AHumano(t *testing.T) {
	hum := D2(516499).Float()
	if hum != 5164.99 {
		fmt.Printf("%.80f\n", hum)
		t.Error("Me lo normalizó mal")
	}
}

func TestD2AMaquina(t *testing.T) {
	hum := NewD2(51.6423)
	if hum != 5164 {
		fmt.Println(hum)
		t.Error("Lo llevó mal a formato base de datos.")
	}
}

func TestSumarD2(t *testing.T) {
	x := NewD2(17.74)
	y := NewD2(119.66)
	total := x + y
	if total != 13740 {
		fmt.Println("Sumando: ", x, " + ", y, " me dio: ", total, ". Diferencia ", total-13740)
		t.Error("No redondeo como debia.")
	}
}

func TestD2UnmarshalJSON(t *testing.T) {

	type prueba struct {
		D2    D2
		Texto string
	}

	pruebas := []prueba{
		{
			D2:    3524,
			Texto: fmt.Sprintf("%v\n", `{"Monto":35.24}`),
		},
		{
			D2:    -1,
			Texto: fmt.Sprintf("%v\n", `{"Monto":-0.01}`),
		},
		{
			D2:    9999999999,
			Texto: fmt.Sprintf("%v\n", `{"Monto":99999999.99}`),
		},
		{
			D2:    0,
			Texto: fmt.Sprintf("%v\n", `{"Monto":0}`),
		},
	}
	for _, v := range pruebas {
		res := struct{ Monto D2 }{}
		err := json.NewDecoder(strings.NewReader(v.Texto)).Decode(&res)
		if err != nil {
			t.Error(fmt.Errorf("al unmarshalizar %v", v.Texto))
		}

		if res.Monto != v.D2 {
			t.Fatalf("no se leyó correctamente `%v`. Se esperaba: `%v`, pero se obtuvo `%v`.", v.Texto, v.D2, res.Monto)
		}
	}
}

func TestD2MarshalJSON(t *testing.T) {
	type item struct {
		Monto D2
	}
	type prueba struct {
		Item  item
		Texto string
	}
	pruebas := []prueba{
		prueba{
			Item:  item{D2(3524)},
			Texto: fmt.Sprint(`{"Monto":35.24}`, "\n"),
		},
		prueba{
			Item:  item{D2(-1)},
			Texto: fmt.Sprint(`{"Monto":-0.01}`, "\n"),
		},
		prueba{
			Item:  item{D2(9999999999)},
			Texto: fmt.Sprint(`{"Monto":99999999.99}`, "\n"),
		},
		prueba{
			Item:  item{D2(0)},
			Texto: fmt.Sprint(`{"Monto":0}`, "\n"),
		},
	}
	for _, v := range pruebas {
		w := strings.Builder{}
		err := json.NewEncoder(&w).Encode(v.Item)
		texto := w.String()
		if err != nil {
			t.Error(t, fmt.Errorf("no se pudo marsahlizar: %w", err))
		}
		if texto != v.Texto {
			t.Error(t, fmt.Sprintf("No se convirtió correctamente `%v`. Se esperaba: `%v`, pero se obtuvo `%v`.", v.Item, v.Texto, texto))
		}
	}
}
