package validator

type CreateTodoInput struct {
	Title     string `json:"title" binding:"required"`
	Completed bool   `json:"completed"`
	UserId    uint   `json:"user_id" binding:""`
}
