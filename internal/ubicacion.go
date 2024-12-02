package internal

import (
	"errors"
)

type Ubicacion struct {
	Provincia  string
	Latitud    float64
	Longitud   float64
}

func NewUbicacion(provincia string, municipio string, latitud float64, longitud float64) (*Ubicacion, error) {
	if provincia == ""   {
		return nil, errors.New("provincia, municipio e idEstacion son obligatorios")
	}
	if latitud < -90 || latitud > 90 {
		return nil, errors.New("latitud debe estar entre -90 y 90")
	}
	if longitud < -180 || longitud > 180 {
		return nil, errors.New("longitud debe estar entre -180 y 180")
	}
	return &Ubicacion{Provincia: provincia,  Latitud: latitud, Longitud: longitud}, nil
}