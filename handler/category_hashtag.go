package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/AlbertusKevin/hashtag-grouper/database"
	"github.com/AlbertusKevin/hashtag-grouper/entity"
	"github.com/gorilla/mux"
)

type CategoryHashtagHandler struct {
}

func Get(w http.ResponseWriter, r *http.Request) {
	var categories []entity.CategoryHashtag
	var response entity.Response

	db := database.Connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		response.Status = "500"
		response.Message = "failed parsing; " + err.Error()

		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	fmt.Println(mux.Vars(r))

	quantity, _ := strconv.Atoi(mux.Vars(r)["quantity"])
	categoryID, _ := strconv.Atoi(mux.Vars(r)["category_id"])

	fmt.Println(quantity)
	fmt.Println(categoryID)

	rows, err := db.Query("SELECT * FROM categories_hashtag")
	if err != nil {
		response.Status = "500"
		response.Message = "failed query db; " + err.Error()

		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	for rows.Next() {
		var category entity.CategoryHashtag

		if err := rows.Scan(&category.ID, &category.Category); err != nil {
			response.Status = "500"
			response.Message = "failed scan data; " + err.Error()

			w.WriteHeader(500)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		}

		categories = append(categories, category)
	}

	response.Status = "200"
	response.Message = "Success"
	response.Data = categories

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
