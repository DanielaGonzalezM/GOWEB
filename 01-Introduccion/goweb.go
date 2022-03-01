package main

import (
	"fmt"
	"log"
	"net/http"
)

//Handlers
func Hola(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("El metodo es +", r.Method)
	fmt.Fprintln(rw, "Hola Mundo de GOWEB")
}
func PaginaNF(rw http.ResponseWriter, r *http.Request) {
	http.NotFound(rw, r)
}
func Error(rw http.ResponseWriter, r *http.Request) {
	http.Error(rw, "La pagina no funciona", http.StatusNotFound)
}
func Saludar(rw http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Println(r.URL.RawQuery)
	fmt.Println(r.URL.Query())
	name := r.URL.Query().Get("name")
	edad := r.URL.Query().Get("edad")
	fmt.Fprintf(rw, "Hola, %s tu edad es %s!!", name, edad)
}

func main() {
	//Mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", Hola)
	mux.HandleFunc("/page", PaginaNF)
	mux.HandleFunc("/error", Error)
	mux.HandleFunc("/saludar", Saludar)

	//Router
	//http.HandleFunc("/", Hola)
	//http.HandleFunc("/page", PaginaNF)
	//http.HandleFunc("/error", Error)
	//http.HandleFunc("/saludar", Saludar)

	//Crear servidor
	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}
	fmt.Println("El servidor está corriendo en el puerto 3000")
	fmt.Println("Run server: http://localhost:3000/")
	//log.Fatal(http.ListenAndServe("localhost:3000", mux))
	log.Fatal(server.ListenAndServe())
}

//pag para ver los status
//https://developer.mozilla.org/es/docs/Web/HTTP/Status
//https://pkg.go.dev/net/http#pkg-constants
