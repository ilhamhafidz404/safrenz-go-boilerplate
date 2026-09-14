package service

import (
	"context"
	"errors"
	"time"

	"safrenz-go-boilerplate/internal/dto"
	"safrenz-go-boilerplate/internal/model"
	"safrenz-go-boilerplate/internal/repository"
	"safrenz-go-boilerplate/pkg/platform"
)

type UserService interface {
	GetAllUsers(ctx context.Context) ([]dto.UserResponse, error)
	GetUserByID(ctx context.Context, id uint) (*dto.UserResponse, error)
	CreateUser(ctx context.Context, req dto.CreateUserRequest) (*dto.UserResponse, error)
	UpdateUser(ctx context.Context, id uint, req dto.UpdateUserRequest) (*dto.UserResponse, error)
	DeleteUser(ctx context.Context, id uint) error
}

type userService struct {
	userRepo   repository.UserRepository
	roleClient platform.RoleClient
}

func NewUserService(userRepo repository.UserRepository, roleClient platform.RoleClient) UserService {
	return &userService{
		userRepo:   userRepo,
		roleClient: roleClient,
	}
}

func (s *userService) getRoleMap(ctx context.Context) map[int]string {
	roleMap := make(map[int]string)
	roles, err := s.roleClient.FetchRoles(ctx)
	if err == nil {
		for _, r := range roles {
			roleMap[r.ID] = r.Name
		}
	}
	return roleMap
}

func (s *userService) GetAllUsers(ctx context.Context) ([]dto.UserResponse, error) {
	users, err := s.userRepo.FindAll()
	if err != nil {
		return nil, err
	}

	roleMap := s.getRoleMap(ctx)
	userResponses := make([]dto.UserResponse, 0, len(users))

	for _, user := range users {
		roleName, exists := roleMap[user.Role]
		if !exists {
			roleName = "Unknown Role"
		}

		userResponses = append(userResponses, dto.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Role:      roleName,
			CreatedAt: user.CreatedAt.Format(time.RFC3339),
		})
	}

	return userResponses, nil
}

func (s *userService) GetUserByID(ctx context.Context, id uint) (*dto.UserResponse, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("user tidak ditemukan")
	}

	roleMap := s.getRoleMap(ctx)
	roleName, exists := roleMap[user.Role]
	if !exists {
		roleName = "Unknown Role"
	}

	return &dto.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      roleName,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}, nil
}

func (s *userService) CreateUser(ctx context.Context, req dto.CreateUserRequest) (*dto.UserResponse, error) {
	user := model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	}

	if err := s.userRepo.Create(&user); err != nil {
		return nil, err
	}

	return s.GetUserByID(ctx, user.ID)
}

func (s *userService) UpdateUser(ctx context.Context, id uint, req dto.UpdateUserRequest) (*dto.UserResponse, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("user tidak ditemukan")
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Role != 0 {
		user.Role = req.Role
	}

	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	return s.GetUserByID(ctx, user.ID)
}

func (s *userService) DeleteUser(ctx context.Context, id uint) error {
	_, err := s.userRepo.FindByID(id)
	if err != nil {
		return errors.New("user tidak ditemukan")
	}

	return s.userRepo.Delete(id)
}
