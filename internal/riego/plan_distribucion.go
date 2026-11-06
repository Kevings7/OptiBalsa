package riego

const DiasEntreRecargas = 14

type porcentajeRiego float64

type PlanDistribucion struct {
	asignaciones [DiasEntreRecargas]porcentajeRiego
}