package internal

import (
	"errors"
	"time"
)

type GestorCalidadAire struct {
	id                string                
	datosCalidadAire  []DatosCalidadAire    
	
}

func NewGestorCalidadAire(id string) (GestorCalidadAire, error) {
	if id == "" {
		return GestorCalidadAire{}, errors.New("el gestor debe tener un identificador único")
	}
	return GestorCalidadAire{
		id:               id,
		datosCalidadAire: []DatosCalidadAire{},
	}, nil
}
