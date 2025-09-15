package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/paulllo-victor/go-hexagonal/adapters/web/handler"
	"github.com/paulllo-victor/go-hexagonal/application"
	"github.com/urfave/negroni/v3"
)

type Webserver struct {
	Service application.ProductPersistenceInterface
}

func MakeNewWebserver() *Webserver{
	return &Webserver{}
}

func (w Webserver) Serve(){

	r := mux.NewRouter();
	n := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakeProductHandlers(r,n,w.Service)
	http.Handle("/", r)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr: ":8080",
		Handler: http.DefaultServeMux,
		ErrorLog: log.New(os.Stderr, "log", log.Lshortfile),
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}