package handlers

import (
	"encoding/json"
	"github.com/NureTymofiienkoSnizhana/arkpz-pzpi-22-9-tymofiienko-snizhana/Pract1/arkpz-pzpi-22-9-tymofiienko-snizhana-task2/src/api/requests"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
)

func Auth(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewAuth(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	usersDB := MongoDB(r).Users()
	user, err := usersDB.FindByEmail(req.Email)
	if err != nil {
		if strings.Contains(err.Error(), "no documents in result") {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "User not found",
			})
			return
		}
		http.Error(w, "Failed to authenticate user", http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User authenticated successfully",
		"userID":  user.ID.Hex(),
		"role":    user.Role,
	})
}
