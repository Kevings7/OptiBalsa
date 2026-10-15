package riego

import "errors"

var (
	ErrHumedadFueraDeRango      = errors.New("La humedad relativa debe estar entre 0 y 100")
	ErrValorClimaticoNegativo   = errors.New("Las variables de viento, radiacion, precipitacion y ETo no pueden ser negativas")
	ErrPorcentajeReservaFallo  = errors.New("El porcentaje actual de la reserva no puede ser negativo ni superar el 100%")
	ErrPorcentajeRiegoInvalido = errors.New("El porcentaje de uso asignado debe estar entre 0 y 100")
)