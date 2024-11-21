package internal

import (
	"errors"
	"time"
)

// La estructura de los datos es diferente en función de su periodicidad

type DatosHorario struct {
	Fecha         time.Time
	Hora          int
	Contaminantes map[string]Contaminante // Contaminantes mapea la magnitud del contaminante con su información detallada.
}

type DatosDiario struct {
	Fecha         time.Time
	Contaminantes map[string]Contaminante 
}

type DatosIrregular struct {
	FechaInicial  time.Time
	FechaFinal    time.Time
	Contaminantes map[string]Contaminante
}

// Rango de fechas válido
// fechaMinima y fechaMaxima se usan para garantizar que las fechas
// ingresadas sean razonables para datos de calidad del aire. Las mediciones
// suelen ser recientes y no deberían incluir fechas como el año 67.
var (
	fechaMinima = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC) // Primer año razonable de medición
	fechaMaxima = time.Now().Add(time.Hour)                   // Hasta 1 hora después del momento actual
)

// Constructor para DatosHorario
func NewDatosHorario(fecha time.Time, hora int, contaminantes []Contaminante) (DatosHorario, error) {
	// Validar rango de fechas razonable
	if fecha.Before(fechaMinima) || fecha.After(fechaMaxima) {
		return DatosHorario{}, errors.New("la fecha debe estar entre el año 2000 y el presente")
	}
	if hora < 1 || hora > 24 {
		return DatosHorario{}, errors.New("la hora debe estar entre 1 y 24")
	}
	if len(contaminantes) == 0 {
		return DatosHorario{}, errors.New("contaminantes no puede estar vacío")
	}

	contaminantesMap := make(map[string]Contaminante)
	for _, contaminante := range contaminantes {
		if contaminante.Magnitud == "" {
			return DatosHorario{}, errors.New("cada contaminante debe tener una magnitud")
		}
		if _, exists := contaminantesMap[contaminante.Magnitud]; exists {
			return DatosHorario{}, errors.New("no se pueden duplicar magnitudes en los contaminantes")
		}
		contaminantesMap[contaminante.Magnitud] = contaminante
	}

	return DatosHorario{Fecha: fecha, Hora: hora, Contaminantes: contaminantesMap}, nil
}

// Constructor para DatosDiario
func NewDatosDiario(fecha time.Time, contaminantes []Contaminante) (DatosDiario, error) {
	// Validar rango de fechas razonable
	if fecha.Before(fechaMinima) || fecha.After(fechaMaxima) {
		return DatosDiario{}, errors.New("la fecha debe estar entre el año 2000 y el presente")
	}
	if len(contaminantes) == 0 {
		return DatosDiario{}, errors.New("contaminantes no puede estar vacío")
	}

	contaminantesMap := make(map[string]Contaminante)
	for _, contaminante := range contaminantes {
		if contaminante.Magnitud == "" {
			return DatosDiario{}, errors.New("cada contaminante debe tener una magnitud")
		}
		if _, exists := contaminantesMap[contaminante.Magnitud]; exists {
			return DatosDiario{}, errors.New("no se pueden duplicar magnitudes en los contaminantes")
		}
		contaminantesMap[contaminante.Magnitud] = contaminante
	}

	return DatosDiario{Fecha: fecha, Contaminantes: contaminantesMap}, nil
}

// Constructor para DatosIrregular
func NewDatosIrregular(fechaInicial, fechaFinal time.Time, contaminantes []Contaminante) (DatosIrregular, error) {	
	// Validar rango de fechas razonable
	if fechaInicial.Before(fechaMinima) || fechaFinal.After(fechaMaxima) {
		return DatosIrregular{}, errors.New("las fechas deben estar entre el año 2000 y el presente")
	}
	if fechaFinal.Before(fechaInicial) {
		return DatosIrregular{}, errors.New("fecha final no puede ser anterior a la fecha inicial")
	}
	if len(contaminantes) == 0 {
		return DatosIrregular{}, errors.New("contaminantes no puede estar vacío")
	}

	contaminantesMap := make(map[string]Contaminante)
	for _, contaminante := range contaminantes {
		if contaminante.Magnitud == "" {
			return DatosIrregular{}, errors.New("cada contaminante debe tener una magnitud")
		}
		if _, exists := contaminantesMap[contaminante.Magnitud]; exists {
			return DatosIrregular{}, errors.New("no se pueden duplicar magnitudes en los contaminantes")
		}
		contaminantesMap[contaminante.Magnitud] = contaminante
	}

	return DatosIrregular{FechaInicial: fechaInicial, FechaFinal: fechaFinal, Contaminantes: contaminantesMap}, nil
}
