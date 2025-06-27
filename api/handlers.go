package api

import (
	"Car-viewer/db"
	"encoding/json"
	"net/http"
	"os"
)

// Simulated car database
var cars []map[string]string

func LoadCars() {
	file, err := os.Open("api/data.json")
	if err != nil {
		panic("Could not open data.json: " + err.Error())
	}
	defer file.Close()
	json.NewDecoder(file).Decode(&cars)
}

func GetCarsHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(cars)
}

func GetCarDetailsHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	for _, car := range cars {
		if car["id"] == id {
			json.NewEncoder(w).Encode(car)
			return
		}
	}
	http.Error(w, "Car not found", http.StatusNotFound)
}

func LikeCarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing car ID", http.StatusBadRequest)
		return
	}
	err := db.AddFavorite(id)
	if err != nil {
		http.Error(w, "Failed to like car", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func UnlikeCarHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing car ID", http.StatusBadRequest)
		return
	}
	err := db.RemoveFavorite(id)
	if err != nil {
		http.Error(w, "Failed to remove favorite", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetFavoritesHandler(w http.ResponseWriter, r *http.Request) {
	ids, err := db.GetFavorites()
	if err != nil {
		http.Error(w, "Failed to get favorites", http.StatusInternalServerError)
		return
	}
	var favs []map[string]string
	for _, id := range ids {
		for _, car := range cars {
			if car["id"] == id {
				favs = append(favs, car)
				break
			}
		}
	}
	json.NewEncoder(w).Encode(favs)
}

func GetManufacturersHandler(w http.ResponseWriter, r *http.Request) {
	seen := map[string]bool{}
	var manufacturers []string
	for _, car := range cars {
		m := car["manufacturer"]
		if !seen[m] {
			manufacturers = append(manufacturers, m)
			seen[m] = true
		}
	}
	json.NewEncoder(w).Encode(manufacturers)
}

func GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	seen := map[string]bool{}
	var categories []string
	for _, car := range cars {
		c := car["category"]
		if !seen[c] {
			categories = append(categories, c)
			seen[c] = true
		}
	}
	json.NewEncoder(w).Encode(categories)
}

func CarImageHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	for _, car := range cars {
		if car["id"] == id {
			http.ServeFile(w, r, car["image"])
			return
		}
	}
	http.Error(w, "Car not found", http.StatusNotFound)
}

func CarImageHTMLHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	for _, car := range cars {
		if car["id"] == id {
			html := `<img src="` + car["image"] + `" alt="` + car["model"] + `" style="max-width:300px;display:block;margin-bottom:1em;">`
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte(html))
			return
		}
	}
	http.Error(w, "Car not found", http.StatusNotFound)
}

// Call this in main.go before starting the server:
// api.LoadCars()
