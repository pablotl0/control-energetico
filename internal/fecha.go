package internal

import (
    "errors"
    "fmt"
    "time"
)

type Fecha struct {
    Año uint
    Mes time.Month
    Día uint
}

func NewFecha(año uint, mes time.Month, día uint) (*Fecha, error) {
    currentYear := uint(time.Now().Year())
    if año < 2000 || año > currentYear {
        return nil, fmt.Errorf("el año debe estar entre 2000 y el año actual (%d)", currentYear)
    }

    if día < 1 || día > 31 {
        return nil, errors.New("el día debe estar entre 1 y 31")
    }

    dateStr := fmt.Sprintf("%04d-%02d-%02d", año, mes, día)
    if _, err := time.Parse("2006-01-02", dateStr); err != nil {
        return nil, fmt.Errorf("la fecha es inválida: %v", err)
    }

    return &Fecha{
        Año: año,
        Mes: mes,
        Día: día,
    }, nil
}
