package handlers

import (
	"net/http"
	"strconv"
	"todo-full/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CreateTodoInput struct {
	Title     string `json:"title" binding:"required"`
	Completed bool   `json:"completed"`
}

type UpdateTodoInput struct {
	Title     string `json:"title"`
	Completed *bool  `json:"completed"`
}

func CreateTodoHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userIDInterface, exists := ctx.Get("user_id")

		if !exists {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "user_id not found in gin context."})
		}
		userId := userIDInterface.(string)

		var input CreateTodoInput

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})

			return
		}
		todo, err := repository.CreateTodo(pool, input.Title, input.Completed, userId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, todo)
	}
}

func GetAllTodoHandler(pool *pgxpool.Pool) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		userIDInterface, exists := ctx.Get("user_id")

		if !exists {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "user_id not found in gin context."})
		}
		userId := userIDInterface.(string)
		todo, err := repository.GetAllTodos(pool, userId)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, todo)
	}
}

func UpdateTodoHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userIDInterface, exists := ctx.Get("user_id")

		if !exists {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "user_id not found in gin context."})
		}
		userId := userIDInterface.(string)

		idStr := ctx.Param("id")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Malformed param"})
		}
		var input UpdateTodoInput

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}
		if input.Title == "" && input.Completed == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Atleast one field must be provided to update"})
			return
		}

		var completed bool
		if input.Completed != nil {
			completed = *input.Completed
		}

		todo, err := repository.UpdateTodo(pool, id, input.Title, completed, userId)
		if err != nil {
			if err == pgx.ErrNoRows {
				ctx.JSON(http.StatusNotFound, gin.H{"Error": "Todo not found"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, todo)
	}
}

func DeleteTodoHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userIDInterface, exists := ctx.Get("user_id")

		if !exists {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "user_id not found in gin context."})
		}
		userId := userIDInterface.(string)

		idStr := ctx.Param("id")

		id, err := strconv.Atoi(idStr)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid param "})
			return
		}
		err = repository.DeleteTodo(pool, id, userId)

		if err != nil {
			if err.Error() == "Todo was not found" {
				ctx.JSON(http.StatusNotFound, gin.H{"Error ": err.Error()})

			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "todo deleted successfully"})

	}
}
