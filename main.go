package main

import (
	"fmt"
	"net/http"
)

func main() {
	//Creación de un servidor web //////////////////////////////////
	/*fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	log.Println("Servidor iniciado en: http://127.0.0.1:8080")
	http.ListenAndServe(":8080", nil)*/

	// Handler /////////////////////////////////////
	//Registrando la ruta y el handler
	http.HandleFunc("/saludar", saludar)
	//sube un servidor en el puerto 9080
	http.ListenAndServe(":9080", nil)
}

func saludar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo")
}
