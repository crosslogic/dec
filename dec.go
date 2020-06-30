// Package dec sirve para trabajar con decimales sin
// tener que preocuparse por la exactitud del punto floatante.
// Se crearon dos tipos D2 (dos decimales) y D4 (cuatro decimales)
package dec

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var (
	// SeparadorDecimal es el carácter que se usa para seprar los enteros
	// de los decimales en un número.
	SeparadorDecimal = ","

	// SeparadorDeMiles es el caracter que se usa para separar los millares.
	SeparadorDeMiles = "."
)

// ------------------------------- Conversion a texto -------------------------------//

// enPalabras es la funcion que devuelve en palabras un numero float.
// Esta funcion es implementada por D2 para hacer mas intiuitivo su uso.
func enPalabras(monto float64, moneda string) string { // 824,32
	var enterosFloat = math.Trunc(monto)         // 824
	var decsFloat = (monto - enterosFloat) * 100 // 0,3199999999998
	var decimales = int(toFixed(decsFloat, 0))   //32

	return moneda + " " + strings.Trim(aLetras(int(enterosFloat)), " ") + describirCentavos(decimales)
}
func correrComa(numero, lugares int) int {
	enFl := float64(numero) / math.Pow(10, float64(lugares))
	return int(math.Trunc(enFl))
}

// Devuelve el numero que se ingresa en letras.
func aLetras(n int) string {

	var concat string
	switch {

	case n < 100:
		return describirDecena(n)

	case n < 1000:
		concat += describirCentena(n)
		concat += describirDecena(n)
		return concat

	case n < 10000:
		corrido := correrComa(n, 3)
		concat += miles(corrido)
		concat += aLetras(n - corrido*1000)
		return concat

	case n < 100000:
		corrido := correrComa(n, 3)
		concat += describirDecena(corrido) + "MIL "
		concat += aLetras(n - corrido*1000)
		return concat
	case n < 1000*1000: //  542325

		fl := float64(n) / 1000            // me da 542,325
		centena := float64(math.Trunc(fl)) // me da 542
		miles := int(toFixed(centena, 0))
		cientos := n - (int(centena) * 1000)

		concat += describirCentena(miles)
		concat += describirDecena(miles)
		concat += "MIL "
		return concat + aLetras(cientos)

	case n < 1000*1000*1000: // mil millones  421.738.111,99
		fl := float64(n) / 1000000         // me da 421,73811199
		centena := float64(math.Trunc(fl)) // me da 421
		millones := int(toFixed(centena, 0))
		miles := n - (int(centena) * 1000000)

		concat += describirCentena(millones)
		concat += describirDecena(millones)
		concat += "MILLONES "
		return concat + aLetras(miles)

	}
	return ""
}
func miles(n int) string {
	switch n {
	case 1:
		return "MIL "
	}
	return unDigito(n) + "MIL "
}

// Describe desde 100 a 900
func describirCentena(n int) string {
	switch {
	case n < 100:
		return describirDecena(n)
	case n == 100:
		return "CIEN "
	case n < 200:
		return "CIENTO "
	case n < 300:
		return "DOSCIENTOS "
	case n < 400:
		return "TRESCIENTOS "
	case n < 500:
		return "CUATROCIENTOS "
	case n < 600:
		return "QUINIENTOS "
	case n < 700:
		return "SEISCIENTOS "
	case n < 800:
		return "SETECIENTOS "
	case n < 900:
		return "OCHOCIENTOS "
	case n < 1000:
		return "NOVECIENTOS "
	}
	return ""
}

// Describe desde 0 a 99
func describirDecena(n int) string {
	if n >= 100 { // 124 => Le saco el 100 y que quede 24
		fl := float64(n) / 100             // me da 1,24
		centena := float64(math.Trunc(fl)) // me da 1
		decena := (fl - centena) * 100     // (1,24 - 1) = 0,24 .         0,24 *100= 24
		n = int(toFixed(decena, 0))
	}

	switch {
	case n < 10:
		return unDigito(n)
	case n == 10:
		return "DIEZ "
	case n == 11:
		return "ONCE "
	case n == 12:
		return "DOCE "
	case n == 13:
		return "TRECE "
	case n == 14:
		return "CATORCE "
	case n == 15:
		return "QUINCE "
	case n == 16:
		return "DECISEIS "
	case n == 17:
		return "DIECISIETE "
	case n == 18:
		return "DIECIOCHO "
	case n == 19:
		return "DIECINUEVE "
	case n == 20:
		return "VEINTE "
	case n == 21:
		return "VEINTIUN "
	case n < 30:
		return "VEINTI" + unDigito(n-20)
	case n == 30:
		return "TREINTA "
	case n < 40:
		return "TREINTA Y " + unDigito(n-30)
	case n == 40:
		return "CUARENTA "
	case n < 50:
		return "CUARENTA Y " + unDigito(n-40)
	case n == 50:
		return "CINCUENTA "
	case n < 60:
		return "CINCUENTA Y " + unDigito(n-50)
	case n == 60:
		return "SESENTA "
	case n < 70:
		return "SESENTA Y " + unDigito(n-60)
	case n == 70:
		return "SETENTA "
	case n < 80:
		return "SETENTA Y " + unDigito(n-70)
	case n == 80:
		return "OCHENTA"
	case n < 90:
		return "OCHENTA Y " + unDigito(n-80)
	case n == 90:
		return "NOVENTA "
	case n < 100:
		return "NOVENTA Y " + unDigito(n-90)
	}
	return ""
}

func unDigito(n int) string {
	switch n {
	case 0:
		return ""
	case 1:
		return "UNO "
	case 2:
		return "DOS "
	case 3:
		return "TRES "
	case 4:
		return "CUATRO "
	case 5:
		return "CINCO "
	case 6:
		return "SEIS "
	case 7:
		return "SIETE "
	case 8:
		return "OCHO "
	case 9:
		return "NUEVE "
	}
	return ""
}

func describirCentavos(n int) string {
	if n == 0 {
		return ".-"
	}
	return " C/" + strconv.Itoa(n) + "/100.-"
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

//------------------------Aux-------------------------//

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func formatNumber(value float64, precision int, separadorMiles, separadorDecimal string) string {
	var x string
	x = fmt.Sprintf(fmt.Sprintf("%%.%df", precision), value)
	return formatNumberString(x, precision, separadorMiles, separadorDecimal)
}

func formatNumberString(x string, precision int, thousand string, decimal string) string {
	lastIndex := strings.Index(x, ".") - 1

	if lastIndex < 0 {
		lastIndex = len(x) - 1
	}

	var buffer []byte
	var strBuffer bytes.Buffer

	j := 0
	for i := lastIndex; i >= 0; i-- {
		j++
		buffer = append(buffer, x[i])

		if j == 3 && i > 0 && !(i == 1 && x[0] == '-') {
			buffer = append(buffer, ',')
			j = 0
		}
	}

	for i := len(buffer) - 1; i >= 0; i-- {
		strBuffer.WriteByte(buffer[i])
	}
	result := strBuffer.String()

	if thousand != "," {
		result = strings.Replace(result, ",", thousand, -1)
	}

	extra := x[lastIndex+1:]
	if decimal != "." {
		extra = strings.Replace(extra, ".", decimal, 1)
	}

	return result + extra
}
