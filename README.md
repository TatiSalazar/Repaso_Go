# Repaso_Go
Repaso de conceptos de Go

1. Gran velocidad de compilacion
2. Alto rendimiento para tareas pesadas
3. Soporte nativo para concurrencia
4. Top 5 amados y mejores pagados
5. Oboliga a implementar buenas practicas
6. Comunidad muy receptiva

## ¿Quienes usan Go?
1. MercadoLibre: 70mil con 20NB ram
2. Twitch: Usuarios concurrente
3. Twitter: Procesar analitica de la app 
4. Uber: Posicion conductores y pasajeros
5. Docker y Kubernetes : Para despliegue de apps

* A Tour of Go
http://tour.golang.com/

* play-with-go.dev
https://play-with-go.dev/

* Go by Example
https://gobyexample.com/

* Slack
http://gophers.slack.com/

* Spotify
https://open.spotify.com/show/2cKdcxETn7jDp7uJCwqmSE?si=q88UkEYQTxS0t1QVws22tw&amp;nd=1

## Valores Primitivos
* Al codificar en Go podemos especificar el tipo de dato, permitiéndonos ganar gran preformas en nuestro código

## Números Enteros
* int Cuando no se declara el tamaño tomara la referencia del OS (Sistema operativo) (32 o 64 bits)
* int8 8bits ⇒ -128 a 127
* int16 16bits ⇒ -2^15 a 2^15-1
* int32 32bits ⇒ -2^31 a 2^31-1
* int64 64bits ⇒ -2^63 a 2^63-1

## Optimizar memoria cuando sabemos que el dato siempre va ser positivo
* uint ⇒ Depende del Sisema Operativo (32 o 64 bits)
* uint8 ⇒ 8bits = 0 a 127
* uint16 ⇒ 16bits = 0 a 2^15-1
* uint32 ⇒ 32bits = 0 a 2^31-1
* uint64 ⇒ 64bits = 0 a 2^63-1

## Números decimales
* float32 ⇒ 32 bits = +/- 1.18e^-38 +/- -3.4e^38
* float64 ⇒ 64 bits = +/- 2.23e^-308 +/- -1.8e^308

## Textos y boleanos
* A diferencia de otros lenguajes de programación donde para definir una variable de tipo string es permitido utilizar “”, ‘’ o ```` en Go solo podemos utilizar las comillas dobles ""
* string ⇒ ""
* bool ⇒ Trueo False

## Números complejos
* Complex64 ⇒ Real e Imaginario float32
* Complex128 ⇒ Real e Imaginario float64
* Ejemplo: c := 10 * 8i

## Documentacion oficial
* https://pkg.go.dev/std
* https://pkg.go.dev/

## realizar un ciclo en Golang:

For condicional
For while
For forever
For Range
For con funciones
For con goto tags

## Operadores de comparacion
* valor1 == valor2
* valor1 != valor2
* valor1 < valor2
* valor1 > valor2
* valor1 >= valor2
* valor1 <= valor2

## Operadores logicos
#### AND -> &&
* Este operador indica que todas las condiciones declaradas deben cumplirse para poderse marcar como TRUE. En Go, se utiliza este símbolo &&.
* 1>0 && 2 >0 Esto retornará TRUE porque tanto la primera como la segunda condición son verdaderas.

#### OR -> ||
* Este operador indica que al menos una de las condiciones debe cumplirse para marcarse como TRUE. En Go, se representa con el símbolo ||.
* 2<0 || 1 > 0 Esto retornará TRUE porque la segunda condición se cumple, a pesar que la primera no.

#### NOT -> !
* Este operador retornará el opuesto al boleano que está dentro de la variable. Ejemplo:
* myBool :=  true
* fmt.Println(!myBool)

#### Arrays - Slices
* En Go, los arrays poseen un tamaño fijo y son inmutables, mientras que en 
* los slices su tamaño es dinámico y los puedes modificar.
* La diferencia principal entre los arrays es que estos tienen una longitud fija e invariable y deben declarase especifiandola
```go
x := [5]int{0, 1 ,2, 3, 4}
```

* mientras que los Slices tienen una longitud variable y no hay que especificarla en la declaración
```go
var x [ ]float64
x := make([]float64, 5)
x := make([]float64, 5, 10) //representa un Slice de longitud 5 y capacidad de 10
```

## Comandos de Go modules
```go
// Inicializar un proyecto
go mod init path_del_proyecto

// Verificar que el código externo no esté corrupto
go mod verify

// Reemplazar fuente del código
go mod edit -replace path_del_repo_online=path_del_repo_en_local

// Quitar el replace
go mod edit -dropreplace path_del_repo_online

// Empaquetar todo el código de terceros que usa nuestro código
go mod vendor

// Eliminar todos los paquetes externos que no estemos usando
go mod tidy

// Aprender más de go modules
go help mod
```
## Zero values de primitivos
```go
var a int // 0
var b float64 // 0
var c string // ""
var d bool // false
```

## Punteros
```go
a := 10 // Variable int
b := &a // "b" es el puntero de "a"
c := *b // "c" adquiere el valor del puntero de "b", es decir toma el mismo valor de "a"
```
## FrontEnd - Hugo
* Hugo es el framework más popular para usar Go en el FrontEnd, claro no es que se utilice en FrontEnd específicamente sino que tiene su propio template para generar archivos estáticos (sin BackEnd, solo HTML, CSS y JavaScript).

* Pero la mayor ventaja de Hugo, es que tienen Hugo Themes una tienda de templates donde muchos miembros de la comunidad suben sus temas para disponerlos al resto de la comunidad para su uso, solo debes prestar atención a la licencia. Pero encontrarás temas con todo tipo de features, desde uso para portafolio, compatibilidad con Google Anlytics y mucho más.

* Y si lo combinas con Vercel podrás desplegar tu aplicación de forma gratuita.

## Backend
* Para el caso de BackEnd tenemos muchas opciones, las que te recomiendo probar son las siguientes:

* **Echo**
* Echo es el Framework que más he llegado a usar en las Apps en Go que he realizado. Esto porque reúne varias características:
* Es minimalista, evitando tener código que no vamos a usar
* Fácil de escalar y un excelente rendimiento al momento de compilar
* Tiene su sistema de enrutamiento, pudiendo incluso implementar middlewares (proceso de autenticación)
* Tiene su propia interfaz para implementar websockets (conexiones en tiempo real)
* Puedes usarlo tanto para crear APIs como páginas web completas.
* Es uno de los pocos framewords que tiene su propia implementación de JWT (otro tipo de autenticación muy usado en APIs)

* **Gin-Gonic**
* Gin-Gonic Es un framework al que fácilmente le puedes agarrar cariño y esto es porque la forma como interactúas para crear Apps es muy intuitiva. De momento no es de los Frameworks que más destaque, pero tampoco se queda atrás. Ideal para tus primeros proyectos de BackEnd con Go.

* **Beego**
* Beego por bastante tiempo llegó a ser uno de los Frameworks más usados, no es que haya dejado de serlo, pero trajo a la comunidad 3 grandes features:
* Su propio ORM (Object Relational Mapping), básicamente son un conjunto de métodos para interactuar con la base de datos desde consulta a creación de tablas.
* Un CLI (Command Line Interface) llamado bee con el cual puedes generar proyecto ya sea para hacer una API como página web completa, empaquetar el proyecto listo para llevarlo a producción y hacer migraciones de bases de datos.
* Un admin, ¡sí un admin! Con el cual puedes ver estadísticas de las rutas de tu app, verificar el estado de la base de datos y ejecutar tareas de forma manual
* Hoy en día Beego tiene muchos más features como por ejemplo crear el Dockerfile para que tu app se ejecute en Docker. Si vas a crear múltiples apps seguidamente, te recomiendo este Framework.


```go
package main

import "fmt"

func main() {
	// Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.14
	fmt.Println(pi, pi2)

	// Declaracion de variables enteras
	base := 12
	var altura int = 14
	var area int
	fmt.Println(base, altura, area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println(a, b, c, d)

	// Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("area cuadrado: ", areaCuadrado)

	// si no le asignas el valor a una variable,
	// el programa asigna un valor por defecto (no coloca nulos)
}
```
