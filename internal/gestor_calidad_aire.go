package internal

import (
	"errors"
	"time"
)

type GestorCalidadAire struct {
	provincia         string                    
	datosCalidadAire  map[string][]DatosCalidadAire   
	
}

func NewGestorCalidadAire(provincia string) (GestorCalidadAire, error) {
	if provincia == "" {
		return GestorCalidadAire{}, errors.New("el nombre de la provincia es obligatorio")
	}
	return GestorCalidadAire{
		provincia:        provincia,
		datosCalidadAire: make(map[string][]DatosCalidadAire), 
	}, nil
}
