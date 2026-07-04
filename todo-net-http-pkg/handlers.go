package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func rootHandler(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Hello, welcome to todo app"))
}

func getTodosHandler(w http.ResponseWriter, _ *http.Request) {
	data, err := json.Marshal(getTodos())
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func postTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	err := json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	todo = postTodo(todo)
	fmt.Fprintln(w, "Todo Created")

}

func deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
	}
	deleted := deleteTodo(id)
	if !deleted {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{
		"message": "Todo deleted successfulyy",
	}
	json.NewEncoder(w).Encode(response)
}

func patchTodoHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
	}
	var todo TodoPatchDto
	error := json.NewDecoder(r.Body).Decode(&todo)

	if error != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	patched := patchTodo(id, todo)

	if !patched {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{
		"message": "Todo patched successfulyy",
	}
	json.NewEncoder(w).Encode(response)

}
