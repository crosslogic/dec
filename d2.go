package dec

import (
	"database/sql/driver"
	"strconv"

	"github.com/pkg/errors"
)

// D2 sirve registros contables, tiene dos decimales.
type D2 int64

// NewD2 toma el float y lo redondea a dos decimales
func NewD2(n float64) D2 {
	porCien := toFixed(n*100, 0) // 1774.0
	d := D2(porCien)
	return d
}

// Float devuelve el número como float con la coma donde corresponde.
func (d D2) Float() float64 {
	return float64(d) / 100
}

// String devuelve 334.897,98
func (d D2) String() string {

	f := float64(d) / 100
	if f == float64(0) {
		return "-"
	}
	return formatNumber(f, 2, SeparadorDeMiles, SeparadorDecimal)

}

// EnPalabras devuelve PESOS TRESCIENTOS CUARENTA Y CINCO C/20/100.-
func (d D2) EnPalabras(moneda string) string {
	enFloat := d.Float()
	return enPalabras(enFloat, moneda)
}

// Value satisface la interface de package sql, para persistir el núnero como debe.
func (d D2) Value() (driver.Value, error) {
	return int64(d), nil
}

// Scan satisface la interfaz de package sql, para correr las comas los lugares
// necesarios.
func (d *D2) Scan(value interface{}) error {
	if value == nil {
		*d = D2(0)
		return nil
	}
	entero, ok := value.(int64)
	if !ok {
		return errors.Errorf("al intentar Scan en un D2. Se esperaba un int64, se obtuvo un %t", value)
	}
	*d = D2(entero)

	return nil
}

// MarshalJSON es para tomar un D2 y pasarlo a JSON.
func (d D2) MarshalJSON() (by []byte, err error) {
	by = []byte(strconv.FormatFloat(d.Float(), 'f', -1, 64))
	return by, nil
}

// UnmarshalJSON Es para pasar un Fecha => JSON
func (d *D2) UnmarshalJSON(input []byte) error {
	texto := string(input)

	if texto == "null" || texto == `""` {
		*d = 0
		return nil
	}

	fl, err := strconv.ParseFloat(texto, 64)
	if err != nil {
		return errors.Errorf("no se pudo convertir %v a float64", texto)
	}
	*d = NewD2(fl)

	return nil
}

// ExportarParaCSV sirve para cuando se debe generar un string con el número
// generalmete para archivos CSV.
func (d D2) ExportarParaCSV(
	cantidadDecimales int,
	separadorMiles string,
	separadorDecimal string,
	largo int,
	llenarCon string,
	alineadoDerecha bool) (rtdo string) {

	texto := formatNumber(float64(d)/100.0, cantidadDecimales, separadorMiles, separadorDecimal)

	// Si el texto es más largo, lo recorto
	if len(texto) > largo {
		return texto[:largo]
	}

	// Si el más corto, relleno
	cantidadAgregar := largo - len(texto)
	relleno := ""
	for i := 1; i <= cantidadAgregar; i++ {
		relleno += llenarCon
	}

	if alineadoDerecha {
		rtdo = relleno + texto
		return
	}
	if !alineadoDerecha {
		rtdo = texto + relleno
		return
	}
	return
}
