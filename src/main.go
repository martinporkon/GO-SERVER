package main

import "net/http"

type fooHandler struct {
	Message string
}

// in order to have at this interface we need a ServeHttp func handler

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(f.Message)) //<< wrties out the message field to the HTTP response using the message writer
}

func main() {
	http.Handle("/foo", fooHandler)
}
