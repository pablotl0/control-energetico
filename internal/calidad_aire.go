package internal

import (
    "errors"
)

type DatosCalidadAire struct {
    Ubicacion     Ubicacion
    DatosMuestreo []MuestraCalidadAire 
}

func NewDatosCalidadAire(ubicacion Ubicacion, datosMuestreo []MuestraCalidadAire) (DatosCalidadAire, error) {
    if len(datosMuestreo) == 0 {
        return DatosCalidadAire{}, errors.New("debe haber al menos un dato de muestreo")
    }

    return DatosCalidadAire{Ubicacion: ubicacion, DatosMuestreo: datosMuestreo}, nil
}
