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

# Datos y aproximación del problema 

El uso de datos meteorológicos será de especial utilidad respecto a determinar la climatología, influyente en la determinación de la frecuencia y tiempo de riego. Para eso es posible  hacer uso de una web climatológica que revele datos relevantes como humedad e información de precipitaciones, tal como la Red de Información Agroclimática de Andalucía (RIA), con el fin de extraer los datos para su análisis.

El mecanismo del goteo de la balsa, canales y presión queda completamente delegado a la construcción de la misma, en definitiva, las héctareas de terreno, tipo exacto de cultivo o litro por segundo no son parámetros relevantes. 

Su uso cotidiano se basa principalmente en la gestión temporal de su capacidad, objetivo de este problema.

Analizar la información disponible del entorno, especialmente la climatológica, resultará vital para hacer una inferencia acerca del uso de la balsa. 

Para realizar la metodología DDD se deberán disponer de datos de expertos. 
Tras consulta, me han derivado a fuentes de información como las siguientes proporcionadas: 

https://www.rainbird.com/es/agencia/consejos-de-diseno-de-riego-necesidades-climaticas-y-de-riego

https://redivia.gva.es/bitstream/handle/20.500.11939/5864/2017_Esteban_Agrometereolog%C3%ADa.pdf?sequence=1&isAllowed=y


La obtención de datos climatológicos estructurados se pueden obtener de las siguientes fuentes: 

https://www.juntadeandalucia.es/agriculturaypesca/ifapa/riaweb/web/datosabiertos (Datos históricos)

https://open-meteo.com/ (Predicciones)

https://www.aemet.es/es/datos_abiertos/catalogo (Predicciones)

La decisión e implementación de los mismos para la heurística(*) del mismo serán decisión del programador a partir del M1. 

(*) Heurística se plantea para aclarar la lógica de negocio con fines educativos. La implementación de la misma dependería de la decisión del programador cuando plantee el problema. 

# Abordando el proyecto

Se han añadido las primeras Historias de Usuario al proyecto junto a los primeros Milestones.

[Ver información relativa.](#enlaces-de-interés-al-desarrollo)

# Decisiones considerables.

Para la modelización se usará una metodología DDD, por su abordamiento del dominio del problema ante uno tan difuso. Posteriormente se refinará con reglas mediante Example Mapping, algo que cobra especial interés con las historias de usuario sobre las que ya hemos trabajado, pero para ello debemos haber abordado el problema de una manera general. 

La resolución de los beneficios de la HU001 resulta previa pues la misma nos permite comenzar a formular la lógica de negocio para el M1. 

# Enlaces de interés al desarrollo.

A continuación se añaden enlaces a las HUs para comprender el desarrollo. 

[Historias de Usuario](docs/historias-de-usuario.md)

Enlace a la UserJourney

[UserJourney](docs/user-journeys.md)

Enlace a los Milestones

[Milestone](docs/milestones.md)


# Enlace a imágenes relativas a la docencia. 


[Cliente](/media/Cliente.jpeg)
[Desarrollador](/media/Desarrollador.jpeg)
[Configuración del repositorio](docs/configuracion.md)












