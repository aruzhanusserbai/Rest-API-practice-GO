package handlers

import (
	"bookstore/models"
	"encoding/json"
	"net/http"
)

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var authorList []models.Author
	for _, author := range models.Authors {
		authorList = append(authorList, author)
	}
	json.NewEncoder(w).Encode(authorList)
}
func AddAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var author models.Author

	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if author.Name == "" {
		http.Error(w, "Author's name is required", http.StatusBadRequest)
		return
	}

	author.ID = models.NextAuthorID
	models.NextAuthorID++
	models.Authors[author.ID] = author

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(author)
}
