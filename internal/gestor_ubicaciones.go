package internal

import "errors"

type GestorUbicaciones struct {
	ubicaciones map[string]Ubicacion 
}

func NewGestorUbicaciones(ubicaciones []Ubicacion) (GestorUbicaciones, error) {
	if len(ubicaciones) == 0 {
		return GestorUbicaciones{}, errors.New("debe haber al menos una ubicación")
	}

	mapUbicaciones := make(map[string]Ubicacion)
	for _, u := range ubicaciones {
		if _, exists := mapUbicaciones[u.Provincia]; exists {
			return GestorUbicaciones{}, errors.New("ubicaciones duplicadas para la misma provincia")
		}
		mapUbicaciones[u.Provincia] = u
	}

	return GestorUbicaciones{ubicaciones: mapUbicaciones}, nil
}
