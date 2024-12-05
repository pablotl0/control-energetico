package internal

import (
	"errors"
	"fmt"
)

type RegistroCalidadAire struct {
	Provincia     string
	Municipio     string
	DatosMuestreo map[Fecha][]MuestraContaminante
}

func NewRegistroCalidadAire(provincia string, municipio string, datosMuestreo []MuestraContaminante, fechas []Fecha) (*RegistroCalidadAire, error) {
	if provincia == "" || municipio == "" {
		return nil, errors.New("la provincia y el municipio son obligatorios")
	}
	if len(datosMuestreo) == 0 {
		return nil, errors.New("debe haber al menos un dato de muestreo")
	}

	if len(datosMuestreo) != len(fechas) {
		return nil, errors.New("el número de fechas no coincide con el número de muestras de contaminantes")
	}

	datosMap := make(map[Fecha][]MuestraContaminante)

	for i, fecha := range fechas {
		if _, exists := datosMap[fecha]; exists {
			return nil, fmt.Errorf("datos duplicados para la fecha: %v", fecha)
		}
		datosMap[fecha] = append(datosMap[fecha], datosMuestreo[i])
	}

	return &RegistroCalidadAire{
		Provincia:     provincia,
		Municipio:     municipio,
		DatosMuestreo: datosMap,
	}, nil
}
