package handlers

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"pet_project_etap_1/JWT"
)

type Identity struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var cr Identity
	if err := json.NewDecoder(r.Body).Decode(&cr); err != nil {
		http.Error(w, "404", http.StatusBadRequest)
		return
	}
	if cr.Username != "admin" || cr.Password != "3636" {
		http.Error(w, "Не правильный пароль", http.StatusBadRequest)
		return
	}

	token, err := jwt.GenerateJWT(1, "admin")
	if err != nil {
		http.Error(w, "Токен не сгенерирован", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
