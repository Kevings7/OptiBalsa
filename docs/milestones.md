A continuación se plantean los Milestones requeridos para el desarrollo;

## Milestone 0: Modelización inicial.

Se realizará un modelizado usando DDD a raíz de la HU001 con el objetivo de tener un modelado inicial codificado.

Será válido si se ha seguido la metodología DDD de manera acorde a la misma, dando lugar a un modelo mínimo del problema a partir de la historia de usuario. 

Con ello tendremos un PMV , manifestado como código, donde tendremos una representación lógica de los elementos del dominio y las relaciones existentes entre ellos.


## Milestone 1: Lógica de negocio de la HU001

Se implementa el reparto de la reserva de la balsa entre los 14 días que hay entre dos recargas, a partir del histórico climatológico de los últimos 5 años.

Será viable si pasan los tests que comprueban que:

- No se pueden crear registros climáticos con valores imposibles ni un planificador con una reserva fuera de 0–100.
- La necesidad de agua de un día es su ETo menos su precipitación, y nunca es negativa.
- El histórico no admite dos registros para el mismo día.
- La necesidad esperada de cada día es la media de los últimos 5 años, y la falta de histórico de un día produce un error.
- La suma del plan nunca supera la reserva, un día con más necesidad recibe más porcentaje y un día sin necesidad recibe 0 %.
- Ningún día recibe más del 50 % de la reserva.

## Milestone 2: Replanificación y diferenciación entre fincas

Con medias de varios años, el reparto dentro de un periodo sale casi uniforme. En este hito se replanifica cada día con la reserva restante y la lluvia real del día anterior, y se aplica el reparto a las fincas de Bácor-Olivar y Pozo Alcón con el histórico de cada una (HU002).