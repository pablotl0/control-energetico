package internal

import (
	"errors"
	"fmt"
)



type MuestreoCalidadAire struct {
	Provincia     string
	DatosMuestreo map[Fecha]MuestraCalidadAire 
}

func NewMuestreoCalidadAire(provincia string, datosMuestreo []MuestraCalidadAire) (*MuestreoCalidadAire, error) {
	if provincia == "" {
		return nil, errors.New("la provincia es obligatoria")
	}
	if len(datosMuestreo) == 0 {
		return nil, errors.New("debe haber al menos un dato de muestreo")
	}

	datosMap := make(map[Fecha]MuestraCalidadAire)

    for _, muestra := range datosMuestreo {
		fecha := Fecha{
			Año: muestra.FechaInicial.Year(),
			Mes: int(muestra.FechaInicial.Month()), 
			Día: muestra.FechaInicial.Day(),
		}

		if _, exists := datosMap[fecha]; exists {
			return nil, fmt.Errorf("datos duplicados para la fecha: %04d-%02d-%02d", fecha.Año, fecha.Mes, fecha.Día)
		}

		datosMap[fecha] = muestra
	}

	return &MuestreoCalidadAire{
		Provincia:     provincia,
		DatosMuestreo: datosMap,
	}, nil
}
