package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	sample "github.com/iguerrero84/gopher-rest-api/cmd/sample-data"
	gopher "github.com/iguerrero84/gopher-rest-api/pkg"

	"github.com/iguerrero84/gopher-rest-api/pkg/server"
	"github.com/iguerrero84/gopher-rest-api/pkg/storage/inmem"
)

func main() {
	withSampleData := flag.Bool("withSampleData", true, "Initialize the API with some Gophers")
	flag.Parse()

	var gophers map[string]*gopher.Gopher
	if *withSampleData {
		gophers = sample.Gophers
	}
	repo := inmem.NewGopherRepository(gophers)
	s := server.New(repo)

	fmt.Println("The gopher server is on tap now: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", s.Router()))
}
