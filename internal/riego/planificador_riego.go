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

type ErrorPorcentajeRiego struct {
	Porcentaje float64
}

func (e ErrorPorcentajeRiego) Error() string {
	return fmt.Sprintf("el porcentaje de uso asignado (%.2f) debe estar entre 0 y 100", e.Porcentaje)
}

type PorcentajeRiego float64

type PlanificadorRiego struct {
	porcentajeReserva float64
	historicoClima    map[time.Time]RegistroClimatico
	planDistribucion  [DiasEntreRecargas]PorcentajeRiego
}
