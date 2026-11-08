package riego

import "fmt"

const DiasEntreRecargas = 14

type ErrorPorcentajeRiego struct {
	Porcentaje float64
}

func (e ErrorPorcentajeRiego) Error() string {
	return fmt.Sprintf("el porcentaje de uso asignado (%.2f) debe estar entre 0 y 100", e.Porcentaje)
}

type porcentajeRiego float64

type PlanDistribucion struct {
	asignaciones [DiasEntreRecargas]porcentajeRiego
}