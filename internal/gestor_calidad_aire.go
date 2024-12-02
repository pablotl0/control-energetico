package internal

import "errors"

type GestorCalidadAire struct {
	datosCalidadAireProvincia map[string][]MuestreoCalidadAire
	gestorUbicaciones         GestorUbicaciones
}

func NewGestorCalidadAire(datosCalidadAire map[string][]MuestreoCalidadAire, gestorUbicaciones GestorUbicaciones) (*GestorCalidadAire, error) {
	for provincia := range datosCalidadAire {
		if _, exists := gestorUbicaciones.ubicaciones[provincia]; !exists {
			return nil, errors.New("la provincia " + provincia + " no existe en el gestor de ubicaciones")
		}
	}

	gestor := &GestorCalidadAire{
		datosCalidadAireProvincia: datosCalidadAire,
		gestorUbicaciones:         gestorUbicaciones,
	}

	return gestor, nil
}

