package riego

import (
	"errors"
	"testing"
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
