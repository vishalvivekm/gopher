package main

import (
	"fmt"
	"net/http"
)

type name string

func main() {
	student := name("Vivek Vishal")

	fmt.Print(student)

	http.HandleFunc("/", Handler)
	err := http.ListenAndServe("2000", nil)
	if err != nil {
		return
	}

}
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello there")
}
