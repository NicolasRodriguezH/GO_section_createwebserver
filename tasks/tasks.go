package main

import "fmt"

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

func (t *taskList) eliminarDeLista(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func (t *taskList) agregarALista(tl *task) {
	t.tasks = append(t.tasks, tl)
}

type taskList struct {
	tasks []*task
}

func (t *task) marcarCompleta() {
	t.completado = true
}

func (t *task) actualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t *task) actualizarNombre(nombre string) {
	t.nombre = nombre
}

func main() {
	/* Se apunta a la memoria de ese struct "task" */
	t1 := &task{
		nombre:      "Completar Ts",
		descripcion: "Voy a culminar curso YOUTUBR",
	}
	t2 := &task{
		nombre:      "Completar Js",
		descripcion: "Voy a culminar curso",
	}
	t3 := &task{
		nombre:      "Completar cursoGO",
		descripcion: "Voy a culminar curso GOYOUTUBR",
	}

	lista := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}
	lista.agregarALista(t3)

	for i := 0; i < len(lista.tasks); i++ {
		fmt.Println("Index", i, "nombre", lista.tasks[i].nombre)
	}

	for _, t := range lista.tasks {
		fmt.Println("Nombre:", t.nombre)
	}

	/* for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	} */

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
}
