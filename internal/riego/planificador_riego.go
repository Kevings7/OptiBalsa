package riego

import (
	"fmt"
	"time"
)

const (
	DiasEntreRecargas = 14
	AniosHistorico    = 5
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

type ErrorSinHistorico struct {
	Fecha time.Time
}

func (e ErrorSinHistorico) Error() string {
	return fmt.Sprintf("no hay registros de los últimos %d años para el %s", AniosHistorico, e.Fecha.Format(formatoFecha))
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

func (p *PlanificadorRiego) necesidadEsperada(dia time.Time) (float64, error) {
	var suma float64
	var anios int
	for atras := 1; atras <= AniosHistorico; atras++ {
		anterior := mismoDiaAnios(dia, atras)
		registro, existe := p.historicoClima[anterior]
		if existe && anterior.Month() == dia.Month() {
			suma += registro.Necesidad()
			anios++
		}
	}
	if anios == 0 {
		return 0, ErrorSinHistorico{Fecha: dia}
	}
	return suma / float64(anios), nil
}

func soloDia(fecha time.Time) time.Time {
	return time.Date(fecha.Year(), fecha.Month(), fecha.Day(), 0, 0, 0, 0, time.UTC)
}

func mismoDiaAnios(dia time.Time, atras int) time.Time {
	return time.Date(dia.Year()-atras, dia.Month(), dia.Day(), 0, 0, 0, 0, time.UTC)
}
