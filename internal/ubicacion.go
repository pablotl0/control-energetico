package internal

import (
	"errors"
)

//Para la localización precisa de la región específica
type Ubicacion struct {
	Provincia  string
	Municipio  string
	IDEstacion string
	Latitud    float64
	Longitud   float64
}

// Constructor para Ubicacion
func NewUbicacion(provincia, municipio, idEstacion string, latitud, longitud float64) (Ubicacion, error) {
	if provincia == "" || municipio == "" || idEstacion == "" {
		return Ubicacion{}, errors.New("provincia, municipio e idEstacion son obligatorios")
	}
	if latitud < -90 || latitud > 90 {
		return Ubicacion{}, errors.New("latitud debe estar entre -90 y 90")
	}
	if longitud < -180 || longitud > 180 {
		return Ubicacion{}, errors.New("longitud debe estar entre -180 y 180")
	}
	return Ubicacion{Provincia: provincia, Municipio: municipio, IDEstacion: idEstacion, Latitud: latitud, Longitud: longitud}, nil
}