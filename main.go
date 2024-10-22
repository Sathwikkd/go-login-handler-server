package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type loginRequest struct {
	UserName string `json:"username"` 
	PassWord string `json:"password"`
}

type loginResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

var validUserName = "admin"
var validPassword = "admin@123"

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid request", http.StatusMethodNotAllowed)
		return
	}

	var loginReq loginRequest
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		http.Error(w, "error parsing request body", http.StatusBadRequest)
		return
	}

	var response loginResponse
	if loginReq.UserName == validUserName && loginReq.PassWord == validPassword {
		response = loginResponse{
			Message: "login successful",
			Success: true,
		}
	} else {
		response = loginResponse{
			Message: "login failed: invalid username or password", // Improved error message
			Success: false,
		}
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/login", loginHandler)
	fmt.Println("server Listening at port no:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
