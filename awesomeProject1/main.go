package main

import (
	"awesomeProject1/crypto"
	"fmt"
	"net/http"
)

func main() {

	encrypted, errs := crypto.Encrypt("hellocemo", "key")
	if errs != nil {
		return
	}
	print(encrypted + ": this is the encrypted text")
	decrypted, errs := crypto.Decrypt(encrypted, "key")
	if errs != nil {
		fmt.Printf("error: %v", errs)
		return
	}
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
	key := r.FormValue("key")

	crypt, err := crypto.Encrypt(text, key)
	if err != nil {
		return
	}

	_, err = w.Write([]byte(crypt))
	if err != nil {
		return
	}
}
