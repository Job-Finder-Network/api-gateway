package service

import (
	"context"
	"github.com/Job-Finder-Network/api-gateway/entity"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user entity.User) error
	SelectUserByEmail(ctx context.Context, email string) (entity.User, error)
}
type UserService interface {
	CreateUser(ctx context.Context, email string, password string,role string) (string, error)
	AuthenticateUser(ctx context.Context, email string, password string) (bool, error)
	FindUserByEmail(ctx context.Context, email string) (entity.User, error)
}

type userService struct {
	userrepository UserRepository
	logger         log.Logger
}

func NewUserService(rep UserRepository, logger log.Logger) UserService {
	return &userService{
		userrepository: rep,
		logger:         logger,
	}
}

func (s userService) CreateUser(ctx context.Context, email string, password string,role string) (string, error) {
	logger := log.With(s.logger, "method", "CreateUser")

	uuid, _ := uuid.NewV4()
	id := uuid.String()
	hashPassword, err := hashPassword(password)
	if err != nil {
		return "", err
	}
	user := entity.User{
		ID:       id,
		Email:    email,
		Password: hashPassword,
		Role: role,
	}

	if err := s.userrepository.CreateUser(ctx, user); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("create user", id)

	return "Success", nil
}

func (s userService) AuthenticateUser(ctx context.Context, email string, password string) (bool, error) {
	user, err := s.userrepository.SelectUserByEmail(ctx, email)
	if err != nil {
		return false, err
	}
	if user.Email == email && (bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))) == nil {
		return true, nil
	}
	return false, nil
}

func (s userService) FindUserByEmail(ctx context.Context, email string) (entity.User, error) {
	user, err := s.userrepository.SelectUserByEmail(ctx, email)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), err
}
