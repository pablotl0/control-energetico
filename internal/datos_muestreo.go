package internal

import (
	"errors"
	"time"
)

// MuestraCalidadAire representa una entrada de datos de calidad del aire
type MuestraCalidadAire struct {
	FechaInicial  time.Time               // Fecha de inicio del muestreo 
	FechaFinal    time.Time               // Fecha de fin de muestreo 
	Contaminantes map[string]Contaminante // Contaminantes mapeados por su magnitud
}

// Rango de fechas válido
var (
	fechaMinima = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	fechaMaxima = time.Now().Add(time.Hour)
)

// Constructor para MuestraCalidadAire
func NewMuestraCalidadAire(fechaInicial, fechaFinal time.Time, contaminantes []Contaminante) (MuestraCalidadAire, error) {
	// Validar rango de fechas
	if fechaInicial.Before(fechaMinima) || fechaInicial.After(fechaMaxima) {
		return MuestraCalidadAire{}, errors.New("la fecha inicial debe estar entre el año 2000 y el presente")
	}
	if fechaFinal.Before(fechaInicial) || fechaFinal.After(fechaMaxima) {
		return MuestraCalidadAire{}, errors.New("la fecha final debe estar después de la inicial y dentro del rango válido")
	}

	// Validar contaminantes
	if len(contaminantes) == 0 {
		return MuestraCalidadAire{}, errors.New("contaminantes no puede estar vacío")
	}

	contaminantesMap := make(map[string]Contaminante)
	for _, contaminante := range contaminantes {
		if contaminante.Magnitud == "" {
			return MuestraCalidadAire{}, errors.New("cada contaminante debe tener una magnitud")
		}
		if _, exists := contaminantesMap[contaminante.Magnitud]; exists {
			return MuestraCalidadAire{}, errors.New("no se pueden duplicar magnitudes en los contaminantes")
		}
		contaminantesMap[contaminante.Magnitud] = contaminante
	}

	return MuestraCalidadAire{
		FechaInicial:  fechaInicial,
		FechaFinal:    fechaFinal,
		Contaminantes: contaminantesMap,
	}, nil
}
