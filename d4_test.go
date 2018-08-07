package dec

import (
	"fmt"
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
