package request

type CreateTodo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
