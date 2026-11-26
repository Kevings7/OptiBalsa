# Tests

## Herramientas

En Go, las herramientas de test se reparten en tres niveles:

- **Biblioteca de aserciones:** las funciones que comparan el resultado obtenido con el esperado.
- **Framework de test:** la estructura de los tests (funciones de test, subtests, funciones auxiliares, fallos que detienen el test).
- **Ejecutor de tests de línea de órdenes:** el programa que encuentra los tests, los compila, los ejecuta e informa del resultado.

GNU Make, elegido en [automatizacion.md](automatizacion.md), no pertenece a ninguno de estos niveles: la tarea `make test` solo lanza el ejecutor elegido.

### Criterios

Los criterios salen de lo que tienen que comprobar los tests de OptiBalsa:

- **Comparar números decimales con tolerancia.** El plan de riego son porcentajes `float64`, y dos repartos correctos pueden diferir en el último decimal por redondeo.
- **Comprobar el tipo de un error.** Los errores del dominio son tipos propios (`ErrorHumedad`, `ErrorPorcentajeReserva`…), y hay que comprobar que se devuelve el tipo correcto, no solo que hay un error.
- **No escribir en el directorio de los fuentes, o poder configurar dónde escribe.** En el siguiente objetivo los tests se ejecutarán en un contenedor con el directorio del proyecto protegido contra escritura.

Además, el proyecto no usa bibliotecas externas, así que una herramienta externa solo se elegiría si la biblioteca estándar no cumpliera alguno de los criterios.

Las opciones se han buscado en la sección de testing de [Awesome Go](https://github.com/avelino/awesome-go#testing) y en [pkg.go.dev](https://pkg.go.dev/), considerando solo las que funcionan con `go test`.

### Bibliotecas de aserciones

- **Ninguna, solo la biblioteca estándar.** El paquete `testing` no trae aserciones: se compara con `if` y se informa del fallo con `t.Errorf`. Para el tipo de error está [`errors.As`](https://pkg.go.dev/errors#As), y la tolerancia se resuelve con `math.Abs` en una función auxiliar de tres líneas (`casiIgual`). Cumple los tres criterios.
- **[testify](https://github.com/stretchr/testify) (`assert`):** tiene `InDelta` para la tolerancia y `ErrorAs` para el tipo de error. Cumple los criterios, pero es una dependencia externa.
- **[gotest.tools](https://pkg.go.dev/gotest.tools/v3/assert) (`assert`):** tiene `ErrorType` para el tipo de error; la tolerancia necesita además las opciones de [go-cmp](https://pkg.go.dev/github.com/google/go-cmp/cmp/cmpopts#EquateApprox). Son dos dependencias externas.
- **[quicktest](https://github.com/frankban/quicktest):** tiene `ErrorAs`; la tolerancia también depende de go-cmp. Son dependencias externas.

Se usa **la biblioteca estándar, sin biblioteca de aserciones**, porque cumple los tres criterios y las demás solo añaden una forma más corta de escribir lo mismo a cambio de dependencias externas.

### Framework de test

- **Paquete [`testing`](https://pkg.go.dev/testing):** forma parte de Go. Ofrece las funciones `TestXxx`, subtests con `t.Run`, funciones auxiliares con `t.Helper` y fallos que detienen el test con `t.Fatal`.
- **[Ginkgo](https://onsi.github.io/ginkgo/):** framework de estilo BDD con su propio ejecutor, que se usa junto a la biblioteca de aserciones Gomega. Es externo y su estructura de especificaciones no aporta nada a estos tres criterios.

Se usa **el paquete `testing`**.

### Ejecutor de tests

- **`go test`:** forma parte de Go. No escribe en el directorio de los fuentes: compila los tests en un directorio temporal y guarda su caché en `GOCACHE`, y los dos lugares se pueden configurar con las variables `GOTMPDIR` y `GOCACHE`.
- **[gotestsum](https://github.com/gotestyourself/gotestsum):** ejecuta `go test` por debajo y cambia el formato de la salida; solo escribe un fichero si se le pide un informe JUnit. Como depende de `go test`, no mejora ninguno de los criterios.

Se usa **`go test`**, lanzado con `make test`.

## Principios F.I.R.S.T.

- **Fast (rápidos):** ningún test lee ficheros ni accede a la red; todos se ejecutan en milisegundos.
- **Independent (independientes):** cada test crea su propio planificador con la función auxiliar `planificadorConNecesidades` o con `NuevoPlanificadorRiego`, sin compartir estado con otros tests.
- **Repeatable (repetibles):** las fechas están fijadas en el código (`inicioPeriodo`) y no se usa la fecha actual, así que el resultado no depende del día en que se ejecuten.
- **Self-validating (autoevaluables):** cada test termina con `t.Errorf` o `t.Fatal` si algo no se cumple, sin necesidad de revisar la salida a mano.
- **Timely (oportunos):** cada test se ha añadido en el mismo commit que el código que resuelve su issue.
