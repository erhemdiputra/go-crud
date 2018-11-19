package user

import (
	"context"

	"github.com/erhemdiputra/go-crud/models"
)

type IUserUsecase interface {
	GetList(ctx context.Context) ([]models.User, error)
	Add(context.Context, models.User) (int64, error)
	GetByID(context.Context, int64) (models.User, error)
	Update(context.Context, models.User) (int64, error)
	Delete(context.Context, int64) (int64, error)
}
