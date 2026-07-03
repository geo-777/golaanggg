//lee's implimentation for reference

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Product represents our data model.
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// In-memory "database"
var products = []Product{
	{ID: 1, Name: "Espresso Machine", Price: 599.99},
	{ID: 2, Name: "Coffee Grinder", Price: 89.50},
}

// getProductsHandler handles GET /products
func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode data directly to the connection writer stream
	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, "Failed to encode products", http.StatusInternalServerError)
	}
}

// getProductByIDHandler handles GET /products/{id}
func getProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	// r.PathValue extracts parameters configured in the multiplexer route (Go 1.22+)
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID format", http.StatusBadRequest)
		return
	}

	for _, p := range products {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	// Product not found
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Product not found"})
}

// createProductHandler handles POST /products
func createProductHandler(w http.ResponseWriter, r *http.Request) {
	var newProduct Product

	// Decode incoming JSON from the request body
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	// Simple validation
	if newProduct.Name == "" || newProduct.Price <= 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(map[string]string{"error": "Name and positive price are required"})
		return
	}

	newProduct.ID = len(products) + 1
	products = append(products, newProduct)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newProduct)
}

func notmain() {
	// Create our custom router (multiplexer)
	mux := http.NewServeMux()

	// Register routes with methods (supported natively in Go 1.22+)
	mux.HandleFunc("GET /products", getProductsHandler)
	mux.HandleFunc("GET /products/{id}", getProductByIDHandler)
	mux.HandleFunc("POST /products", createProductHandler)

	// Configure and boot server
	serverAddr := ":8080"
	log.Printf("Starting server on %s...", serverAddr)
	server := &http.Server{
		Addr: serverAddr,

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
