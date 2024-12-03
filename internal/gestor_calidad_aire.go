package internal

import "errors"

type GestorCalidadAire struct {
	datosCalidadAireProvincia map[string][]CalidadAireUbicacion
	gestorUbicaciones         GestorUbicaciones
}

func NewGestorCalidadAire(datosCalidadAire map[string][]CalidadAireUbicacion, gestorUbicaciones GestorUbicaciones) (*GestorCalidadAire, error) {
	for provincia, muestreos := range datosCalidadAire {
		ubicaciones, exists := gestorUbicaciones.ubicaciones[provincia]
		if !exists {
			return nil, errors.New("la provincia " + provincia + " no existe en el gestor de ubicaciones")
		}

		for _, muestreo := range muestreos {
			encontrado := false
			for _, ubicacion := range ubicaciones {
				if ubicacion.Municipio == muestreo.Municipio {
					encontrado = true
					break
				}
			}
			if !encontrado {
				return nil, errors.New("el municipio " + muestreo.Municipio + " no existe en la provincia " + provincia + " del gestor de ubicaciones")
			}
		}
	}

	gestor := &GestorCalidadAire{
		datosCalidadAireProvincia: datosCalidadAire,
		gestorUbicaciones:         gestorUbicaciones,
	}

	return gestor, nil
}
