package handler

import (
	"encoding/json"
	"github.com/davide-c/rest-sample-b/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
)

func GetProperties(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	property := []model.Property{}
	db.Preload("Asset").
		Find(&property)
	respondJSON(w, http.StatusOK, property)
}

func GetProperty(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	url := vars["url"]

	project := getPropertyOr404(db, url, w, r)
	if project == nil {
		return
	}
	respondJSON(w, http.StatusOK, project)
}

func getPropertyOr404(db *gorm.DB, url string, w http.ResponseWriter, r *http.Request) *model.Property {
	project := model.Property{}
	if err := db.First(&project, model.Property{Url: url}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &project
}

func CreateProperty(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	property := model.Property{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&property); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&property).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, property)
}

func GetAssets(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	asset := []model.Asset{}
	db.Find(&asset)
	respondJSON(w, http.StatusOK, asset)
}

func CreateAsset(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	asset := []model.Asset{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&asset); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&asset).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, asset)
}
