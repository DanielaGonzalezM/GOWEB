package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//Funciones
func Saludar(nombre string) string {
	return ("Hola " + nombre + " desde una funcion")
}

//Handler
func Index(rw http.ResponseWriter, r *http.Request) {

	//fmt.Fprintf(rw, "Hola mundo")
	funciones := template.FuncMap{
		"saludar": Saludar,
	}
	template, error := template.New("index.html").Funcs(funciones).
		ParseFiles("index.html")

	if error != nil {
		panic(error)
	} else {
		template.Execute(rw, nil)
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
