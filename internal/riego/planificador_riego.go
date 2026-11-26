package riego

import (
	"fmt"
	"time"
)

const DiasEntreRecargas = 14

type ErrorPorcentajeReserva struct {
	Porcentaje float64
}

func (e ErrorPorcentajeReserva) Error() string {
	return fmt.Sprintf("el porcentaje actual de la reserva (%.2f) debe estar entre 0 y 100", e.Porcentaje)
}

type PorcentajeRiego float64

type PlanificadorRiego struct {
	porcentajeReserva float64
	historicoClima    map[time.Time]RegistroClimatico
	planDistribucion  [DiasEntreRecargas]PorcentajeRiego
}

func NuevoPlanificadorRiego(porcentajeReserva float64) (*PlanificadorRiego, error) {
	if porcentajeReserva < 0 || porcentajeReserva > 100 {
		return nil, ErrorPorcentajeReserva{Porcentaje: porcentajeReserva}
	}
	return &PlanificadorRiego{
		porcentajeReserva: porcentajeReserva,
		historicoClima:    map[time.Time]RegistroClimatico{},
	}, nil
}
