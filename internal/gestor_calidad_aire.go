package internal

import "errors"

type GestorCalidadAire struct {
	datosCalidadAireProvincia map[string][]CalidadAireUbicacion
	ubicaciones               map[string][]Ubicacion
}

func NewGestorCalidadAire(datosCalidadAire map[string][]CalidadAireUbicacion, ubicaciones []Ubicacion) (*GestorCalidadAire, error) {
	if len(ubicaciones) == 0 {
        return nil, errors.New("debe haber al menos una ubicación")
    }

    ubicacionesMap := make(map[string][]Ubicacion)
    for _, u := range ubicaciones {
        ubicacionesMap[u.Provincia] = append(ubicacionesMap[u.Provincia], u)
    }

    for provincia, muestreos := range datosCalidadAire {
        ubicacionesProvincia, existe := ubicacionesMap[provincia]
        if !existe {
            return nil, errors.New("la provincia " + provincia + " no existe en las ubicaciones")
        }

        for _, muestreo := range muestreos {
            if !ubicacionValida(muestreo.Municipio, ubicacionesProvincia) {
                return nil, errors.New("el municipio " + muestreo.Municipio + " no existe en la provincia " + provincia)
            }
        }
    }

    return &GestorCalidadAire{
        datosCalidadAireProvincia: datosCalidadAire,
        ubicaciones:               ubicacionesMap,
    }, nil
}
