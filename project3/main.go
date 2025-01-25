package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ShreyashPG/project3/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Server is running on port 9010");
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
