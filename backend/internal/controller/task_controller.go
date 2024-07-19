package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/HwaI12/task-management-app/backend/internal/model"
	"github.com/HwaI12/task-management-app/backend/internal/service"
	"github.com/gorilla/mux"
)

type TaskController struct {
	service *service.TaskService
}

func NewTaskController(service *service.TaskService) *TaskController {
	return &TaskController{service: service}
}

func (c *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		log.Printf("リクエストの解析に失敗しました: %v", err)
		http.Error(w, "リクエスト形式が正しくありません", http.StatusBadRequest)
		return
	}

	taskID, err := c.service.CreateTask(task)
	if err != nil {
		log.Printf("タスクの保存に失敗しました: %v", err)
		http.Error(w, "タスクの登録に失敗しました", http.StatusInternalServerError)
		return
	}

	task.ID = int(taskID)

	log.Println("タスクが正常に追加されました")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (c *TaskController) GetTasks(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "ユーザーIDが指定されていません", http.StatusBadRequest)
		return
	}

	tasks, err := c.service.GetTasksByUserID(userID)
	if err != nil {
		log.Printf("タスクの取得に失敗しました: %v", err)
		http.Error(w, "タスクの取得に失敗しました", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		log.Printf("タスクのエンコードに失敗しました: %v", err)
		http.Error(w, "タスクの取得に失敗しました", http.StatusInternalServerError)
	}
}

func (c *TaskController) GetUserTasks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID := vars["task_id"]
	if taskID == "" {
		http.Error(w, "タスクIDが指定されていません", http.StatusBadRequest)
		return
	}

	task, err := c.service.GetTaskByID(taskID)
	if err != nil {
		log.Printf("タスクの取得に失敗しました: %v", err)
		http.Error(w, "タスクの取得に失敗しました", http.StatusInternalServerError)
		return
	}

	if task.ID == 0 {
		http.Error(w, "タスクが見つかりません", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(task); err != nil {
		log.Printf("タスクのエンコードに失敗しました: %v", err)
		http.Error(w, "タスクのエンコードに失敗しました", http.StatusInternalServerError)
	}
}
