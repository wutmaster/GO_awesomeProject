package main


import (
"log"
"net/http"
"net/http/httptest"
	"fmt"
)
ts := httptest.NewServer(http.HandlerFunc(func(w ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
fmt.Fprintln(w, "Hello, client")
}))
defer ts.Close()

res, err := http.Get(ts.URL)
if err != nil {
log.Fatal(err)
}