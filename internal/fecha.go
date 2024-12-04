package internal

import (
    "errors"
    "time"
)

type Fecha struct {
    Año int       
    Mes time.Month 
    Día int        
}

func NewFecha(año int, mes time.Month, día int) (*Fecha, error) {
    if año < 2000 || año > time.Now().Year() {
        return nil, errors.New("el año debe estar entre 2000 y el actual")
    }

    if mes < 1 || mes > 12 {
        return nil, errors.New("el mes debe estar entre 1 y 12")
    }

    if día < 1 || día > 31 { 
        return nil, errors.New("el día debe estar entre 1 y 31")
    }

    if _, err := time.Parse("2006-1-2", fmt.Sprintf("%04d-%d-%02d", año, mes, día)); err != nil {
        return nil, errors.New("la fecha es inválida: " + err.Error())
    }

    return &Fecha{
        Año: año,
        Mes: mes,
        Día: día,
    }, nil
}