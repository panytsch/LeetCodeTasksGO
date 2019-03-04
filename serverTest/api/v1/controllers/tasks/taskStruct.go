package tasks

type Task struct {
	IsDone      bool   `json:"isDone"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TaskPostRequest struct {
	Id   int  `json:"id"`
	Task Task `json:"task"`
}
