package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GiovanniASBF/api_rest_hello-world/controllers"
)

func main() {
	http.HandleFunc("/deploy", controllers.RunContainer)
	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Println("Iniciando o servidor Rest com Go")
}
