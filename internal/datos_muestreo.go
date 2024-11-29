package internal

import (
	"errors"
	"time"
)

type MuestraCalidadAire struct {
	FechaInicial  time.Time               
	Contaminantes map[string]Contaminante 
}

var (
	fechaMinima = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	fechaMaxima = time.Now().Add(time.Hour)
)

func NewMuestraCalidadAire(fechaInicial time.Time, contaminantes []Contaminante) (MuestraCalidadAire, error) {
	fechaInicial = time.Date(fechaInicial.Year(), fechaInicial.Month(), fechaInicial.Day(), 0, 0, 0, 0, time.UTC)

	if fechaInicial.Before(fechaMinima) || fechaInicial.After(fechaMaxima) {
		return MuestraCalidadAire{}, errors.New("la fecha inicial debe estar entre el año 2000 y el presente")
	}


	if len(contaminantes) == 0 {
		return MuestraCalidadAire{}, errors.New("contaminantes no puede estar vacío")
	}

	contaminantesMap := make(map[string]Contaminante)
	for _, contaminante := range contaminantes {
		if contaminante.MagnitudUnidad.Magnitud == "" {
			return MuestraCalidadAire{}, errors.New("cada contaminante debe tener una magnitud")
		}
		if _, exists := contaminantesMap[contaminante.MagnitudUnidad.Magnitud]; exists {
			return MuestraCalidadAire{}, errors.New("no se pueden duplicar magnitudes en los contaminantes")
		}
		contaminantesMap[contaminante.MagnitudUnidad.Magnitud] = contaminante
	}

	return MuestraCalidadAire{
		FechaInicial:  fechaInicial,
		Contaminantes: contaminantesMap,
	}, nil
}
