package internal

import (
	"errors"
	"fmt"
)

type Fecha struct {
    Año int
    Mes int
    Día int
}

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

	datosMap := make(map[int]map[int]map[int]MuestraCalidadAire)

	for _, muestra := range datosMuestreo {
		year, month, day := muestra.FechaInicial.Date()
		
		if _, ok := datosMap[year]; !ok {
			datosMap[year] = make(map[int]map[int]MuestraCalidadAire)
		}
		if _, ok := datosMap[year][int(month)]; !ok {
			datosMap[year][int(month)] = make(map[int]MuestraCalidadAire)
		}
		if _, exists := datosMap[year][int(month)][day]; exists {
			return nil, fmt.Errorf("datos duplicados para la fecha: %04d-%02d-%02d", year, month, day)
		}

		datosMap[year][int(month)][day] = muestra
	}

	return &MuestreoCalidadAire{
		Provincia:     provincia,
		DatosMuestreo: datosMap,
	}, nil
}
