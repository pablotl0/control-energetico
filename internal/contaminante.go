	package internal

	import (
		"errors"
	)

	type MagnitudUnidad struct {
		Magnitud string
		Unidad   string
	}

	func NewMagnitudUnidad(magnitud, unidad string) (MagnitudUnidad, error) {
		if magnitud == "" || unidad == "" {
			return MagnitudUnidad{}, errors.New("magnitud y unidad son obligatorios")
		}
		return MagnitudUnidad{Magnitud: magnitud, Unidad: unidad}, nil
	}

	type Contaminante struct {
		MagnitudUnidad MagnitudUnidad
		Concentracion  float64
	}

	func NewContaminante(magnitud, unidad string, concentracion float64) (Contaminante, error) {
		if concentracion < 0 {
			return Contaminante{}, errors.New("la concentración no puede ser negativa")
		}

		mu, err := NewMagnitudUnidad(magnitud, unidad)
		if err != nil {
			return Contaminante{}, err
		}

		return Contaminante{MagnitudUnidad: mu, Concentracion: concentracion}, nil
	}
