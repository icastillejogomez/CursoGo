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

## Workspace

Go por defecto trabaja con workspaces. El workspace por defecto lo podemos averiguar ejecutando el siguiente comando aunque en las nuevas versiones de Go este se establece automáticamente al instalar el ejecutable de Go en nuestra carpeta Home bajo el directorio go, es decir: `~/go`

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
