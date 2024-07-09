package models

// Task represents a task in the system
type Task struct {
	ID          int    `json:"id"`
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	Deadline    string `json:"deadline"`
	Priority    string `json:"priority"`
	Status      string `json:"status"`
	Purpose     string `json:"purpose"`
	Description string `json:"description"`
	Steps       string `json:"steps"`
	Memo        string `json:"memo"`
	Remarks     string `json:"remarks"`
}
