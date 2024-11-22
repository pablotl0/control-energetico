package internal

import (
	"errors"
)

// Representación de un contaminante
type Contaminante struct {
	Magnitud      string
	Concentracion float64
	Unidad        string
}

// Constructor para Contaminante
func NewContaminante(magnitud string, concentracion float64, unidad string) (Contaminante, error) {
	if magnitud == "" || unidad == "" {
		return Contaminante{}, errors.New("magnitud y unidad son obligatorios")
	}
	if concentracion < 0 {
		return Contaminante{}, errors.New("la concentración no puede ser negativa")
	}
	return Contaminante{Magnitud: magnitud, Concentracion: concentracion, Unidad: unidad}, nil
}