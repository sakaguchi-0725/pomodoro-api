package usecase

import (
	"pomodoro-api/domain"
	"pomodoro-api/repository"
	"pomodoro-api/validator"
)

type ITaskUseCase interface {
	GetAllTasks(userId uint) ([]domain.TaskResponse, error)
	GetTaskById(userId uint, taskId uint) (domain.TaskResponse, error)
	CreateTask(task domain.Task) (domain.TaskResponse, error)
	UpdateTask(task domain.Task, userId uint, taskId uint) (domain.TaskResponse, error)
	DeleteTask(userId uint, taskId uint) error
}

type taskUsecase struct {
	tr repository.ITaskRepository
	tv validator.ITaskValidator
}

func NewTaskUsecase(tr repository.ITaskRepository, tv validator.ITaskValidator) ITaskUseCase {
	return &taskUsecase{tr, tv}
}

func (tu *taskUsecase) GetAllTasks(userId uint) ([]domain.TaskResponse, error) {
	tasks := []domain.Task{}
	if err := tu.tr.GetAllTasks(&tasks, userId); err != nil {
		return nil, err
	}

	resTasks := []domain.TaskResponse{}
	for _, v := range tasks {
		t := domain.TaskResponse{
			ID:        v.ID,
			Title:     v.Title,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resTasks = append(resTasks, t)
	}

	return resTasks, nil
}

func (tu *taskUsecase) GetTaskById(userId uint, taskId uint) (domain.TaskResponse, error) {
	task := domain.Task{}
	if err := tu.tr.GetTaskById(&task, userId, taskId); err != nil {
		return domain.TaskResponse{}, err
	}

	resTask := domain.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}

	return resTask, nil
}

func (tu *taskUsecase) CreateTask(task domain.Task) (domain.TaskResponse, error) {
	if err := tu.tv.TaskValidate(task); err != nil {
		return domain.TaskResponse{}, err
	}

	if err := tu.tr.CreateTask(&task); err != nil {
		return domain.TaskResponse{}, err
	}

	resTask := domain.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}

	return resTask, nil
}

func (tu *taskUsecase) UpdateTask(task domain.Task, userId uint, taskId uint) (domain.TaskResponse, error) {
	if err := tu.tv.TaskValidate(task); err != nil {
		return domain.TaskResponse{}, err
	}

	if err := tu.tr.UpdateTask(&task, userId, taskId); err != nil {
		return domain.TaskResponse{}, err
	}

	resTask := domain.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}

	return resTask, nil
}

func (tu *taskUsecase) DeleteTask(userId uint, taskId uint) error {
	if err := tu.tr.DeleteTask(userId, taskId); err != nil {
		return err
	}

	return nil
}
