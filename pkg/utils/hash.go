package utils

import "golang.org/x/crypto/bcrypt"

//hashed password
func HashPassword(password string) (string, error) {
  bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}


//compair password check if password is coreceet or not 
func CompairPassword(hashedPassword, plainPassword string) bool {
  err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	 return err == nil
}
