package repositories

import (
	"context"
	"github.com/bonus2k/lab8/internal/models"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"sync"
	"time"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, id uuid.UUID) (*models.User, error)
	GetAllUsers(ctx context.Context) ([]models.User, error)
}

var (
	once sync.Once
	repo UserRepository
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (u *UserRepositoryImpl) CreateUser(ctx context.Context, user *models.User) error {
	result := u.db.WithContext(ctx).Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *UserRepositoryImpl) GetUser(ctx context.Context, id uuid.UUID) (*models.User, error) {
	var user models.User
	result := u.db.WithContext(ctx).First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (u *UserRepositoryImpl) GetAllUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User
	result := u.db.WithContext(ctx).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func Init(dsn string) (UserRepository, error) {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	once.Do(func() {
		db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
		if err != nil {
			return
		}
		repo = &UserRepositoryImpl{db: db}
		err = db.WithContext(ctx).AutoMigrate(&models.User{})
		if err != nil {
			repo = nil
		}
	})
	return repo, err
}
