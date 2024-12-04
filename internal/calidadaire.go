package internal

import (
	"errors"
	"fmt"
)



type RegistroCalidadAire struct {
	Provincia     string
	Municipio 	  string
	DatosMuestreo map[Fecha]MuestraCalidadAire 
}

func NewRegistroCalidadAire(provincia string,municipio string datosMuestreo []MuestraCalidadAire) (*RegistroCalidadAire, error) {
	if provincia == "" || municipio == ""{
		return nil, errors.New("la provincia y el municipio es obligatorio")
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

	return &RegistroCalidadAire{
		Provincia:     provincia,
		Municipio: 	   municipio,
		DatosMuestreo: datosMap,
	}, nil
}
