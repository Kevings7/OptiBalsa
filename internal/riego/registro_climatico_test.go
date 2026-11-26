package riego

import (
	"errors"
	"testing"
)

func TestNuevoRegistroClimaticoRechazaHumedadFueraDeRango(t *testing.T) {
	for _, humedad := range []float64{-1, 101} {
		_, err := NuevoRegistroClimatico(30, humedad, 2, 25, 0, 6)
		var errHumedad ErrorHumedad
		if !errors.As(err, &errHumedad) {
			t.Errorf("humedad %.2f: se esperaba ErrorHumedad, se obtuvo %v", humedad, err)
		}
	}
}

func TestNuevoRegistroClimaticoRechazaValoresNegativos(t *testing.T) {
	casos := []struct {
		variable                       string
		viento, radiacion, precip, eto float64
	}{
		{"viento", -1, 25, 0, 6},
		{"radiación", 2, -1, 0, 6},
		{"precipitación", 2, 25, -1, 6},
		{"ETo", 2, 25, 0, -1},
	}
	for _, caso := range casos {
		_, err := NuevoRegistroClimatico(30, 40, caso.viento, caso.radiacion, caso.precip, caso.eto)
		var errValor ErrorValorClimatico
		if !errors.As(err, &errValor) || errValor.Variable != caso.variable {
			t.Errorf("%s negativa: se esperaba ErrorValorClimatico, se obtuvo %v", caso.variable, err)
		}
	}
}

func TestNuevoRegistroClimaticoAceptaValoresValidos(t *testing.T) {
	if _, err := NuevoRegistroClimatico(30, 40, 2, 25, 0, 6); err != nil {
		t.Errorf("no se esperaba error, se obtuvo %v", err)
	}
}
