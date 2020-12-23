package main

import "net/http"

type fooHandler struct {
	Message string
}

// in order to have at this interface we need a ServeHttp func handler

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(f.Message)) //<< wrties out the message field to the HTTP response using the message writer
}

func main() { // see fail on main fail serveri alustamiseks ning tööle panemiseks.
	http.Handle("/foo", &fooHandler{Message: "foo called"}) // set the message
	// set the serer to listen and serve
	http.HandleFunc("/bar", barHandler) // bar patterna and the HTTP handleFunc funciton
	http.ListenAndServe(":5000", nil)   // nil for the handler and ServeMux. This will tell it to use the default ServeMux
}

// here is the simpler Http call function

func barHandler(w http.ResponseWriter, r *http.Request) {
	// accepts a response writer and a pointer to the request
	w.Write([]byte("bar called"))
}
