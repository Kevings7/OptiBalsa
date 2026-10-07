# OptiBalsa
Sistema para optimizar regadío con balsa.

# Origen

Conocer el momento idóneo para aprovechar el riego no resulta fácil. Desde la moda de construcción de balsas y la desaparición paulatina del riego a manta, la importancia de saber gestionar el riego de una finca es una cuestión de especial interés para adaptarse a los tiempos actuales y venideros. 

Perder el cultivo por ahogo o sequedad es un problema real si no se gestiona adecuadamente el riego del mismo. 

El agua es un recurso finito junto con el dinero puesto en el mantenimiento de los cultivos. Por ello conocer el momento idóneo para aprovechar al máximo este primer recurso ayudará al no detrimento del segundo. 

# Perspectiva personal hacia un problema común. 

Actualmente gestiono una finca en pleno Geoparque de Granada. El clima árido siempre ha proporcionado de incertidumbre a todos los agricultores de la zona. El tiempo ha causado el deterioro fluvial de la zona, manifestado en un escaso rendimiento de los pozos. Siempre que entablo conversación con mis vecinos aparece la misma duda fruto de la preocupación de la pérdida del cultivo, "¿Lo estoy haciendo bien?". Actualmente la comunidad de regantes establece un tiempo de relleno de la balsa de dos semanas aproximadamente. 

Saber cuándo regar es vital para evitar perder el cultivo y mantener su máxima productividad, por eso es vital la correcta gestión de la balsa. 

# Objetivo

Con este sistema se busca la optimalidad de las horas de riego en balsa en virtud de aprovechar al máximo su capacidad. Con la información proporcionada el agricultor conocerá la gestión de su riego para usarlo de manera provechosa.  

# Breve proximación del problema 

El uso de datos meteorológicos será de especial utilidad respecto a determinar la climatología, influyente en la determinación de la frecuencia y tiempo de riego. Para eso es posible  hacer uso de una web climatológica que revele datos relevantes como humedad e información de precipitaciones, tal como la Red de Información Agroclimática de Andalucía (RIA), con el fin de extraer los datos para su análisis.

El mecanismo del goteo de la balsa, canales y presión queda completamente delegado a la construcción de la misma, en definitiva, las héctareas de terreno, tipo exacto de cultivo o litro por segundo no son parámetros relevantes. 

Su uso cotidiano se basa principalmente en la gestión temporal de su capacidad, objetivo de este problema.

Analizar la información disponible del entorno, especialmente la climatológica, resultará vital para hacer una inferencia acerca del uso de la balsa. 

Para realizar la metodología DDD se deberán disponer de datos de expertos especificados en el siguiente apartado.

# Datos y conocimiento del dominio

Para abordar el problema se tendrán en cuenta tanto datos propios de la explotación como información meteorológica externa  y conocimiento relativo al riego, que se mencionarán a continuación.

Los datos a extraer se encuentran estructurados en formato CSV.

# Datos meteorológicos

Los principales factores climatológicos considerados son:

- **Temperatura:** temperaturas elevadas tienden a aumentar la demanda hídrica del cultivo.
- **Humedad relativa:** una humedad baja favorece una mayor pérdida de agua, mientras que valores altos reducen esa demanda.
- **Viento:** puede incrementar la pérdida de agua por evapotranspiración.
- **Radiación solar:** una mayor radiación aumenta la energía disponible para la evapotranspiración.
- **Precipitación:** disminuye la necesidad de aportar agua mediante riego.
- **Evapotranspiración de referencia (ETo):** permite representar de forma conjunta el efecto de varios factores meteorológicos sobre la demanda de agua.

Estos datos se obtendrán de una fuente meteorológica estructurada (datos en CSV), indicando para cada dato su procedencia y formato. La Red de Información Agroclimática de Andalucía (RIA)  proporciona, entre otros, datos de temperatura, humedad, viento, radiación, precipitación y evapotranspiración de referencia. 

https://www.juntadeandalucia.es/agriculturaypesca/ifapa/riaweb/web/

https://www.juntadeandalucia.es/agriculturaypesca/ifapa/riaweb/web/estacion/23/2

https://www.juntadeandalucia.es/agriculturaypesca/ifapa/riaweb/web/estacion/18/1


# Factores propios de la explotación

Además de las condiciones meteorológicas, la decisión de riego depende de factores propios del problema:

- **Porcentaje de agua disponible en la balsa:** limita el agua que puede utilizarse.
- **Tiempo restante hasta la siguiente recarga:** condiciona cuánto agua puede emplearse sin comprometer el cultivo. 

Datos obtenidos en el propio problema. 

# Reglas prácticas utilizadas en el riego

A partir de la experiencia de Juan se tienen en cuenta las siguientes reglas:

- Cuando ha habido precipitaciones recientes, reduce el porcentaje de reserva destinado al riego.
- Evita agotar una parte excesiva de la reserva en un único riego cuando todavía quedan varios días hasta la siguiente recarga.


# Abordando el proyecto

Se han añadido las primeras Historias de Usuario al proyecto junto a los primeros Milestones.

[Ver información relativa.](#enlaces-de-interés-al-desarrollo)

# Decisiones considerables.

Para la modelización se usará una metodología DDD, por su abordamiento del dominio del problema. Para ello se pueden consultar la información de expertos de cara al desarrollo en [información relativa](#datos-y-conocimiento-del-dominio). 

# Enlaces de interés al desarrollo.

A continuación se añaden enlaces a las HUs para comprender el desarrollo. 

[Historias de Usuario](docs/historias-de-usuario.md)

Enlace a la UserJourney

[UserJourney](docs/user-journeys.md)

Enlace a los Milestones

[Milestone](docs/milestones.md)

Enlace a descripción de las personas. 

[Personas](docs/personas.md)


# Enlace a imágenes relativas a la docencia. 


[Cliente](/media/Cliente.jpeg)
[Desarrollador](/media/Desarrollador.jpeg)
[Configuración del repositorio](docs/configuracion.md)












