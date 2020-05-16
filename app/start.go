package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"walk-the-camino/utils"
)

func Start() (err error) {
	// create http server to accept transaction
	errCh := make(chan error)
	go func(chan error) {
		url := fmt.Sprintf("%s:%s", os.Getenv(utils.HttpListenAddress), os.Getenv(utils.HttpPort))
		router := mux.NewRouter().StrictSlash(true)
		router.HandleFunc("/employee", createEmployee).Methods("POST")
		router.HandleFunc("/employees", getAllEmployee).Methods("GET")
		router.HandleFunc("/employees/{id}", getOneEmployee).Methods("GET")
		router.HandleFunc("/employees/{id}", updateEmpoyee).Methods("PUT")
		router.HandleFunc("/employees/{id}", deleteEmployee).Methods("DELETE")
		router.HandleFunc("/flushData", flushEmployee).Methods("GET")
		if os.Getenv(utils.TLSEnabled) == "true" {
			log.Println("starting the http server on tls :", url)
			if err := http.ListenAndServeTLS(url, os.Getenv(utils.Cert), os.Getenv(utils.Key), router); err != nil {
				errCh <- err
			}
		} else {
			log.Println("starting the http server on non tls :", url)
			if err := http.ListenAndServe(url, router); err != nil {
				errCh <- err
			}
		}
	}(errCh)
	return <-errCh
}
