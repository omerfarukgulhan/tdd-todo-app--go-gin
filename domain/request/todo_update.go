package request

type TodoUpdate struct {
	UserId      int    `json:"userId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"isCompleted"`
}
