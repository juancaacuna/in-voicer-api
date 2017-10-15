package invoicer

import (
	"encoding/json"
	"net/http"
)

// @Title GetInvoices
// @Description Dummy. Gets basic structure of json, in order to post /send
// @Produce  json
// @Success 200 {array}  Invoice
// @Router /invoices [get]
func GetInvoices(w http.ResponseWriter, r *http.Request) {
	var invoices []Invoice
	invoices = append(invoices, Invoice{})
	json.NewEncoder(w).Encode(invoices)
}

// @Title CreateAndSendInvoices
// @Description Generate PDF invoices given a json input, and send them via email.
// @Accept  json
// @Router /send [post]
func CreateAndSendInvoices(w http.ResponseWriter, r *http.Request) {
	var invoices []Invoice
	err := json.NewDecoder(r.Body).Decode(&invoices)
	if (err == nil) {
		CreateInvoices(&invoices)
		for i, _ := range invoices {
			SendInvoice(&invoices[i])
		}
	}
    return
}
