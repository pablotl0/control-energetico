package internal

import (
    "errors"
    "time"
)

type DatosCalidadAire struct {
    Ubicacion     Ubicacion
    DatosMuestreo map[string]MuestraCalidadAire //YYYY-MM-DD
}

func NewDatosCalidadAire(ubicacion Ubicacion, datosMuestreo []MuestraCalidadAire) (DatosCalidadAire, error) {
    if len(datosMuestreo) == 0 {
        return DatosCalidadAire{}, errors.New("debe haber al menos un dato de muestreo")
    }

    datosMap := make(map[string]MuestraCalidadAire)
	for _, muestra := range datosMuestreo {
		claveFecha := fmt.Sprintf("%04d-%02d-%02d", muestra.FechaInicial.Year(), muestra.FechaInicial.Month(), muestra.FechaInicial.Day())
		
		if _, exists := datosMap[claveFecha]; exists {
			return DatosCalidadAire{}, errors.New("datos duplicados para la misma fecha")
		}
		
		datosMap[claveFecha] = muestra
	}

    return DatosCalidadAire{Ubicacion: ubicacion, DatosMuestreo: datosMap}, nil
}
