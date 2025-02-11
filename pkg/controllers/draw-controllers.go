package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/iAmImran007/draw-app-js-go/pkg/config"
	"github.com/iAmImran007/draw-app-js-go/pkg/middleware"
	"github.com/iAmImran007/draw-app-js-go/pkg/models"
	"github.com/iAmImran007/draw-app-js-go/pkg/utils"
)

// Get All Drawings (Public)
func GetDraw(w http.ResponseWriter, r *http.Request) {
	drawings := models.GetDraw()
	res, _ := json.Marshal(drawings)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Get a Single Drawing (Protected)
func GetDrawById(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserKey).(float64) // JWT stores numeric values as float64

	vars := mux.Vars(r)
	drawId := vars["drawId"]
	ID, err := strconv.ParseInt(drawId, 10, 64)
	if err != nil {
		http.Error(w, "Invalid draw ID", http.StatusBadRequest)
		return
	}

	// Fetch drawing from DB
	drawDetails, err := models.GetDrawById(ID)
	if err != nil {
		http.Error(w, "Drawing not found", http.StatusNotFound)
		return
	}

	// Restrict access: Only allow the owner to view
	if drawDetails.UserID != strconv.Itoa(int(userID)) {
		http.Error(w, "Unauthorized access", http.StatusUnauthorized)
		return
	}

	res, _ := json.Marshal(drawDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Create a Drawing (Public)
func CreateDraw(w http.ResponseWriter, r *http.Request) {
	createDraw := &models.Drawing{}
	utils.ParseBody(r, createDraw)

	// Save to DB
	d := createDraw.CreateDraw()
	res, _ := json.Marshal(d)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

// Delete a Drawing (Protected)
func DeleteDraw(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserKey).(float64)

	vars := mux.Vars(r)
	drawId := vars["drawId"]
	ID, err := strconv.ParseInt(drawId, 10, 64)
	if err != nil {
		http.Error(w, "Invalid draw ID", http.StatusBadRequest)
		return
	}

	// Fetch the drawing
	drawDetails, err := models.GetDrawById(ID)
	if err != nil {
		http.Error(w, "Drawing not found", http.StatusNotFound)
		return
	}

	// Restrict access: Only owner can delete
	if drawDetails.UserID != strconv.Itoa(int(userID)) {
		http.Error(w, "Unauthorized action", http.StatusUnauthorized)
		return
	}

	models.DeleteDraw(ID)
	w.WriteHeader(http.StatusNoContent)
}

// Update a Drawing (Protected)
func UpdateDraw(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserKey).(float64)

	vars := mux.Vars(r)
	drawId := vars["drawId"]
	ID, err := strconv.ParseInt(drawId, 10, 64)
	if err != nil {
		http.Error(w, "Invalid draw ID", http.StatusBadRequest)
		return
	}

	// Fetch the drawing
	drawDetails, err := models.GetDrawById(ID)
	if err != nil {
		http.Error(w, "Drawing not found", http.StatusNotFound)
		return
	}

	// Restrict access: Only owner can update
	if drawDetails.UserID != strconv.Itoa(int(userID)) {
		http.Error(w, "Unauthorized action", http.StatusUnauthorized)
		return
	}

	// Parse updated data
	updatedDraw := &models.Drawing{}
	utils.ParseBody(r, updatedDraw)

	// Update fields
	if updatedDraw.Name != "" {
		drawDetails.Name = updatedDraw.Name
	}
	if updatedDraw.Data != "" {
		drawDetails.Data = updatedDraw.Data
	}

	// Save changes
	config.DB.Save(&drawDetails)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(drawDetails)
}
