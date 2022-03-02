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

var templates = template.Must(template.New("T").ParseGlob("templates/**/*.html"))
var errorTemplate = template.Must(template.ParseFiles("templates/error/error.html"))

func handleError(rw http.ResponseWriter, status int) {
	rw.WriteHeader(status)
	errorTemplate.Execute(rw, nil)

}

//Funcion de rendertemplate
func renderTemplate(rw http.ResponseWriter, name string, data interface{}) {
	err := templates.ExecuteTemplate(rw, name, data)
	if err != nil {
		handleError(rw, http.StatusInternalServerError)
	}
}

//Handler
func Index(rw http.ResponseWriter, r *http.Request) {

	//fmt.Fprintf(rw, "Hola mundo")
	//Must valida el error y entre solo el template
	//template := template.Must(template.New("index.html").ParseFiles("index.html", "base.html"))

	usuario := Usuarios{"Daniela", 30}
	renderTemplate(rw, "index.html", usuario)
}
func Registro(rw http.ResponseWriter, r *http.Request) {

	renderTemplate(rw, "regisstro.html", nil)
}

func main() {
	//mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)
	mux.HandleFunc("/registro", Registro)
	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}
	fmt.Println("El servidor está corriendo en el puerto 3000")
	fmt.Println("Run server: http://localhost:3000/")
	log.Fatal(server.ListenAndServe())
}
