package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tarea struct {
	nombre      string
	descripcion string
	completada  bool
}

type ListaTareas struct {
	tareas []Tarea
}

func (lista *ListaTareas) mostrarTareas() {
	fmt.Println("Listado de Tareas:")
	fmt.Println("===================")
	for ind, t := range lista.tareas {
		fmt.Printf("%d. %s - %s - completada: %t \n", ind, t.nombre, t.descripcion, t.completada)
	}
	fmt.Println("===================")
}

func (lista *ListaTareas) agregarTarea(t Tarea) {
	lista.tareas = append(lista.tareas, t)
}

func (lista *ListaTareas) marcarCompletada(indice int) {
	lista.tareas[indice].completada = true
}

func (lista *ListaTareas) editarTarea(indice int, t Tarea) {
	lista.tareas[indice] = t
}

func (lista *ListaTareas) eliminarTarea(indice int) {
	lista.tareas = append(lista.tareas[:indice], lista.tareas[indice+1]...)
}

func displayMenuSelectOption() int {
	fmt.Print(
		"Seleccione una opción: \n",
		"1. Agregar tarea\n",
		"2. Marcar tarea como completada\n",
		"3. Editar Tarea\n",
		"4. Eliminar Tarea\n",
		"5. Salir\n",
		"Ingrese la opción:\n",
	)
	var option int
	fmt.Scanln(&option)
	return option
}

func crearTarea(action string) Tarea {
	fmt.Printf("ingrese el nombre de la tarea que desea %s: \n", action)
	leer := bufio.NewReader(os.Stdin)
	nombre, _ := leer.ReadString('\n')
	fmt.Printf("Ingrese descripcion de la tarea que desea %s: \n", action)
	descr, _ := leer.ReadString('\n')
	return Tarea{nombre: nombre, descripcion: descr}
}

func obtenerIndice(action string) int {
	fmt.Printf("Ingrese indice de la tarea que desea %s: \n ", action)
	var ind int
	fmt.Scanln(&ind)
	return ind
}
func main() {
	lista := ListaTareas{}

	for {
		option := displayMenuSelectOption()

		switch option {
		case 1:
			//Agregar tarea
			t := crearTarea("agregar")
			lista.agregarTarea(t)
		case 2:
			//Marcar como completada
			ind := obtenerIndice("marcar como completada")
			lista.marcarCompletada(ind)

		case 3:
			//Editar
			ind := obtenerIndice("editar")

			t := crearTarea("editar")
			lista.editarTarea(t)
		case 4:
			//Eliminar
			ind := obtenerIndice("eliminar")
			lista.eliminarTarea(ind)
		case 5:
			fmt.Println("Saliendo del programa...")
			return
		default:
			fmt.Println("Opcion ingresada no es válida")
		}

		lista.mostrarTareas()
	}
}
