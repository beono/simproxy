[![Build Status](https://travis-ci.org/beono/simproxy.svg?branch=master)](https://travis-ci.org/beono/simproxy)
[![GoDoc](https://godoc.org/github.com/beono/simproxy?status.svg)](https://godoc.org/github.com/beono/simproxy)
[![Go Report Card](https://goreportcard.com/badge/github.com/beono/simproxy)](https://goreportcard.com/report/github.com/beono/simproxy)

# About the project
The project is an example of an application that uses reverse proxy from the standartd library.

# Why
Let's assume you want to migrate your api that from language X to golang.
Sometimes it's hard to rewrite everything at once and you want to move some handlers to golang, but use old application as a fallback.
I just want to show that you don't have to use `nginx` or write your own proxy server based of standard http client.

# How to use it
The solution is quite simple and straightforward.
We use [reverse proxy](https://golang.org/pkg/net/http/httputil/#NewSingleHostReverseProxy) and http server from the standard library.
Just clone or `go get -u github.com/beono/simproxy` this project.
Run `go test && go run main.go` to see it in action.

Send a request to http://localhost/hello and you will get a tiny text response.

Send a request to http://localhost:8080/oauth2/v3/certs and you will get a json response.
This request was proxied to https://www.googleapis.com/oauth2/v3/certs, because this handler doesn't exist in our new application.


