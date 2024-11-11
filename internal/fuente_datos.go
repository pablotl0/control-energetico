package internal

import (
    "time"
)

type FuenteDatos struct {
    Nombre         string    
    Tipo           string    // Tipo de fuente (estaciones de medición, informes gubernamentales y bases de datos de investigación)
    URL            string    
    Formato        string    // ( "CSV", "PDF", "Excel")
    UltimaActualizacion time.Time 
}