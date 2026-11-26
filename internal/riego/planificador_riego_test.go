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
