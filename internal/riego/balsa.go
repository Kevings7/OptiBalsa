package riego

import "errors"

var ErrPorcentajeReservaFallo  = errors.New("El porcentaje actual de la reserva no puede ser negativo ni superar el 100%")

type Balsa struct {
	porcenajeReserva float64
	plan              PlanDistribucion
	historicoClima    []RegistroClimatico
}