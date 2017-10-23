package main

// @APIVersion 1.0.0
// @APITitle in-voicer api
// @APIDescription Generates PDF invoices
// @Contact jcarlos323@hotmail.com
// @License BSD
// @LicenseUrl http://opensource.org/licenses/BSD-2-Clause

import (
    "github.com/juancaacuna/in-voicer-api/invoicer"
	"github.com/gorilla/mux"
    "log"
    "net/http"
    "os"
)

func main() {
    os.Setenv("INVOICER_EMAIL_PASSWORD", "putpasswordhere")

    router := mux.NewRouter()
    api := router.PathPrefix("/api").PathPrefix("/v1").Subrouter()
    api.HandleFunc("/invoices", invoicer.GetInvoices).Methods("GET")
    api.HandleFunc("/send", invoicer.CreateAndSendInvoices).Methods("POST")
    log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), router))
}
