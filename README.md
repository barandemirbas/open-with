[![Go Report Card](https://goreportcard.com/badge/github.com/barandemirbas/open-with)](https://goreportcard.com/report/github.com/barandemirbas/open-with)
[![Go Reference](https://pkg.go.dev/badge/github.com/barandemirbas/open-with.svg)](https://pkg.go.dev/github.com/barandemirbas/open-with)
# Getting Started
Install **Open With**
```sh
go get -u github.com/barandemirbas/open-with
```
Add to your imports to start using **Open With**
```go
import "github.com/barandemirbas/open-with"
```
# Functions
```go
openwith.Browser("https://github.com/barandemirbas/open-with")
openwith.Browser("127.0.0.1", 8080) // http://127.0.0.1:8080
```
**Browser** function opens a new browser tab for pointing url with default browser.
# Examples
Example with **Browser** function:
```go
package main

import (
	"github.com/barandemirbas/open-with"
	"net/http"
)

type HttpHandler struct{}

func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello World!"))
}

func main() {
	openwith.Browser("127.0.0.1", 9000)
	http.ListenAndServe(":9000", HttpHandler{})
}
```
