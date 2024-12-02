package internal

import (
	"errors"
	"time"
)

type Fecha struct {
	Año int
	Mes int
	Día int
}

func NewFecha(año, mes, día int) (*Fecha, error) {
	fechaActual := time.Now().UTC()
	fechaMinima := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	fechaMaxima := fechaActual.Add(-24 * time.Hour) 

	fecha := time.Date(año, time.Month(mes), día, 0, 0, 0, 0, time.UTC)

	if fecha.Before(fechaMinima) {
		return nil, errors.New("la fecha debe ser después del 1 de enero de 2000")
	}
	if fecha.After(fechaMaxima) {
		return nil, errors.New("la fecha no puede ser en el futuro")
	}

	return &Fecha{Año: año, Mes: mes, Día: día}, nil
}
