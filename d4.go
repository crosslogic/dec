package dec

import (
	"database/sql/driver"

	"github.com/pkg/errors"
)

// D4 sirve para registros que necesita los 4 decimales (precios unitarios,
// descuentos, tipos de cambio, etc.)
type D4 int64

// NewD4 redondea a cuatro decimales y le corre la coma para contabilizarlo como corresponde.
func NewD4(n float64) D4 {
	porMil := toFixed(n*10000, 0)
	d := D4(porMil)
	return d
}

// StringDos devuelve 334.897,98
func (d D4) StringDos() string {

	f := float64(d) / 10000
	if f == float64(0) {
		return "-"
	}
	return formatNumber(f, 2)
}

// Float devuelve el número como float con la coma donde corresponde.
func (d D4) Float() float64 {
	return float64(d) / 10000
}

// String devuelve 334.897,9815
func (d D4) String() string {

	f := float64(d) / 10000
	if f == float64(0) {
		return "-"
	}
	return formatNumber(f, 4)
}

// Value satisface la interface de package sql, para persistir el núnero como debe.
func (d D4) Value() (driver.Value, error) {
	return int64(d), nil
}

// Scan satisface la interfaz de package sql, para correr las comas los lugares
// necesarios.
func (d *D4) Scan(value interface{}) error {
	if value == nil {
		*d = D4(0)
		return nil
	}
	entero, ok := value.(int64)
	if !ok {
		return errors.Errorf("al intentar Scan en un D2. Se esperaba un int, se obtuvo un %t", value)
	}
	*d = D4(entero)

	return nil
}