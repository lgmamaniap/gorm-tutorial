package main

import (
	"fmt"
	"time"
)

/*
Una goroutine es una función que se ejecuta de manera concurrente con otras funciones.
En términos simples, es como un hilo ligero que Go maneja de manera eficiente.
A diferencia de los hilos tradicionales del sistema operativo, las goroutines son gestionadas por el runtime de Go, lo que las hace más eficientes en términos de memoria y tiempo de creación.

1.1. Creando una Foroutine
Para crear una goroutine, se debe usar la palabra clave go antes de la llamada a la función.
*/

func sayHello() {
	fmt.Println("Hello world!")
}

func largeTask() {
	fmt.Println("Task completed!")
}

func main() {
	go sayHello()
	go largeTask()
	fmt.Println("End of main function")
	time.Sleep(2 * time.Second) // Espera para que las goroutines tengan tiempo de ejecutarse
}

/*
En este ejemplo, sayHello() se ejecuta en una goroutine. Sin el tiem.Sleep, el programa principal podría teminar antes de que la goroutine tenga la oportunidad de ejcutarse.

1.2. Goroutines y el Scheduler de Go
Go tiene un scheduler que se encarga de gestionar las goroutines. Este scheduler asigna
tiempo de CPU a las goroutines y las ejecuta en uno o más hilos del sistema operativo. Esto
permite que miles o incluso millones de goroutines se ejecuten de manera eficiente.

2. Indroducción a los channels
Los channels son un
*/
