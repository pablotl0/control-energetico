package internal

import (
	"errors"
	"fmt"
)

type MuestreoCalidadAire struct {
	Ubicacion    Ubicacion
	DatosMuestreo map[Fecha]MuestraCalidadAire
}

func NewMuestreoCalidadAire(ubicacion Ubicacion, datosMuestreo []MuestraCalidadAire) (*MuestreoCalidadAire, error) {
	if ubicacion.Provincia == "" || ubicacion.Municipio =="" {
		return nil, errors.New("la provinciay municipio es obligatorio")
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
		Ubicacion:    ubicacion,
		DatosMuestreo: datosMap,
	}, nil
}
