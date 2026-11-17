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

Las alternativas se compararán ejecutando operaciones equivalentes del proyecto y no únicamente una orden trivial.

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