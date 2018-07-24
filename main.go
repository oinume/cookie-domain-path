package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	domain = flag.String("domain", "", "domain")
)

func main() {
	flag.Parse()
	if *domain == "" {
		log.Fatal("Flag -domain required")
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/cookie/set", setCookie)
	mux.HandleFunc("/cookie/get", getCookie)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:   "c",
		Value:  "hello",
		Domain: "." + *domain,
		Path:   "/cookie",
	}
	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "setCookie ok")
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("c")
	if err != nil {
		http.Error(w, "Failed to get cookie", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Cookie value is %s", cookie.Value)
}
