package riego

import (
	"errors"
	"testing"
	"time"
)

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
