package request

type TodoUpdate struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"isCompleted"`
}
