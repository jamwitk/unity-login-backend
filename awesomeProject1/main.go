package main

import (
	"awesomeProject1/crypto"
	"net/http"
)

func main() {

	encrypted := crypto.Encrypt("hellocemo")

	print(encrypted + ": this is the encrypted text")
	decrypted := crypto.Decrypt(encrypted)

	print(decrypted + ": this is the decrypted text")

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/encrypt", encrypt)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
func login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	if username == "admin" && password == "admin" {
		_, err := w.Write([]byte("success"))
		if err != nil {
			return
		}
	} else {
		_, err := w.Write([]byte("failed"))
		if err != nil {
			return
		}
	}
}
func hello(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello World"))
	if err != nil {
		return
	}
}
func encrypt(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")

	crypt := crypto.Encrypt(text)

	_, _ = w.Write([]byte(crypt))
}
