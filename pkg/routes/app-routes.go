
package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iAmImran007/draw-app-js-go/pkg/controllers"
	"github.com/iAmImran007/draw-app-js-go/pkg/middleware"
)

func RegisterApproutes(router *mux.Router) {
	// Public Route: Anyone can access
	router.HandleFunc("/draw/", controllers.GetDraw).Methods("GET")

	// Protected Routes: Require JWT authentication
	router.Handle("/draw/", middleware.JWTMiddleware(http.HandlerFunc(controllers.CreateDraw))).Methods("POST")
	router.Handle("/draw/{drawId}", middleware.JWTMiddleware(http.HandlerFunc(controllers.GetDrawById))).Methods("GET")
	router.Handle("/draw/{drawId}", middleware.JWTMiddleware(http.HandlerFunc(controllers.UpdateDraw))).Methods("PUT")
	router.Handle("/draw/{drawId}", middleware.JWTMiddleware(http.HandlerFunc(controllers.DeleteDraw))).Methods("DELETE")
}
