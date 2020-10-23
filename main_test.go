package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/abdulcrimping/movies/models"
)

var a App

func testMain(m *testing.M) {
	a.InitializeDB(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	t := m.Run()
	os.Exit(t)
}

func TestMovies(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	getMovies(res, req)

	var m []models.Movie
	json.Unmarshal(res.Body.Bytes(), &m)
	if res.Code != http.StatusOK {
		t.Errorf("Expected status code `%d', got '%d' instead", http.StatusOK, res.Code)
	}
	if m[0].Title != movies[0].Title {
		t.Errorf("Expected '%s', got '%s' instead", movies[0].Title, m[0].Title)
	}
}

const createTable = `
	CREATE TABLE IF NOT EXISTS movies(
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		genre TEXT NOT NULL
	)
`

func tableMustExist(t *testing.T) {
	if _, err := a.DB.Exec(createTable); err != nil {
		log.Fatal(err)
	}
}
