package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Shulammite-Aso/bazel-demo-app/pkg/greetings"
)

func Greet(w http.ResponseWriter, r *http.Request) {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	greeting, err := greetings.Hello("Shula")

	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(greeting))

}

func GreetMany(w http.ResponseWriter, r *http.Request) {

	names := []string{"Prisca", "Nana", "Derin"}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(messages)
}
