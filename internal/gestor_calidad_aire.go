package internal

import "errors"

type GestorCalidadAire struct {
	provincia                 string
	datosCalidadAireProvincia map[string][]DatosCalidadAire
	gestorUbicaciones         GestorUbicaciones 
}

func NewGestorCalidadAire(provincia string,datosCalidadAire []DatosCalidadAire,gestorUbicaciones GestorUbicaciones, )(GestorCalidadAire, error) {
	if provincia == "" {
		return GestorCalidadAire{}, errors.New("el nombre de la provincia es obligatorio")
	}

	if _, exists := gestorUbicaciones.ubicaciones[provincia]; !exists {
		return GestorCalidadAire{}, errors.New("la provincia no existe en el gestor de ubicaciones")
	}

	for _, datos := range datosCalidadAire {
		if datos.Ubicacion.Provincia != provincia {
			return GestorCalidadAire{}, errors.New("la provincia de los datos no coincide con la provincia del gestor")
		}
	}

	gestor := GestorCalidadAire{
		provincia:                 provincia,
		datosCalidadAireProvincia: make(map[string][]DatosCalidadAire),
		gestorUbicaciones:         gestorUbicaciones,
	}

	gestor.datosCalidadAireProvincia[provincia] = append(gestor.datosCalidadAireProvincia[provincia], datosCalidadAire...)

	return gestor, nil
}
