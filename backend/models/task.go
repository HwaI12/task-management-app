package models

// Task represents a task in the system
type Task struct {
	ID       int    `json:"id"`
	UserID   string `json:"user_id"`
	Title    string `json:"title"`
	Purpose  string `json:"purpose"`
	Deadline string `json:"deadline"`
	Priority string `json:"priority"`
	Status   string `json:"status"`
	Steps    string `json:"steps"`
	Memo     string `json:"memo"`
	Remarks  string `json:"remarks"`
}
