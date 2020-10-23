package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/abdulcrimping/movies/models"
)

var movies []models.Movie = []models.Movie{
	models.Movie{
		ID:    1,
		Title: "Menculik si udin",
		Genre: "Comdey",
	},
	models.Movie{
		ID:    2,
		Title: "Bapakku Polwan",
		Genre: "parody",
	},
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(movies)
}

func main() {
	a := App{}
	a.InitializeDB(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	http.HandleFunc("/", cors(getMovies))
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func cors(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		next(w, r)
	})
}
