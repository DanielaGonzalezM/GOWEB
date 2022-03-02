package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//Estrucuras
type Usuarios struct {
	UserName string
	Edad     int
	Activo   bool
	Admin    bool
	Cursos   []Curso
}
type Curso struct {
	Nombre string
}

//Funciones

//Handler
func Index(rw http.ResponseWriter, r *http.Request) {

	c1 := Curso{"Go"}
	c2 := Curso{"Python"}
	c3 := Curso{"Java"}
	c4 := Curso{"JavaScript"}

	//fmt.Fprintf(rw, "Hola mundo")
	template, error := template.ParseFiles("index.html")
	cursos := []Curso{c1, c2, c3, c4}
	usuario := Usuarios{"Daniela", 30, true, false, cursos}

	if error != nil {
		panic(error)
	} else {
		template.Execute(rw, usuario)
	}
}
func main() {
	//mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)
	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}
	fmt.Println("El servidor está corriendo en el puerto 3000")
	fmt.Println("Run server: http://localhost:3000/")
	log.Fatal(server.ListenAndServe())
}
