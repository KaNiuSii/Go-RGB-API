package main

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"net/http"
	"strconv"
)

type RGBColor struct {
	R string `json:"R"`
	G string `json:"G"`
	B string `json:"B"`
}

func getColor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method is supported", http.StatusMethodNotAllowed)
		return
	}

	var rgb RGBColor

	err := json.NewDecoder(r.Body).Decode(&rgb)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rInt, err := strconv.Atoi(rgb.R)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	gInt, err := strconv.Atoi(rgb.G)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	bInt, err := strconv.Atoi(rgb.B)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	col := color.RGBA{R: uint8(rInt), G: uint8(gInt), B: uint8(bInt), A: 255}

	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			img.Set(x, y, col)
		}
	}

	w.Header().Set("Content-Type", "image/png")

	err = png.Encode(w, img)
	if err != nil {
		http.Error(w, "Failed to create image", http.StatusInternalServerError)
	}
}

func CORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		if r.Method == "OPTIONS" {
			http.Error(w, "No Content", http.StatusNoContent)
			return
		}

		next(w, r)
	}
}

func handleRequest() {
	http.HandleFunc("/getColor", CORS(getColor))
}

func main() {
	handleRequest()

	port := ":8081"
	fmt.Printf(" * Server running at: http://localhost%s * \n", port)

	log.Fatal(http.ListenAndServe(port, nil))
}
