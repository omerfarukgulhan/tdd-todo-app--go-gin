package request

type TodoCreate struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
