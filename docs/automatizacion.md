# Automatización del proyecto

## Gestión de dependencias

OptiBalsa utiliza Go Modules mediante `go.mod`.

Go incorpora de forma nativa la gestión de módulos y dependencias, por lo que
no es necesario añadir un gestor de dependencias externo al proyecto.

## Gestor de tareas

### Alternativas consideradas

Se estudiarán:

- GNU Make
- Just
- Task

La elección se realizará después de comparar las alternativas.

### Criterios

#### Requisitos de instalación

Se estudiará qué necesita cada herramienta para estar disponible en un entorno
limpio. Este aspecto es relevante de cara a futuros entornos de CI y
contenedores, donde añadir herramientas supone tiempo y componentes adicionales.

#### Tiempo de ejecución

Se medirá el tiempo necesario para lanzar mediante cada herramienta una misma
operación sencilla del proyecto.

La operación será equivalente en las tres alternativas para que los resultados
sean comparables.

## Resultados

| Herramienta | Requisitos de instalación | Tiempo |
|---|---|---:|
| GNU Make | Pendiente | Pendiente |
| Just | Pendiente | Pendiente |
| Task | Pendiente | Pendiente |

## Elección

La herramienta se elegirá después de completar las mediciones anteriores.