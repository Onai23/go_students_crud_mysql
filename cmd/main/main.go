package main

import (
	"fmt"
	"github.com/Onai23/go_students_crud_mysql/pkg/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// inisialisasi routes
	r := mux.NewRouter()
	routes.RegisterStudentRoutes(r)
	http.Handle("/", r)
	fmt.Print("Starting Server localhost:9010")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
