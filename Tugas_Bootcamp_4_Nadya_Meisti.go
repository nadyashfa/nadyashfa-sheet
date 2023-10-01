package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
)

type Input struct {
	Sisi  float64 `json:"sisi"`
	Jenis string  `json:"jenis"`
}

type Output struct {
	Keliling float64 `json:"keliling"`
	Luas     float64 `json:"luas"`
}

func main() {
	http.HandleFunc("/hitung", Hitung)
	fmt.Println("Server is listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Hitung(w http.ResponseWriter, r *http.Request) {
	var input Input
	var output Output

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch input.Jenis {
	case "lingkaran":
		output.Keliling = 2 * math.Pi * input.Sisi
		output.Luas = math.Pi * math.Pow(input.Sisi, 2)
	case "segitiga":
		output.Keliling = 3 * input.Sisi
		output.Luas = (math.Pow(input.Sisi, 2) * math.Sqrt(3)) / 4
		output.Keliling = 4 * input.Sisi
		output.Luas = math.Pow(input.Sisi, 2)
	default:
		http.Error(w, "Jenis bentuk tidak valid", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
