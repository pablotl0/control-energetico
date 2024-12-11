package internal

import (
    "errors"
    "fmt"
)

type MuestraCalidadAire struct {
	Contaminantes map[string]Contaminante
}

func NewMuestraCalidadAire(fechaInicial Fecha, contaminantes []Contaminante) (*MuestraCalidadAire, error) {
	if len(contaminantes) == 0 {
		return nil, errors.New("contaminantes no puede estar vacío")
	}

	contaminantesMap := make(map[string]Contaminante)
	for _, contaminante := range contaminantes {
		if contaminante.MagnitudUnidad.Magnitud == "" {
			return nil, errors.New("cada contaminante debe tener una magnitud")
		}
		if _, exists := contaminantesMap[contaminante.MagnitudUnidad.Magnitud]; exists {
			return nil, errors.New("no se pueden duplicar magnitudes en los contaminantes")
		}
		contaminantesMap[contaminante.MagnitudUnidad.Magnitud] = contaminante
	}

	return &MuestraCalidadAire{
		FechaInicial:  *fechaInicial,
		Contaminantes: contaminantesMap,
	}, nil
}
