package internal

import "errors"

type GestorCalidadAire struct {
	datosCalidadAireUbicacion map[Ubicacion][]MuestreoCalidadAire
}

func NewGestorCalidadAire(datosCalidadAire map[Ubicacion][]MuestreoCalidadAire) (*GestorCalidadAire, error) {
	for ubicacion := range datosCalidadAire {
		if ubicacion.Provincia == "" || ubicacion.Municipio == ""|| ubicacion.Latitud == 0 || ubicacion.Longitud == 0 {
			return nil, errors.New("ubicación inválida, la provincia,municipio, latitud y longitud son obligatorios")
		}
	}

	gestor := &GestorCalidadAire{
		datosCalidadAireUbicacion: datosCalidadAire,
	}

	return gestor, nil
}
