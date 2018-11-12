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