package riego

type Balsa struct {
	porcenajeReserva float64
	plan              PlanDistribucion
	historicoClima    []RegistroClimatico
}