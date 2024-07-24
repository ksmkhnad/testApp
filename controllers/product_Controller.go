package controllers

import (
	"app/models"
	"app/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productName := strings.ToLower(params["name"])
	var product models.Product
	if utils.DB.Where("LOWER(name) = ?", productName).First(&product).RecordNotFound() {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(product)
}
