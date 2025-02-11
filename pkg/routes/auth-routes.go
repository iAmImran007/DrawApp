package routes

import (
	"encoding/json"
	"net/http"

	"github.com/iAmImran007/draw-app-js-go/pkg/config"
	"github.com/iAmImran007/draw-app-js-go/pkg/models"
	"github.com/iAmImran007/draw-app-js-go/pkg/utils"
)


func RegisterUser(w http.ResponseWriter, r *http.Request) {
   var user models.User
	 err := json.NewDecoder(r.Body).Decode(&user)
	 if err != nil {
		http.Error(w, "Envalid input", http.StatusBadRequest)
		return
	 }
   
	 //hashed password
	 hashedPassword, err := utils.HashPassword(user.Password)
	 if err != nil {
		http.Error(w, "Feild to hashed password", http.StatusInternalServerError)
		return
	 }
	 user.Password = hashedPassword

   
	 //save user to db
   result := config.DB.Create(&user)
   if result.Error != nil {
		http.Error(w, "feild to store the user details", http.StatusInternalServerError)
		return
	 }

	 //gneret jwt tokens 
	 utils.GaneretTokenResponse(w, user.ID)
}
