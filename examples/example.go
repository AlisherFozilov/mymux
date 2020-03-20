package main

import (
	"context"
	"fmt"
	exm "github.com/AlisherFozilov/exactmux/pkg/exactmux"
	"log"
	"net/http"
)

func main() {
	mux := exm.NewExactMux()

	context.Background()
	mux.GET("/hello/{id}/comments/{hello_coursar}/hi/{haha}", f)

	err := http.ListenAndServe("0.0.0.0:9999", mux)
	log.Fatal(err)
}

func f(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Context().Value("id"))
	fmt.Println(r.Context().Value("hello_coursar"))
	fmt.Println(r.Context().Value("haha"))
	fmt.Println(r.URL.Path)
	w.Write([]byte(r.URL.Path))
}
