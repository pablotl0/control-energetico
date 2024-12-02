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

var (
	fechaActual = time.Now().UTC()
    fechaMinima = Fecha{
		Año: fechaActual.AddDate(0, 0, -1).Year(),  
		Mes: int(fechaActual.AddDate(0, 0, -1).Month()),
		Día: fechaActual.AddDate(0, 0, -1).Day(),
	}

	fechaMaxima = Fecha{
		Año: fechaActual.Year(),
		Mes: int(fechaActual.Month()),
		Día: fechaActual.Day(),
	}
)

func NewMuestraCalidadAire(fechaInicial Fecha, contaminantes []Contaminante) (*MuestraCalidadAire, error) {
    if fechaInicial.Año < fechaMinima.Año || (fechaInicial.Año == fechaMinima.Año && fechaInicial.Mes < fechaMinima.Mes) ||
        (fechaInicial.Año == fechaMinima.Año && fechaInicial.Mes == fechaMinima.Mes && fechaInicial.Día < fechaMinima.Día) {
        return nil, errors.New("la fecha inicial debe estar después del 1 de enero de 2000")
    }

    if fechaInicial.Año > fechaMaxima.Año || (fechaInicial.Año == fechaMaxima.Año && fechaInicial.Mes > fechaMaxima.Mes) ||
        (fechaInicial.Año == fechaMaxima.Año && fechaInicial.Mes == fechaMaxima.Mes && fechaInicial.Día > fechaMaxima.Día) {
        return nil, errors.New("la fecha inicial no puede estar en el futuro")
    }

    if len(contaminantes) == 0 {
        return nil, errors.New("contaminantes no puede estar vacío")
    }

    contaminantesMap := make(map[string]Contaminante)
    for _, contaminante := range contaminantes {
        if contaminante.MagnitudUnidad.Magnitud == "" {
            return nil, errors.New("cada contaminante debe tener una magnitud")
        }
        if _, exists := contaminantesMap[contaminante.MagnitudUnidad.Magnitud]; exists {
            return nil, errors.New("no se pueden duplicar magnitudes en los contaminantes")
        }
        contaminantesMap[contaminante.MagnitudUnidad.Magnitud] = contaminante
    }

    return &MuestraCalidadAire{
        FechaInicial:  fechaInicial,
        Contaminantes: contaminantesMap,
    }, nil
}
