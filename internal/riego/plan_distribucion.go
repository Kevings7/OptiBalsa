package riego

const DiasEntreRecargas = 14

type PlanDistribucion struct {
	balsa        Balsa
	asignaciones [DiasEntreRecargas]AsignacionDiaria
}