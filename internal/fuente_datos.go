package internal

import (
	"time"
)

type FuenteDatos struct {
	Nombre              string
	Tipo                string // Tipo de fuente (estaciones de medición, informes gubernamentales y bases de datos de investigación)
	URL                 string
	Formato             string // ( "CSV", "PDF", "Excel")
	UltimaActualizacion time.Time
}

// Constructor para FuenteDatos
func NewFuenteDatos(nombre, tipo, url, formato string, ultimaActualizacion time.Time) (FuenteDatos, error) {
	if nombre == "" || tipo == "" || url == "" || formato == "" {
		return FuenteDatos{}, errors.New("nombre, tipo, url y formato son obligatorios")
	}
	if formato != "CSV" && formato != "PDF" && formato != "Excel" {
		return FuenteDatos{}, errors.New("formato debe ser 'CSV', 'PDF' o 'Excel'")
	}
	return FuenteDatos{Nombre: nombre, Tipo: tipo, URL: url, Formato: formato, UltimaActualizacion: ultimaActualizacion}, nil
}
