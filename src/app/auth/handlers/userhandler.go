package handlers

import (
	"fmt"
	"net/http"

	AuthModels "github.com/faisd405/go-restapi-chi/src/app/auth/models"
	AuthUtils "github.com/faisd405/go-restapi-chi/src/app/auth/utils"
	database "github.com/faisd405/go-restapi-chi/src/config"
	"github.com/faisd405/go-restapi-chi/src/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	user := AuthModels.User{}
	userBody := AuthModels.User{}
	utils.ReadJson(r, &userBody)

	// Find the user in the database
	database.DB.Where("username = ?", userBody.Username).First(&user)

	// Check if the password is correct
	if !AuthUtils.CheckPassword(userBody.Password, user.Password) {
		utils.WriteJson(w, map[string]string{"error": "invalid credentials"})
		return
	}

	// Create a new token for the user
	token, err := utils.NewToken(user.ID)
	if err != nil {
		utils.WriteJson(w, map[string]string{"error": "error creating token"})
		return
	}
	fmt.Println(user.ID)

	// Return the token
	utils.WriteJson(w, map[string]string{"token": token})

}

func Register(w http.ResponseWriter, r *http.Request) {
	user := AuthModels.User{}

	utils.ReadJson(r, &user)

	// Check if the username is already taken
	checkUsername := database.DB.Where("username = ?", user.Username).First(&user)
	if checkUsername.RowsAffected != 0 {
		utils.WriteJson(w, map[string]string{"error": "username already taken"})
		return
	}

	// Hash the password
	hash, err := AuthUtils.HashPassword(user.Password)
	if err != nil {
		utils.WriteJson(w, map[string]string{"error": "error hashing password"})
		return
	}

	// Set the password to the hashed password
	user.Password = hash

	// Save the user in the database
	database.DB.Create(&user)

	// Return the user
	utils.WriteJson(w, user)
}

func User(w http.ResponseWriter, r *http.Request) {
	// Check if the token is valid
	userJWT, err := utils.VerifyToken(r)
	if err != nil {
		utils.WriteJson(w, map[string]string{"error": "invalid token"})
		return
	}

	// Find the user in the database
	user := AuthModels.User{}
	database.DB.Where("id = ?", userJWT["id"]).First(&user)

	// Remove the password from the user
	user.Password = ""

	// Return the user
	utils.WriteJson(w, user)
}
