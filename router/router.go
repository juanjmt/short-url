package router

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"short-url/handle"
)

func Get() {
	router:=mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/ping", handle.Ping).Methods("GET")
	router.HandleFunc("/acortar", handle.Acortar).Queries("url","{url}").Methods("GET")
	router.HandleFunc("/revertir", handle.Revertir).Queries("url","{url}").Methods("GET")
	log.Fatal( http.ListenAndServe(":3000",router))
}


