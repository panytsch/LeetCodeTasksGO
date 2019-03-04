package tasks

type Task struct {
	IsDone      bool   `json:"isDone"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
