package main

import (
	"bufio"
	"errors"
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
	lista.tareas = append(lista.tareas[:indice], lista.tareas[indice+1:]...)
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

func (lista *ListaTareas) obtenerIndice(action string) (int, error) {
	fmt.Printf("Ingrese indice de la tarea que desea %s: \n ", action)
	var ind int
	fmt.Scanln(&ind)

	if ind < 0 || ind >= len(lista.tareas) {
		return -1, errors.New("indice fuera de rango")
	}
	return ind, nil
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
			ind, err := lista.obtenerIndice("marcar como completada")
			if err != nil {
				fmt.Println("Error", err)
				break
			}
			lista.marcarCompletada(ind)

		case 3:
			//Editar
			ind, err := lista.obtenerIndice("editar")
			if err != nil {
				fmt.Println("Error", err)
				break
			}
			t := crearTarea("editar")
			lista.editarTarea(ind, t)
		case 4:
			//Eliminar
			ind, err := lista.obtenerIndice("eliminar")
			if err != nil {
				fmt.Println("Error", err)
				break
			}
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
