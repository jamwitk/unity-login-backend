package controllers

import (
	"backend/models"
	"backend/utils"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
)

var db = utils.ConnectToDB()

func CreateNewUser(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		response := map[string]interface{}{"status": "ERROR", "message": "There is error when binding"}
		fmt.Println(response)
	}

	pass, dErr := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if dErr != nil {
		fmt.Println("Error ")
		err := json.NewEncoder(w).Encode(dErr)
		if err != nil {
			return
		}
	}

	user.Password = string(pass)
	createdUser := db.Create(user)
	if createdUser.Error != nil {
		fmt.Println("Error Code: R02-", createdUser)
	}

	userModelError := json.NewEncoder(w).Encode(user)
	if userModelError != nil {
		fmt.Println("Error Code: R02-", userModelError)
		return
	}

	_, userErr := w.Write([]byte("User created successfully"))
	if userErr != nil {
		fmt.Println("Error Code: R04-", userErr)
		return
	}
}
func CreateToken(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query().Get("s")
	//password := r.URL.Query().Get("password")
	//sessionToken := r.URL.Query().Get("tokenId")
	username := strings.Split(data, "~")
	fmt.Println(username[0], username[1], username[2])
}

//2. login a user
//3.
