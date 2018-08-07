package dec

import (
	"fmt"
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
