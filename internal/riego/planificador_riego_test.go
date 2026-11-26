package riego

import (
	"errors"
	"math"
	"testing"
	"time"
)

var inicioPeriodo = time.Date(2026, time.August, 1, 0, 0, 0, 0, time.UTC)

func planificadorConNecesidades(t *testing.T, reserva float64, necesidades [DiasEntreRecargas]float64) *PlanificadorRiego {
	t.Helper()
	planificador, err := NuevoPlanificadorRiego(reserva)
	if err != nil {
		t.Fatal(err)
	}
	for atras := 1; atras <= AniosHistorico; atras++ {
		for i, necesidad := range necesidades {
			registro, _ := NuevoRegistroClimatico(30, 40, 2, 25, 0, necesidad)
			fecha := inicioPeriodo.AddDate(-atras, 0, i)
			if err := planificador.RegistrarClima(fecha, registro); err != nil {
				t.Fatal(err)
			}
		}
	}
	return planificador
}

func suma(plan [DiasEntreRecargas]PorcentajeRiego) float64 {
	var total float64
	for _, porcentaje := range plan {
		total += float64(porcentaje)
	}
	return total
}

func casiIgual(a, b float64) bool {
	return math.Abs(a-b) < 1e-9
}

func TestNuevoPlanificadorRechazaReservaFueraDeRango(t *testing.T) {
	for _, reserva := range []float64{-1, 100.1} {
		_, err := NuevoPlanificadorRiego(reserva)
		var errReserva ErrorPorcentajeReserva
		if !errors.As(err, &errReserva) {
			t.Errorf("reserva %.2f: se esperaba ErrorPorcentajeReserva, se obtuvo %v", reserva, err)
		}
	}
}

func TestNuevoPlanificadorAceptaReservaEnLosLimites(t *testing.T) {
	for _, reserva := range []float64{0, 100} {
		if _, err := NuevoPlanificadorRiego(reserva); err != nil {
			t.Errorf("reserva %.2f: no se esperaba error, se obtuvo %v", reserva, err)
		}
	}
}

func TestRegistrarClimaRechazaFechaRepetida(t *testing.T) {
	planificador, _ := NuevoPlanificadorRiego(80)
	registro, _ := NuevoRegistroClimatico(30, 40, 2, 25, 0, 6)
	fecha := time.Date(2015, time.December, 29, 0, 0, 0, 0, time.UTC)
	_ = planificador.RegistrarClima(fecha, registro)
	err := planificador.RegistrarClima(fecha.Add(12*time.Hour), registro)
	var errFecha ErrorFechaRepetida
	if !errors.As(err, &errFecha) {
		t.Errorf("se esperaba ErrorFechaRepetida, se obtuvo %v", err)
	}
}

func TestNecesidadEsperadaEsLaMediaDeLaNecesidadDeCadaAnio(t *testing.T) {
	planificador, _ := NuevoPlanificadorRiego(80)
	seco, _ := NuevoRegistroClimatico(30, 40, 2, 25, 0, 6)
	lluvioso, _ := NuevoRegistroClimatico(20, 90, 2, 10, 30, 6)
	antiguo, _ := NuevoRegistroClimatico(30, 40, 2, 25, 0, 100)
	dia := time.Date(2026, time.August, 3, 0, 0, 0, 0, time.UTC)
	_ = planificador.RegistrarClima(dia.AddDate(-1, 0, 0), seco)
	_ = planificador.RegistrarClima(dia.AddDate(-2, 0, 0), lluvioso)
	_ = planificador.RegistrarClima(dia.AddDate(-AniosHistorico-1, 0, 0), antiguo)
	necesidad, err := planificador.necesidadEsperada(dia)
	if err != nil || necesidad != 3 {
		t.Errorf("se esperaba una necesidad de 3.00 sin error, se obtuvo %.2f y %v", necesidad, err)
	}
}

func TestNecesidadEsperadaFallaSinHistoricoDelDia(t *testing.T) {
	planificador, _ := NuevoPlanificadorRiego(80)
	dia := time.Date(2026, time.August, 3, 0, 0, 0, 0, time.UTC)
	_, err := planificador.necesidadEsperada(dia)
	var errHistorico ErrorSinHistorico
	if !errors.As(err, &errHistorico) {
		t.Errorf("se esperaba ErrorSinHistorico, se obtuvo %v", err)
	}
}

func TestPlanificarFallaSiFaltaElHistoricoDeUnDia(t *testing.T) {
	planificador, _ := NuevoPlanificadorRiego(80)
	_, err := planificador.Planificar(inicioPeriodo)
	var errHistorico ErrorSinHistorico
	if !errors.As(err, &errHistorico) || !errHistorico.Fecha.Equal(inicioPeriodo) {
		t.Errorf("se esperaba ErrorSinHistorico para el %v, se obtuvo %v", inicioPeriodo, err)
	}
}

func TestPlanificarRepartePorNecesidadSinSuperarLaReserva(t *testing.T) {
	necesidades := [DiasEntreRecargas]float64{8, 4, 0, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6}
	plan, err := planificadorConNecesidades(t, 80, necesidades).Planificar(inicioPeriodo)
	if err != nil {
		t.Fatal(err)
	}
	if !casiIgual(suma(plan), 80) {
		t.Errorf("el plan debería sumar 80, suma %.2f", suma(plan))
	}
	if plan[0] <= plan[1] {
		t.Errorf("el día con más necesidad debería recibir más: %.2f frente a %.2f", plan[0], plan[1])
	}
	if plan[2] != 0 {
		t.Errorf("un día sin necesidad debería recibir 0, recibe %.2f", plan[2])
	}
}

func TestPlanificarSinNecesidadNoAsignaRiego(t *testing.T) {
	plan, err := planificadorConNecesidades(t, 80, [DiasEntreRecargas]float64{}).Planificar(inicioPeriodo)
	if err != nil || suma(plan) != 0 {
		t.Errorf("se esperaba un plan vacío sin error, se obtuvo suma %.2f y %v", suma(plan), err)
	}
}

func TestPlanificarNingunDiaSuperaLaMitadDeLaReserva(t *testing.T) {
	necesidades := [DiasEntreRecargas]float64{20, 1, 1}
	plan, err := planificadorConNecesidades(t, 80, necesidades).Planificar(inicioPeriodo)
	if err != nil {
		t.Fatal(err)
	}
	if !casiIgual(float64(plan[0]), 40) || !casiIgual(float64(plan[1]), 20) || !casiIgual(suma(plan), 80) {
		t.Errorf("se esperaba 40 para el día 1, 20 para el día 2 y suma 80, se obtuvo %v", plan)
	}
}
