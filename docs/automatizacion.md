# Automatización del proyecto

## Gestión de dependencias

OptiBalsa utiliza Go Modules mediante `go.mod`.

Go incorpora de forma nativa la gestión de módulos y dependencias, por lo que no es necesario añadir un gestor de dependencias externo al proyecto.

## Gestor de tareas

### Alternativas consideradas

Se estudian las siguientes alternativas:

- GNU Make
- Just
- Task

La elección se realizará después de comparar las alternativas mediante mediciones realizadas sobre el mismo entorno.

## Criterios de comparación

Las herramientas de automatización se utilizarán posteriormente en entornos automatizados, especialmente contenedores. Por ello, debemos tener en cuenta ciertas características de cara a su elección.

Los criterios utilizados son los siguientes:

### Coste de instalación

Se medirá el tiempo necesario para disponer de cada gestor de tareas en un entorno en el que no esté instalado previamente.

Además del tiempo, se comprobará el tamaño que ocupa la herramienta instalada. Esto permite estimar el coste adicional que introduciría en futuros entornos de CI o contenedores.

### Dependencias adicionales

Se comprobará qué software adicional necesita cada alternativa para poder utilizarse. Se priorizarán las soluciones que introduzcan menos componentes adicionales en la infraestructura del proyecto.

### Gestión de dependencias del proyecto

OptiBalsa utiliza Go Modules como mecanismo de gestión de dependencias. Se comprobará cómo se integra la descarga de dependencias mediante `go mod download` con cada gestor de tareas.

Se distinguirá entre el coste de instalar el propio gestor de tareas y el coste de descargar las dependencias del proyecto.

### Tiempo de ejecución de tareas reales

Las alternativas se compararán ejecutando operaciones equivalentes del proyecto.

Se medirán las siguientes tareas:

- descarga de dependencias;
- compilación de los fuentes;
- ejecución de los tests.

Para que la comparación sea válida, los tres gestores ejecutarán exactamente las mismas órdenes de Go.

### Ejecución en un entorno limpio

Cuando sea relevante se realizarán mediciones sin reutilizar las dependencias descargadas previamente. De esta forma se aproxima el comportamiento de una ejecución de CI sin caché.

También se podrá comparar con una segunda ejecución utilizando los recursos ya descargados, que representa un entorno con caché.

### Uso posterior en automatización

Se tendrá en cuenta el coste que supone disponer de la herramienta en futuros workflows y contenedores. Este criterio se evaluará mediante datos medibles, como los pasos necesarios para instalarla, el tiempo de instalación y el tamaño añadido, evitando utilizar como criterio únicamente la familiaridad con la herramienta.

### Ligereza

Para comparar la ligereza de las herramientas se ha medido el tamaño en disco de sus ejecutables utilizando `ls -lh`. Las mediciones se han realizado sobre las versiones seleccionadas para la comparación.

| Herramienta | Versión | Tamaño del ejecutable |
| --- | --- | ---: |
| GNU Make | 4.4.1 | 281 KiB |
| Just | 1.59.0 | 6.2 MiB |
| Task | 3.54.0 | 50 MiB |

![Medición del tamaño de los gestores de tareas](img/tamanioherramientas.png)

En el entorno medido, GNU Make presenta el ejecutable de menor tamaño. Just requiere aproximadamente 6.2 MiB, mientras que Task alcanza aproximadamente 50 MiB.

Esta diferencia es relevante de cara al uso posterior de la herramienta en contenedores y sistemas de integración continua, ya que una herramienta de mayor tamaño incrementa la cantidad de datos que deben incorporarse al entorno.


### Procedencia y mantenimiento de las herramientas

Para evitar medir herramientas obtenidas desde fuentes no oficiales o desactualizadas, se comprobó la procedencia de cada alternativa antes de realizar las mediciones.

GNU Make ya se encontraba instalado en el sistema, por lo que no fue necesario descargarlo para realizar la prueba. La versión utilizada es GNU Make 4.4.1. La versión estable 4.4.1 fue publicada oficialmente por el proyecto GNU el 26 de febrero de 2023. Aunque su última versión estable es más antigua que las de las otras alternativas, GNU Make continúa siendo un proyecto mantenido dentro del ecosistema GNU y dispone de un repositorio de desarrollo con miles de commits.

Just se descargó directamente desde la sección de releases de su repositorio oficial de GitHub, `casey/just`, utilizando el binario publicado para Linux. El repositorio continúa activo: la versión 1.59.0 fue publicada el 15 de octubre de 2026, con cambios funcionales, correcciones y modificaciones en su sistema de integración continua.

Task se descargó igualmente desde las releases de su repositorio oficial de GitHub, `go-task/task`. El proyecto muestra actividad reciente: la versión estable 3.54.0 fue publicada el 1 de octubre de 2026 y posteriormente se han publicado versiones nightly en noviembre de 2026.

La existencia de publicaciones recientes se utiliza como evidencia de que Just y Task son proyectos mantenidos activamente. En el caso de GNU Make, se distingue entre la fecha de su última versión estable y la actividad de su repositorio de desarrollo, evitando asumir que una versión estable antigua implica necesariamente que el proyecto esté abandonado.

### Versiones utilizadas

Para realizar las mediciones se han utilizado las versiones estables más recientes disponibles de cada herramienta en el momento de realizar la comparación.

En GNU Make se utiliza la versión 4.4.1, que es la versión estable publicada actualmente por el proyecto GNU.

Para Just se utiliza la versión 1.59.0, publicada el 15 de octubre de 2026 en el repositorio oficial del proyecto.

Para Task se utiliza la versión 3.54.0, publicada el 1 de octubre de 2026 en su repositorio oficial.

En todos los casos se utilizan versiones estables y no versiones nightly, beta o de desarrollo. De esta forma las mediciones representan las versiones que razonablemente se instalarían actualmente en un entorno nuevo de integración continua o contenedores.