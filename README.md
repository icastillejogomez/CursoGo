# Curso de Go

## Introducción

Go es un lenguaje de programación concurrente y compilado inspirado en la sintaxis de C. Ha sido desarrollado por Google, y sus diseñadores iniciales son Robert Griesemer, Rob Pike y Ken Thompson. 

Actualmente está disponible en formato binario para los sistemas operativos Windows, GNU/Linux, FreeBSD y Mac OS X, pudiendo también ser instalado en estos y en otros sistemas con el código fuente.2​3​ Go es un lenguaje de programación compilado, concurrente, imperativo, estructurado, orientado a objetos —de una manera bastante especial— y con recolector de basura que de momento está soportado en diferentes tipos de sistemas UNIX, incluidos Linux, FreeBSD y Mac OS X. 

También está disponible en Plan 9, puesto que parte del compilador está basado en un trabajo previo sobre el sistema operativo Inferno. Las arquitecturas soportadas son i386, amd64 y ARM.

### Nombre

El día de la publicación del lenguaje Go, Francis McCabe, desarrollador del lenguaje de programación Go! (anteriormente llamado Go), solicitó que se le cambiase el nombre al lenguaje de Google para evitar confusiones con su lenguaje.4​ McCabe creó Go! en el año 2003; sin embargo, aún no ha registrado el nombre.5​ Go es un nuevo lenguaje de programación para sistemas lanzado por Google en noviembre de 2009. Aunque empezó a ser desarrollado en septiembre de 2007 por Robert Griesemer, Rob Pike y Ken Thompson.

### Características

1. Go usa una sintaxis parecida a C, por lo que los programadores que hayan usado dicho lenguaje se sienten muy cómodos con él.
2. Go usa tipado estático (statically typed) y es tan eficiente como C.
3. Go tiene muchas de las características y facilidad de lenguajes dinámicos como Python
4. Aún siendo un lenguaje diseñado para la programación de sistemas, provee de un recolector de basura, reflexión y otras capacidades de alto nivel que lo convierten en un lenguaje muy potente.
5. Go admite el paradigma de programación orientada a objetos, pero a diferencia de los lenguajes de programación más populares no dispone de herencia de tipos y tampoco de palabras clave que denoten claramente que soporta este paradigma. Otro detalle que puede resultar confuso es que la definición de un tipo ("clase") se realiza por medio de declaraciones separadas (interfaces, structs, embedded values). Go permite el uso de delegación (a través de embedded values) y polimorfismo (por medio de interfaces).

## Instalación

Dirígete a la [página web](https://golang.org/dl/) de Go para descargar los binarios de tu sistema operativo.

Para comprobar si se ha instalado correctamente puedes ejecutar el siguiente comando:

```
$ go version
```

Un ejemplo en MacOS es el siguiente:

```
$ go version
go version go1.10.3 darwin/amd64
```

Si vas a instalar Go en Windows porque no eres capaz de esforzarte en aprender Linux lee este artículo para que no te explote la cabeza: [Instalación de GO (Golang)
](https://medium.com/@golang_es/instalaci%C3%B3n-de-go-golang-6fd5d7b9eb48)

## Primeros pasos con Go después de instalarlo

### Workspace ¿Qué es?

Go trabaja con workspaces. Un workspace no es más que la carpeta donde estará tu proyecto de go.

El workspace por defecto lo podemos averiguar ejecutando el siguiente comando aunque en las nuevas versiones de Go este se establece automáticamente al instalar el ejecutable de Go en nuestra carpeta Home bajo el directorio go, es decir: `~/go`

```
$ go env GOPATH
/Users/<username>/go
```

o en Linux, por defecto:

```
$ go env GOPATH
/home/<username>/go
```

Podemos modificar el workspace con el siguiente comando:

```
$ export GOPATH="/tu/ruta"
```

El cual exporta una variable de entorno en el sistema operativo para que el ejecutable de Go pueda obtener su contenido.


Bueno, empezemos configurando nuestro espacio de trabajo. Yo, en esta caso, tengo la carpeta donde quiero guardar mi trabajo en la ruta `~/Github/CursoGo` por lo que ejecutaré el siguiente comando:

```
$ export GOPATH="/Users/<username>/Github/CursoGo"
```

O bien, estando con el terminal en la carpeta que queramos que sea nuestro workspace ejecutamos:

```
$ export GOPATH="$(pwd)"
```

Dentro de esta carpeta crearemos la siguiente estructura de carpetas, la cual no va a tener ninguna lógica para el lector nobel pero más adelante le encontraremos la explicación. **Tenéis que hacer un acto de fe**.

```
$ mkdir -p $GOPATH/src/github.com/icastillejogomez/CursoGo
```

Una explicación muy breve es: Crearemos todo nuestro código dentro de la carpeta `src` dentro de nuestro workspace. Por motivos de importación (que se verá más adelante) tenemos que crear una estructura de carpetas que coincida con el repositorio Git en el que guardaremos nuestro código.

### Hola mundo

Crearemos un archivo llamado `main.go` dentro de la carpeta `CursoGo/HolaMundo`:

```
$ mkdir HolaMundo
$ touch HolaMundo/main.go
```

Dentro de este archivo añadiremos el siguiente contenido:

```
$ cat HolaMundo/main.go
package main

// Importamos la libreria para poder imprimir mensajes por consola
import "fmt"

// Punto de entrada de nuestro programa
func main() {
	fmt.Print("Hola mundo\n")
}
```
###### Puedes encontrar el archivo [aquí.](https://github.com/icastillejogomez/CursoGo/blob/master/src/github.com/icastillejogomez/CursoGo/HolaMundo/main.go)

Y para correr nuestro hola mundo ejecutaramos el siguiente comando:

```
$ go run HolaMundo/main.go
Hola mundo
```

Y si! Ya se que dijé que Go es un lenguaje compilado y aquí esta la prueba:

Compilamos nuestro código a un archivo ejecutable con nombre `main`:
```
$ go build -o HolaMundo/main HolaMundo/main.go
```

Comprobamos que nos ha generado un archivo `main` con permisos de ejecución:
```
$ ls -la HolaMundo/
total 4104
drwxr-xr-x  4 nacho  staff      128 Nov 12 08:31 .
drwxr-xr-x  3 nacho  staff       96 Nov 12 08:30 ..
-rwxr-xr-x  1 nacho  staff  2093184 Nov 12 08:31 main
-rw-r--r--@ 1 nacho  staff      177 Nov 12 08:27 main.go
```

Ejecutamos nuestro programa y obtenemos el mismo resultado:
```
$ HolaMundo/main
Hola mundo
```

Podemos realizar una prueba de velocidad de la siguiente manera:

```
$ time HolaMundo/main
Hola mundo

real	0m0.006s
user	0m0.001s
sys	0m0.004s
```

¿Qué tan rápido es 0.006s? ¿6 milisegundos es lento? Te dejo la busqueda de esas respuestas a ti!

### Hola Mundo 2.0

Vamos a ver un ejemplo un poco más complejo pidiendo datos al usuario mediente el teclado de su ordenador, guardandolo en una variable e imprimiendolo en la consola.

Para este ejemplo lo primero que necesitamos saber es: **¿Cómo puedo crear una variable en Go?**

Lo haremos de la siguiente manera :

```
var <nombre_variable> <tipo_variable>
```

Por ejemplo para crear una variable de tipo nombre llamada `name` de tipo String lo haremos de la siguiente forma:

```
var name string
```

Lo siguiente que demos saber hacer es **pedir datos al usuario mediante el teclado de su ordenador**. Para hacer esto usaremos una función de la libreria `fmt` llamada `Scanf`. Con esta función podemos pedir un dato al usuario para que lo introduzca mediante su teclado y guardar el contenido en una variable que nosotros queramos. 

Usando la variable `name` del ejemplo anterior quedaría de la siguiente manera: 

```
fmt.Scanf("%s", &name)
```

1. Con `%s` le estamos diciendo a Scanf que queremos que el usuario nos devuelva contenido de tipo `string`.
2. Con `&name` le estamos diciendo a Scanf que guarde el contenido en la variable `name`.

Con todo esto estamos listos para crear el siguient ejemplo:

```
$ cat HolaMundo2.0/main.go
package main

// Importamos la libreria para poder imprimir mensajes por consola
import "fmt"

// Punto de entrada de nuestro programa
func main() {

	// Creamos la variable donde vamos a guardar el nombre
	var name string

	// Pedimos el nombre del usuario
	fmt.Print("Por favor, indique su nombre: > ")
	fmt.Scanf("%s", &name)

	// Imprimimos el nombre del usuario por pantalla
	fmt.Printf("Registro completado con exito %s\n", name)
}
```
###### Puedes encontrar el archivo [aquí.](https://github.com/icastillejogomez/CursoGo/blob/master/src/github.com/icastillejogomez/CursoGo/HolaMundo2.0/main.go)

Cuyo resultado si lo ejecutamos es:

```
$ go run HolaMundo2.0/main.go
Por favor, indique su nombre: > Ignacio
Registro completado con exito Ignacio
```

Y si... te estarás pregutando **¿Qué es Printf?** Es una función que nos ofrece la libreria `fmt` para concatenar un mensaje que queremos imprimir con variables. De esta forma cuando le estamos escribiendo `%s` le estamos diciendo que ahí irá un texto. **¿Cuál?** El que nosotros le pasamos a continuación que es `name`. 

**¿Es facil verdad?**

Solo nos queda aclarar algo de lo que seguro ya te has dado cuenta... ¿Cómo podemos poner comentarios? O... ¿Qué narices es un comentario?

### Comentarios

Un comentario es una zona de nuestro código que **no** será ejecutada. Se utilizan como anotaciones para explicar de una manera sencilla que hace el código que les acompaña. Pensad que en un proyecto en el que trabajan muchas personas y que va a durar en el tiempo se escribirá código que no volverá a ser visto pasado el suficiente tiempo como para que ni el que lo escribió se acuerde de que hace su propio código.

Existen dos formas de comentario:

1. Comentarios de lineas. Son los que habéis visto hasta ahora. Se inician con `//` y indican que desde ahí hasta **el final de esa misma linea** todo será interpretado como un comentario a la hora de ejecutar el programa.
2. Comentarios de varias lineas. Como podéis imaginar los comentarios de varias lineas afectan a varias lineas. Se inician con `/*` y terminan con `*/` y todo lo que haya entre medias será interpretado como un comentario a la hora de ejecutar el programa.

