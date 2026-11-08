package riego

import (
	"fmt"
	"time"
)

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
	fecha         time.Time
	temperatura   float64
	humedad       float64
	viento        float64
	radiacion     float64
	precipitacion float64
	eto           float64
}