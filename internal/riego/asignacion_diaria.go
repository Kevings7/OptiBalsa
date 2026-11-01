package riego

import "errors"

var ErrPorcentajeRiegoInvalido = errors.New("El porcentaje de uso asignado debe estar entre 0 y 100")

type AsignacionDiaria struct {
	dia             int
	porcentajeRiego float64
}