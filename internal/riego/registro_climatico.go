package riego

import "fmt"

type ErrorHumedad struct {
	Humedad float64
}

func (e ErrorHumedad) Error() string {
	return fmt.Sprintf("la humedad relativa (%.2f) debe estar entre 0 y 100", e.Humedad)
}

type ErrorValorClimatico struct {
	Variable string
	Valor    float64
}

func (e ErrorValorClimatico) Error() string {
	return fmt.Sprintf("la variable %s no puede ser negativa (valor aportado: %.2f)", e.Variable, e.Valor)
}

type RegistroClimatico struct {
	temperatura   float64
	humedad       float64
	viento        float64
	radiacion     float64
	precipitacion float64
	eto           float64
}

func NuevoRegistroClimatico(temperatura, humedad, viento, radiacion, precipitacion, eto float64) (RegistroClimatico, error) {
	if humedad < 0 || humedad > 100 {
		return RegistroClimatico{}, ErrorHumedad{Humedad: humedad}
	}
	noNegativas := []struct {
		nombre string
		valor  float64
	}{
		{"viento", viento},
		{"radiación", radiacion},
		{"precipitación", precipitacion},
		{"ETo", eto},
	}
	for _, variable := range noNegativas {
		if variable.valor < 0 {
			return RegistroClimatico{}, ErrorValorClimatico{Variable: variable.nombre, Valor: variable.valor}
		}
	}
	return RegistroClimatico{temperatura, humedad, viento, radiacion, precipitacion, eto}, nil
}

func (r RegistroClimatico) Necesidad() float64 {
	return max(r.eto-r.precipitacion, 0)
}
