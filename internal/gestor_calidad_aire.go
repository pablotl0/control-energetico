package internal

import (
	"errors"
)

type GestorCalidadAire struct {
	provincia         string                    
	datosCalidadAire  map[string][]DatosCalidadAire   
	
}

func NewGestorCalidadAire(provincia string, datosCalidadAire []DatosCalidadAire) (GestorCalidadAire, error) {
	if provincia == "" {
		return GestorCalidadAire{}, errors.New("el nombre de la provincia es obligatorio")
	}

	for _, datos := range datosCalidadAire {
		if datos.Ubicacion.Provincia != provincia {
			return GestorCalidadAire{}, errors.New("la provincia de los datos no coincide con la provincia del gestor")
		}
	}

	gestor := GestorCalidadAire{
		provincia:        provincia,
		datosCalidadAire: make(map[string][]DatosCalidadAire),
	}

	gestor.datosCalidadAire[provincia] = append(gestor.datosCalidadAire[provincia], datosCalidadAire...)

	return gestor, nil
}
