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

var templates = template.Must(template.New("T").ParseFiles("index.html", "base.html"))

//Handler
func Index(rw http.ResponseWriter, r *http.Request) {

	//fmt.Fprintf(rw, "Hola mundo")
	//Must valida el error y entre solo el template
	//template := template.Must(template.New("index.html").ParseFiles("index.html", "base.html"))

	usuario := Usuarios{"Daniela", 30}

	//template.Execute(rw, usuario)
	err := templates.ExecuteTemplate(rw, "index.html", usuario)
	if err != nil {
		panic(err)
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
