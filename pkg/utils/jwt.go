package utils

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/iAmImran007/draw-app-js-go/pkg/config"
)

//create a jwt token
func GenerateToken(userID uint) (string, error) {
	jwtSecret := []byte(config.GetEnv("JWT_SECRET_KEY"))
	if len(jwtSecret) == 0 {
		jwtSecret = []byte("default_secret_key")
	}

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}


func GaneretTokenResponse(w http.ResponseWriter, userID uint) {
   token, err := GenerateToken(userID)
	 if err != nil {
		http.Error(w, `{"errro": "Feild to ganaret token"}`, http.StatusInternalServerError)
		return
	 }
	 w.Header().Set("Content-Type", "application/json")
	 w.WriteHeader(http.StatusOK)
	 json.NewEncoder(w).Encode(map[string]string{"token": token})
}