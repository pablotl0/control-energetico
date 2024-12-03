package internal

import "errors"

type GestorUbicaciones struct {
	ubicaciones map[string][]Ubicacion
}

func NewGestorUbicaciones(ubicaciones []Ubicacion) (*GestorUbicaciones, error) {
	if len(ubicaciones) == 0 {
		return nil, errors.New("debe haber al menos una ubicación")
	}

	mapUbicaciones := make(map[string][]Ubicacion)

	for _, u := range ubicaciones {
		provincia := u.Provincia
		municipio := u.Municipio

		ubicacionesProvincia, exists := mapUbicaciones[provincia]
		if exists {
			for _, ubicacionExistente := range ubicacionesProvincia {
				if ubicacionExistente.Municipio == municipio {
					return nil, errors.New("ubicación duplicada para el municipio " + municipio + " en la provincia " + provincia)
				}
			}
			mapUbicaciones[provincia] = append(ubicacionesProvincia, u)
		} else {
			mapUbicaciones[provincia] = []Ubicacion{u}
		}
	}

	return &GestorUbicaciones{ubicaciones: mapUbicaciones}, nil
}
