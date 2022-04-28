package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
	first example of http.Handler interface
*/

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

// the web server invokes each handler in a new goroutine, so handlers must take precautions such as locking
// when accessing variables that other goroutines, including other requests to the same handler, may be accessing.
func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	// returns ServeMux which also implements http.Handler interface
	mux := http.NewServeMux()
	// HandlerFunc is thus an adapter that lets a function value satisfy an interface,
	// where the function and the interface’s sole method have the same signature.
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	//mux.HandleFunc("/list", db.list)
	//mux.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

// net/http provides a global ServeMux instance called DefaultServeMux and package-level functions called http.Handle and http.HandleFunc.
// To use DefaultServeMux as the server’s main handler, we needn’t pass it to ListenAndServe; nil will do.
/*
func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
*/
