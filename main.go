package main

import (
	"fmt"
	"net/http"

	"github.com/fulit103/truoratest/models"
	"github.com/fulit103/truoratest/routes"
	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("init Server")

	models.InitDB()

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Truora api by Julian Toro"))
	})

	r.Get("/domains/{domain}", routes.InfoDomainEndPoint)

	r.Get("/domains", routes.ListDomainsEndPoint)

	http.ListenAndServe(":4000", r)
}
