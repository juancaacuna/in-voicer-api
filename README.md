in-voicer-api with GoLang
==========

This API provides an easy way to generate PDF invoices and send them via email. You do the business, we send the invoices.

It was developed with [GoLang](https://golang.org/), which is an open source programming language that makes it easy to build simple, reliable, and efficient software.


To run it locally...
-----

1. 
```bash
$ go get github.com/juancaacuna/in-voicer-api
```
2. 
```bash
$ cd $GOPATH/src/github.com/juancaacuna/in-voicer-api
```
3. 
```bash
$ export INVOICER_EMAIL_PASSWORD=putemailsenderpasswordhere
$ export INVOICER_TEMPLATES_PATH=$GOPATH/src/github.com/juancaacuna/in-voicer-api/invoicer/templates
```
And update email information on `email.go`.
4. 
```bash
$ go run *.go
```


Testing
-----
To run Unit Tests just do:

```bash
$ go test in-voicer/invoicer
```
