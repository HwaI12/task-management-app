package models

type User struct {
	ID           int    `json:"id"`
	User_id      string `json:"user_id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	Password     string `json:"password" gorm:"-"`
	CreatedAt    string `json:"created_at"`
}

type RegisterRequest struct {
	UserID       string `json:"user_id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

// type SubSection struct {
// 	ID      int    `json:"id"`
// 	TaskID  int    `json:"task_id"`
// 	Title   string `json:"title"`
// 	Content string `json:"content"`
// }

// type PersonalArtifact struct {
// 	ID        int       `json:"id"`
// 	TaskID    int       `json:"task_id"`
// 	UserID    int       `json:"user_id"`
// 	FilePath  string    `json:"file_path"`
// 	CreatedAt time.Time `json:"created_at"`
// }

// type SharedArtifact struct {
// 	ID        int       `json:"id"`
// 	UserID    int       `json:"user_id"`
// 	Title     string    `json:"title"`
// 	Content   string    `json:"content"`
// 	FilePath  string    `json:"file_path"`
// 	CreatedAt time.Time `json:"created_at"`
// }

// type Progress struct {
// 	ID              int `json:"id"`
// 	TaskID          int `json:"task_id"`
// 	ProgressPercent int `json:"progress_percent"`
// }

// type Comment struct {
// 	ID               int       `json:"id"`
// 	SharedArtifactID int       `json:"shared_artifact_id"`
// 	UserID           int       `json:"user_id"`
// 	CommentText      string    `json:"comment_text"`
// 	CreatedAt        time.Time `json:"created_at"`
// }
