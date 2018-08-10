package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("------ START --------")
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("value: ", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello World! %s", time.Now())
	fmt.Println("------ END --------")

}

func main() {
	http.HandleFunc("/", sayHelloName)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	} else {
		fmt.Print("Server listening on http://localhost:9090...")
	}
}
