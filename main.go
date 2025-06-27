package main

import (
	"Car-viewer/api"
	"Car-viewer/db"
	"log"
	"net/http"
)

func main() {
	db.InitDB("favorites.db")
	api.LoadCars() // <-- This loads your car data from api/data.json

	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/", fs)

	http.HandleFunc("/api/cars", api.GetCarsHandler)
	http.HandleFunc("/api/cars/details", api.GetCarDetailsHandler)
	http.HandleFunc("/api/manufacturers", api.GetManufacturersHandler)
	http.HandleFunc("/api/categories", api.GetCategoriesHandler)
	http.HandleFunc("/api/favorites", api.GetFavoritesHandler)
	http.HandleFunc("/api/like", api.LikeCarHandler)
	http.HandleFunc("/api/unlike", api.UnlikeCarHandler)
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./api/img"))))

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
