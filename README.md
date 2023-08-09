# Dec

Package dec sirve para trabajar con decimales sin tener que preocuparse por la exactitud del punto flotante en la base de datos.

## Tipos

- **D2:** usa dos decimales, ideal para trabajar con registros contables.
- **D4:** cuatro decimales, para otros números que necesitan más precisión.

## Funciones

Se agregaron dos funciones adicionales:

- **EnPalabras:** devuelve el número escrito en palabras (ideal para recibos)
- **ExportarParaCSV:** devuelve un string configurable para generar archivos planos.

## Bases de datos y JSON

Al funcionar con un int64 subyacente, no hay problemas para Marshal ni Unmarshal.

## Dependencias

Contribuciones bienvenidas.
