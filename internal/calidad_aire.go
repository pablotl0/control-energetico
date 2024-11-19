package internal

import (
	"errors"

)


// Es la estructura de datos principal
type DatosCalidadAire struct {
	Ubicacion          Ubicacion
	MetodoEvaluacion   string
	FrecuenciaMuestreo string
	DatosHorario       []DatosHorario
	DatosDiario        []DatosDiario
	DatosIrregular     []DatosIrregular
	Fuente             FuenteDatos
}


// Constructor para DatosCalidadAire
func NewDatosCalidadAire(ubicacion Ubicacion, metodoEvaluacion, frecuenciaMuestreo string, datosHorario []DatosHorario, datosDiario []DatosDiario, datosIrregular []DatosIrregular, fuente FuenteDatos) (DatosCalidadAire, error) {
	if metodoEvaluacion == "" || frecuenciaMuestreo == "" {
		return DatosCalidadAire{}, errors.New("metodoEvaluacion y frecuenciaMuestreo son obligatorios")
	}
	if len(datosHorario) == 0 && len(datosDiario) == 0 && len(datosIrregular) == 0 {
		return DatosCalidadAire{}, errors.New("debe haber al menos un tipo de datos de muestreo")
	}
	return DatosCalidadAire{Ubicacion: ubicacion, MetodoEvaluacion: metodoEvaluacion, FrecuenciaMuestreo: frecuenciaMuestreo, DatosHorario: datosHorario, DatosDiario: datosDiario, DatosIrregular: datosIrregular, Fuente: fuente}, nil
}
