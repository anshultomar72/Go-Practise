package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	greeter()

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("heyaaaa")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Ya i gotcha </h1>"))
}

//Some cli commands for mod related task
//any operation including go mod ---- are expensive
// go build .
// go mod tidy
// go mod verify --- verifies ur modules using the hash present in go.sum
// go list --- list all the modules on which ur project is dependent doesnt list all
// go list all --- lists all the modules present in the cache
// go list -m all --- lists all the modules on which it is actually dependent
// go list -m -versions github.com/gorilla/mux --- displays all the versions of something till now
// go mod why github.com/gorilla/mux --- tells u why ur modules is dependent on this
// go mod graph --- displays a list which modules depend on which
// go mod edit --- to edit the mod file
// go mod vendor --- just like nodemodules