package request

type TodoCreate struct {
	UserId      int    `json:"userId"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
