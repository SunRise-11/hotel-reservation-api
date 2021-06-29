package repositories

import (
	"gorm.io/gorm"
	"hotel-reservation/internal/commons"
	"hotel-reservation/internal/dto"
	"hotel-reservation/internal/models"
	"hotel-reservation/pkg/application_loger"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) Create(user *models.User) (*models.User, error) {

	if tx := r.DB.Create(&user); tx.Error != nil {
		application_loger.LogError(tx.Error)
		return nil, tx.Error
	}

	return user, nil
}

func (r *UserRepository) Update(user *models.User) (*models.User, error) {

	if tx := r.DB.Updates(&user); tx.Error != nil {
		application_loger.LogError(tx.Error)
		return nil, tx.Error
	}

	return user, nil
}

func (r *UserRepository) Find(id uint64) (*models.User, error) {

	model := models.User{}
	if tx := r.DB.Where("id=?", id).Find(&model); tx.Error != nil {
		application_loger.LogError(tx.Error)
		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *UserRepository) FindAll(input *dto.PaginationInput) (*commons.PaginatedList, error) {

	return finAll(&models.User{}, r.DB, input)
}

func (r *UserRepository) FindBySymbol(symbol string) (*models.User, error) {
	model := models.User{}
	if tx := r.DB.Where("symbol=?", symbol).Find(&model); tx.Error != nil {
		application_loger.LogError(tx.Error)
		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}
