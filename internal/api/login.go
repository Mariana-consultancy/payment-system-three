package api

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
// log in system 
var users = map[string]string{
	"jeffery":   hashPassword("password123"),
	"wayne":     hashPassword("qwerty456"),
	"Matthew": hashPassword("zxcvbn789"),
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
// check creditional of user name and password 
	if checkCredentials(user.Username, user.Password) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Login successful"))
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}

func checkCredentials(username, password string) bool {
	hashedPassword := hashPassword(password)
	storedPassword, exists := users[username]
	return exists && storedPassword == hashedPassword
}

func hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

