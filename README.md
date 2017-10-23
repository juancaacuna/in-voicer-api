in-voicer-api with GoLang
==========

This API provides an easy way to generate PDF invoices and send them via email. You do the business, we send the invoices.

It was developed with [GoLang](https://golang.org/), which is an open source programming language that makes it easy to build simple, reliable, and efficient software.

It's currently deployed on [Heroku](https://www.heroku.com), which is a cloud platform based on a managed container system, with integrated data services and a powerful ecosystem, for deploying and running modern apps.

Ready for sending some invoices?
-----
The [API is live and running here!](https://in-voicer-api.herokuapp.com/api/v1/invoices)
```
/invoices (GET): Dummy, for getting invoices json structure. No invoice is stored.
/send (POST): Using the invoices json structure, use this method to send the PDF invoices via email.
```

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
$ export PORT=8000
```
And update email information on `email.go`, like sender email, host and port.

4. Download [wkhtmltopdf](https://wkhtmltopdf.org/downloads.html).

5. Run the in-voicer-api :)
```bash
$ go run *.go
```


Testing
-----
To run Unit Tests just do:

```bash
$ go test in-voicer/invoicer
```
