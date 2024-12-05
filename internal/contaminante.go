package internal

import (
	"errors"
	"fmt"
)

type Contaminante struct {
	Nombre string
	Unidad string
}

func NewContaminante(nombre string, unidad string) (*Contaminante, error) {
	if nombre == "" || unidad == "" {
		return nil, errors.New("nombre y unidad son obligatorios")
	}
	return &Contaminante{
		Nombre: nombre,
		Unidad: unidad,
	}, nil
}

type MuestraContaminante struct {
	Contaminante  *Contaminante
	Concentracion float64
}

func NewMuestraContaminante(contaminante *Contaminante, concentracion float64) (*MuestraContaminante, error) {
	if contaminante == nil {
		return nil, errors.New("el contaminante no puede estar vacío")
	}
	if concentracion < 0 {
		return nil, errors.New("la concentración no puede ser negativa")
	}

	return &MuestraContaminante{
		Contaminante:  contaminante,
		Concentracion: concentracion,
	}, nil
}
