package riego

import (
	"fmt"
	"time"
)

const (
	DiasEntreRecargas = 14
	formatoFecha      = "02/01/2006"
)

type ErrorPorcentajeReserva struct {
	Porcentaje float64
}

func (e ErrorPorcentajeReserva) Error() string {
	return fmt.Sprintf("el porcentaje actual de la reserva (%.2f) debe estar entre 0 y 100", e.Porcentaje)
}

type ErrorFechaRepetida struct {
	Fecha time.Time
}

func (e ErrorFechaRepetida) Error() string {
	return fmt.Sprintf("ya existe un registro climático para el %s", e.Fecha.Format(formatoFecha))
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

func (p *PlanificadorRiego) RegistrarClima(fecha time.Time, registro RegistroClimatico) error {
	dia := soloDia(fecha)
	if _, existe := p.historicoClima[dia]; existe {
		return ErrorFechaRepetida{Fecha: dia}
	}
	p.historicoClima[dia] = registro
	return nil
}

func soloDia(fecha time.Time) time.Time {
	return time.Date(fecha.Year(), fecha.Month(), fecha.Day(), 0, 0, 0, 0, time.UTC)
}
