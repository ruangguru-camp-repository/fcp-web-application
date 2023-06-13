package repository

import (
	"a21hc3NpZ25tZW50/apperror"
	"a21hc3NpZ25tZW50/model"
	"errors"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type TaskRepository interface {
	Store(task *model.Task) error
	Update(id int, task *model.Task) error
	Delete(id int) error
	GetByID(id int) (*model.Task, error)
	GetList() ([]model.Task, error)
	GetTaskCategory(id int) ([]model.TaskCategory, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) *taskRepository {
	return &taskRepository{db}
}

func (t *taskRepository) Store(task *model.Task) error {
	err := t.db.Create(task).Error
	if err != nil {
		return err
	}

	return nil
}

func (t *taskRepository) Update(id int, task *model.Task) error {
	if err := t.db.Where("id = ?", id).Updates(model.Task{
		Title: task.Title,
		Deadline: task.Deadline,
		Priority: task.Priority,
		Status: task.Status,
		CategoryID: task.CategoryID,
		UserID: task.UserID,
	}).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.ForeignKeyViolation {
				return apperror.ErrInvalidUserIdOrCategoryId
			}
			return err
		}
	}
	return nil 
}

func (t *taskRepository) Delete(id int) error {
	if err := t.db.Where("id = ?", id).Delete(&model.Task{}).Error; err != nil {
		return err
	}
	return nil 
}

func (t *taskRepository) GetByID(id int) (*model.Task, error) {
	var task model.Task
	err := t.db.First(&task, id).Error
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (t *taskRepository) GetList() ([]model.Task, error) {
	var listOfTask []model.Task
	if err := t.db.Find(&listOfTask).Error; err != nil {
		return nil, err
	}
	return listOfTask, nil 
}

func (t *taskRepository) GetTaskCategory(id int) ([]model.TaskCategory, error) {
	var listOfTaskCategory []model.TaskCategory
	if err := t.db.Table("tasks").
	Select("tasks.id as id, tasks.title as title, categories.name as category").
	Joins("join categories on tasks.category_id = categories.id").
	Where("tasks.id = ?", id).
	Find(&listOfTaskCategory).Error; err != nil {
		return nil, err
	}
	return listOfTaskCategory, nil 
}
