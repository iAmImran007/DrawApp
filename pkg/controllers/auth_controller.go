package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/iAmImran007/draw-app-js-go/pkg/config"
	"github.com/iAmImran007/draw-app-js-go/pkg/models"
	"github.com/iAmImran007/draw-app-js-go/pkg/utils"
)

//login user handels user and authanticate and token ganeret
func LoginUser(w http.ResponseWriter, r *http.Request) {
  var user models.User
	var foundUser models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Envalied inputs", http.StatusBadRequest)
		return
	}
	result := config.DB.Where("email = ?", user.Email).First(&foundUser)
	if result.Error != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	//check user 
	if utils.CompairPassword(foundUser.Password, user.Password) {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	//ganeret jwt token 
	utils.GaneretTokenResponse(w, foundUser.ID)

}