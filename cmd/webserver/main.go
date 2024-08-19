package main

import (
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/whatsapp", handler)
	http.ListenAndServe(":80", nil)
}
