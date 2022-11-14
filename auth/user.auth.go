package auth

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/Thrashy190/ForumApiGo/connection"
	"github.com/Thrashy190/ForumApiGo/helpers"
	"github.com/Thrashy190/ForumApiGo/models"
	"github.com/golang-jwt/jwt/v4"
)


func SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.Users

	var errPassword error

	json.NewDecoder(r.Body).Decode(&user)

	user.Password, errPassword = helpers.HashPassword(user.Password)

	if errPassword!= nil {
        w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errPassword.Error()))
	}

	createdUser := connection.DB.Create(&user)
	err := createdUser.Error

	if err!= nil {
        w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&user)
}

func Login(w http.ResponseWriter, r *http.Request){
	var user models.Users;

	json.NewDecoder(r.Body).Decode(&user)

	password := user.Password

	connection.DB.Where("username = ?", user.Username).First(&user)

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
        w.Write([]byte("Incorrect username"))
		return;
	}

	match := helpers.CheckPasswordHash(password, user.Password)
    
	if !match {
		w.WriteHeader(http.StatusUnauthorized)
        w.Write([]byte("Incorrect password"))
		return;
	}

	// Create Token with JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour*24*30).Unix(),
	})

	// Sign token with secret key
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
        w.Write([]byte("Failed to sign token"))
		return;
	}

	// Send back token to user using headers

	



	
	json.NewEncoder(w).Encode(tokenString)

}