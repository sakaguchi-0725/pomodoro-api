package repository

import (
	"fmt"
	"pomodoro-api/domain"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITaskRepository interface {
	GetAllTasks(tasks *[]domain.Task, userId uint) error
	GetTaskById(task *domain.Task, userId uint, taskId uint) error
	CreateTask(task *domain.Task) error
	UpdateTask(task *domain.Task, userId uint, taskId uint) error
	DeleteTask(userId uint, taskId uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &taskRepository{db}
}

func (tr *taskRepository) GetAllTasks(tasks *[]domain.Task, userId uint) error {
	err := tr.db.Joins("User").Where("user_id=?", userId).Order("created_at").Find(tasks).Error
	if err != nil {
		return err
	}

	return nil
}

func (tr *taskRepository) GetTaskById(task *domain.Task, userId uint, taskId uint) error {
	err := tr.db.Joins("User").Where("user_id=?", userId).First(task, taskId).Error
	if err != nil {
		return err
	}

	return nil
}

func (tr *taskRepository) CreateTask(task *domain.Task) error {
	if err := tr.db.Create(task).Error; err != nil {
		return err
	}

	return nil
}

func (tr *taskRepository) UpdateTask(task *domain.Task, userId uint, taskId uint) error {
	result := tr.db.Model(task).Clauses(clause.Returning{}).Where("id=? AND user_id=?", taskId, userId).Update("title", task.Title)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}

	return nil
}

func (tr *taskRepository) DeleteTask(userId uint, taskId uint) error {
	result := tr.db.Where("id=? AND user_id=?", taskId, userId).Delete(&domain.Task{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}

	return nil
}
