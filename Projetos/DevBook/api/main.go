package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("Rodando a Api")

	log.Fatal(http.ListenAndServe(":5000", router.Gerar()))
}
