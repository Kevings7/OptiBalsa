package riego

import (
	"errors"
	"time"
)

var (
	ErrHumedadFueraDeRango      = errors.New("La humedad relativa debe estar entre 0 y 100")
	ErrValorClimaticoNegativo   = errors.New("Las variables de viento, radiacion, precipitacion y ETo no pueden ser negativas")
)
type RegistroClimatico struct {
	fecha         time.Time
	temperatura   float64
	humedad       float64
	viento        float64
	radiacion     float64
	precipitacion float64
	eto           float64
}