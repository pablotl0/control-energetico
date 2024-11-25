package internal

import (
    "errors"
)

// Es la estructura de datos principal
type DatosCalidadAire struct {
    Ubicacion     Ubicacion
    DatosMuestreo []MuestraCalidadAire 
}

// Constructor para DatosCalidadAire
func NewDatosCalidadAire(ubicacion Ubicacion, datosMuestreo []MuestraCalidadAire) (DatosCalidadAire, error) {
    if len(datosMuestreo) == 0 {
        return DatosCalidadAire{}, errors.New("debe haber al menos un dato de muestreo")
    }

    return DatosCalidadAire{Ubicacion: ubicacion, DatosMuestreo: datosMuestreo}, nil
}
