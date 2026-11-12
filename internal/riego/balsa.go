package riego

import "fmt"

type ErrPorcentajeReserva struct{
	Porcentaje float64
}

func (e ErrorPorcentajeReserva) Error() string {
	return fmt.Sprintf("el porcentaje actual de la reserva (%.2f) no puede ser negativo ni superar el 100%%", e.Porcentaje)
}

type Balsa struct {
	porcenajeReserva float64
	plan              PlanDistribucion
	historicoClima    []RegistroClimatico
}