package handlers

import (
	"bookstore/models"
	"encoding/json"
	"net/http"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var categoryList []models.Category

	for _, category := range models.Categories {
		categoryList = append(categoryList, category)
	}

	json.NewEncoder(w).Encode(categoryList)
}
func AddCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var category models.Category

	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if category.Name == "" {
		http.Error(w, "Category's name is required", http.StatusBadRequest)
		return
	}

	category.ID = models.NextCategoryID
	models.NextCategoryID++
	models.Categories[category.ID] = category

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)
}
