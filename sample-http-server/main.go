package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// trying out Handle interface
type HomeHandler struct{}

func (h HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	name := r.URL.Query().Get("name")
	fmt.Println("Got /")
	fmt.Fprintln(w, "/ page has been loaded for", name)
}

func getHello(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("got /hello request")
	io.WriteString(w, "Hello! HTTP\n")
}

func echoFunc(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Failed to read body", http.StatusBadRequest)
		return
	}
	fmt.Println(body) // prints bytes
	fmt.Fprintln(w, string(body))
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	fmt.Println(user.Age, user.Name)

	fmt.Fprintln(w, "User Created")
}

func getUser(w http.ResponseWriter, _ *http.Request) {
	var user User = User{
		Name: "John",
		Age:  20,
	}
	data, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

//logger (middle-ware)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		fmt.Println("Request recieved : ", r.Method, r.URL.Path, time.Since(start))

	})
}

func main() {

	mux := http.NewServeMux()
	mux.Handle("/", HomeHandler{})
	mux.HandleFunc("/hello", getHello)
	mux.HandleFunc("/echo", echoFunc)

	mux.HandleFunc("POST /users", createUser)
	mux.HandleFunc("GET /users", getUser)

	handler := loggingMiddleware(mux)

	err := http.ListenAndServe(":3333", handler)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server was closed.")
	} else if err != nil {
		fmt.Println("Error starting server", err)
		os.Exit(1)
	}
}
