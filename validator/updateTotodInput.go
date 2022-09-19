package validator

type UpdateTodoInput struct {
	Title     string `json:"title,omitempty"`
	Completed bool   `json:"completed,omitempty"`
}
