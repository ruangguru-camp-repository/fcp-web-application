package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(email string) (model.User, error)
	CreateUser(user model.User) (model.User, error)
	GetUserTaskCategory() ([]model.UserTaskCategory, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByEmail(email string) (model.User, error) {
	var user model.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return model.User{}, nil
	}
	return user, nil 
}

func (r *userRepository) CreateUser(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) GetUserTaskCategory() ([]model.UserTaskCategory, error) {
	var listData []model.UserTaskCategory
	 result := r.db.Table("tasks").Select(`users.id as id, users.fullname as fullname, 
	users.email as email, tasks.title as task, tasks.deadline as deadline, 
	tasks.priority as priority, tasks.status as status, categories.name as category`).
	Joins("join users on tasks.user_id = users.id").
	Joins("join categories on tasks.category_id = categories.id").
	Find(&listData)
	
	if result.Error != nil {
		return nil, result.Error
	}
	return listData, nil 
}
