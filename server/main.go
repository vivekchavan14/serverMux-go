package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"server/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)
	//servemux
	sm := http.NewServeMux()
	sm.Handle("/tryme", hh)
	sm.Handle("/gogogo", gh)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Oops! Something went wrong"))
			return
		}
		log.Printf("Data %s", data)

	})

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "This is a test route")
		log.Printf("Data %s", data)

	})
	http.HandleFunc("/bye", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Good-bye, World!")
	})

	http.ListenAndServe(":8080", sm)
}
