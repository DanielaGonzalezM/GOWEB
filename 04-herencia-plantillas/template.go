package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//Estructuras
type Usuarios struct {
	UserName string
	Edad     int
}

//Handler
func Index(rw http.ResponseWriter, r *http.Request) {

	//fmt.Fprintf(rw, "Hola mundo")

	template, error := template.New("index.html").ParseFiles("index.html", "base.html")

	usuario := Usuarios{"Daniela", 30}

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
