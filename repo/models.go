package repo

type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type Task struct {
	Status      bool   `json:"status"`
	ID          int    `json:"id"`
	UserTaskID  int    `json:"user_task_id"`
	UserID      int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Deadline    string `json:"deadline"`
}
