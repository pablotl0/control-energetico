package internal

import (
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
