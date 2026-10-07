# [HU001] Planificar el uso de la reserva según el histórico climatológico

[Viaje de Usuario](/docs/user-journeys.md#user-journey-1-juan-gonzalez)  
[Datos](/README.md#datos-y-conocimiento-del-dominio)

Juan dispone aproximadamente de catorce días entre dos recargas de su balsa. Durante ese intervalo debe decidir en qué momentos realizar el riego y qué porcentaje de la reserva utilizar.

Para estudiar el problema se dispone de registros climatológicos históricos de la zona, entre ellos temperatura, humedad relativa, viento, radiación solar, precipitación y evapotranspiración de referencia (ETo), obtenidos de Baza (estación próxima a Bácor Olivar).

Las condiciones registradas no son iguales durante todo el periodo. Un intervalo con varios días secos y elevada demanda evaporativa presenta unas necesidades distintas de otro con precipitaciones o menor evapotranspiración.

El problema consiste en establecer, utilizando los datos históricos y las reglas de riego conocidas por Juan, cómo distribuir porcentualmente el uso de la reserva durante el intervalo entre dos recargas.

Una distribución inadecuada puede consumir una parte excesiva de la reserva antes de llegar a los días de mayor necesidad o, por el contrario, restringir el riego cuando las condiciones históricas indican una mayor necesidad de agua.


# [HU002] Diferenciar el riego entre fincas con condiciones climatológicas distintas

[Viaje de Usuario](/docs/user-journeys.md#user-journey-2-diferencias-entre-fincas)  
[Datos](/README.md#datos-y-conocimiento-del-dominio)

Juan gestiona dos fincas situadas en Bácor-Olivar y Pozo Alcón. Aunque el sistema de riego utilizado es similar, las condiciones climatológicas históricas de ambas ubicaciones no son necesariamente iguales.

Para cada zona se dispone de registros históricos de temperatura, humedad relativa, viento, radiación solar, precipitación y evapotranspiración de referencia (ETo), procedentes de Baza y Pozo Alcón.

El problema aparece cuando se aplica un mismo criterio de riego a las dos fincas pese a que sus condiciones climatológicas históricas presentan diferencias.

El problema consiste en determinar cómo esas diferencias modifican la distribución porcentual del riego, evitando aplicar el mismo criterio a fincas sometidas a condiciones climatológicas distintas.

Dos periodos equivalentes en ambas fincas no tienen por qué producir la misma distribución de riego si las condiciones climatológicas registradas son diferentes.