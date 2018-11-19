package usecase

import (
	"context"

	"github.com/erhemdiputra/go-crud/models"
	"github.com/erhemdiputra/go-crud/user"
)

type UserUsecase struct {
	UserRepository user.IUserRepository
}

func NewUserUsecase(userRepository user.IUserRepository) user.IUserUsecase {
	return &UserUsecase{userRepository}
}

func (u *UserUsecase) GetList(ctx context.Context) ([]models.User, error) {
	userList, err := u.UserRepository.GetList(ctx)
	if err != nil {
		return nil, err
	}
	return userList, nil
}

func (u *UserUsecase) Add(ctx context.Context, user models.User) (int64, error) {
	id, err := u.UserRepository.Add(ctx, user)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (u *UserUsecase) GetByID(ctx context.Context, id int64) (models.User, error) {
	user, err := u.UserRepository.GetByID(ctx, id)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (u *UserUsecase) Update(ctx context.Context, user models.User) (int64, error) {
	id, err := u.UserRepository.Update(ctx, user)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (u *UserUsecase) Delete(ctx context.Context, id int64) (int64, error) {
	id, err := u.UserRepository.Delete(ctx, id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
