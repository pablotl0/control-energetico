package internal

import (
	"errors"
	"time"
)


//Para la localización precisa de la región específica
type Ubicacion struct {
	Provincia  string
	Municipio  string
	IDEstacion string
	Latitud    float64
	Longitud   float64
}

// Representación de un contaminante
type Contaminante struct {
	Magnitud      string
	Concentracion float64
	Unidad        string
	Cumplimiento  bool
}

// La estructura de los datos es diferente en función de su periodicidad

type DatosHorario struct {
	Fecha         time.Time
	Hora          int
	Contaminantes []Contaminante
}

type DatosDiario struct {
	Fecha         time.Time
	Contaminantes []Contaminante
}

type DatosIrregular struct {
	FechaInicial  time.Time
	FechaFinal    time.Time
	Contaminantes []Contaminante
}

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

// Constructor para Ubicacion
func NewUbicacion(provincia, municipio, idEstacion string, latitud, longitud float64) (Ubicacion, error) {
	if provincia == "" || municipio == "" || idEstacion == "" {
		return Ubicacion{}, errors.New("provincia, municipio e idEstacion son obligatorios")
	}
	if latitud < -90 || latitud > 90 {
		return Ubicacion{}, errors.New("latitud debe estar entre -90 y 90")
	}
	if longitud < -180 || longitud > 180 {
		return Ubicacion{}, errors.New("longitud debe estar entre -180 y 180")
	}
	return Ubicacion{Provincia: provincia, Municipio: municipio, IDEstacion: idEstacion, Latitud: latitud, Longitud: longitud}, nil
}

// Constructor para Contaminante
func NewContaminante(magnitud string, concentracion float64, unidad string, cumplimiento bool) (Contaminante, error) {
	if magnitud == "" || unidad == "" {
		return Contaminante{}, errors.New("magnitud y unidad son obligatorios")
	}
	if concentracion < 0 {
		return Contaminante{}, errors.New("la concentración no puede ser negativa")
	}
	return Contaminante{Magnitud: magnitud, Concentracion: concentracion, Unidad: unidad, Cumplimiento: cumplimiento}, nil
}

// Constructor para DatosHorario
func NewDatosHorario(fecha time.Time, hora int, contaminantes []Contaminante) (DatosHorario, error) {
	if hora < 1 || hora > 24 {
		return DatosHorario{}, errors.New("la hora debe estar entre 1 y 24")
	}
	if len(contaminantes) == 0 {
		return DatosHorario{}, errors.New("contaminantes no puede estar vacío")
	}
	return DatosHorario{Fecha: fecha, Hora: hora, Contaminantes: contaminantes}, nil
}

// Constructor para DatosDiario
func NewDatosDiario(fecha time.Time, contaminantes []Contaminante) (DatosDiario, error) {
	if len(contaminantes) == 0 {
		return DatosDiario{}, errors.New("contaminantes no puede estar vacío")
	}
	return DatosDiario{Fecha: fecha, Contaminantes: contaminantes}, nil
}

// Constructor para DatosIrregular
func NewDatosIrregular(fechaInicial, fechaFinal time.Time, contaminantes []Contaminante) (DatosIrregular, error) {
	if fechaFinal.Before(fechaInicial) {
		return DatosIrregular{}, errors.New("fecha final no puede ser anterior a la fecha inicial")
	}
	if len(contaminantes) == 0 {
		return DatosIrregular{}, errors.New("contaminantes no puede estar vacío")
	}
	return DatosIrregular{FechaInicial: fechaInicial, FechaFinal: fechaFinal, Contaminantes: contaminantes}, nil
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
