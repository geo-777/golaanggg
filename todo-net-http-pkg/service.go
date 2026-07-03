package main

// handles todo storing and manipulation

type Priority string

const (
	Low    Priority = "low"
	Medium Priority = "medium"
	High   Priority = "high"
)

type Todo struct {
	Id        int      `json:"int,omitempty"`
	Task      string   `json:"task"`
	Completed bool     `json:"completed"`
	Priority  Priority `json:"priority"`
}

type TodoPatchDto struct {
	Task      *string   `json:"task,omitempty"`
	Completed *bool     `json:"completed,omitempty"`
	Priority  *Priority `json:"priority,omitempty"`
}

var Todos []Todo //todos declaration
var nextId int = 1

func postTodo(todoItem Todo) Todo {
	todoItem.Id = nextId
	nextId++
	Todos = append(Todos, todoItem)

	return todoItem
}
func getTodos() []Todo {
	//we can't send Todos directly here (Todos is a pointer and the caller may make breaking changes)
	copied := make([]Todo, len(Todos))
	copy(copied, Todos)
	return copied
}

func patchTodo(id int, patch TodoPatchDto) bool {
	for i := range Todos {
		if Todos[i].Id == id {
			if patch.Task != nil {
				Todos[i].Task = *patch.Task
			}
			if patch.Completed != nil {
				Todos[i].Completed = *patch.Completed
			}
			if patch.Priority != nil {
				Todos[i].Priority = *patch.Priority
			}

			return true
		}
	}
	return false
}

func deleteTodo(id int) bool {
	for i, todo := range Todos {
		if todo.Id == id {
			Todos = append(Todos[:i], Todos[i+1:]...)
			return true
		}
	}

	return false
}
