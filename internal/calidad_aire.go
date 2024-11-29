package internal

import (
    "errors"
    "time"
)

type DatosCalidadAire struct {
    Ubicacion     Ubicacion
    DatosMuestreo map[time.Time]MuestraCalidadAire
}

func NewDatosCalidadAire(ubicacion Ubicacion, datosMuestreo []MuestraCalidadAire) (DatosCalidadAire, error) {
    if len(datosMuestreo) == 0 {
        return DatosCalidadAire{}, errors.New("debe haber al menos un dato de muestreo")
    }

    datosMap := make(map[time.Time]MuestraCalidadAire)
	for _, muestra := range datosMuestreo {
		if _, exists := datosMap[muestra.FechaInicial]; exists {
			return DatosCalidadAire{}, errors.New("datos duplicados para la misma fecha")
		}
		datosMap[muestra.FechaInicial] = muestra
	}

    return DatosCalidadAire{Ubicacion: ubicacion, DatosMuestreo: datosMap}, nil
}
