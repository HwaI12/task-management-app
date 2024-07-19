package service

import (
	"github.com/HwaI12/task-management-app/backend/internal/model"
	"github.com/HwaI12/task-management-app/backend/internal/repository"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task model.Task) (int64, error) {
	return s.repo.Create(task)
}

func (s *TaskService) GetTasksByUserID(userID string) ([]model.Task, error) {
	return s.repo.GetByUserID(userID)
}

func (s *TaskService) GetTaskByID(taskID string) (model.Task, error) {
	return s.repo.GetByID(taskID)
}
