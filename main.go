package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Thrashy190/ForumApiGo/auth"
	"github.com/Thrashy190/ForumApiGo/connection"
	"github.com/Thrashy190/ForumApiGo/models"
	"github.com/joho/godotenv"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("connecting to 3000...")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connection.Dbconnection()
	
	connection.DB.AutoMigrate(models.Users{})

	r := mux.NewRouter();

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Connection successful on port 3000"))
	});

	r.HandleFunc("/signUp", auth.SignUp).Methods("POST");
	r.HandleFunc("/login", auth.Login).Methods("GET");

	http.ListenAndServe(":3000",r);
}