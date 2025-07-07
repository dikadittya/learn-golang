package mysql

import (
	"context"

	"github.com/dikadittya/learn-golang/tree/code-with-skeleton/go-skeleton/config"
	"github.com/dikadittya/learn-golang/tree/code-with-skeleton/go-skeleton/internal/helper"
	"github.com/dikadittya/learn-golang/tree/code-with-skeleton/go-skeleton/internal/repository/mysql/entity"

	apperr "github.com/dikadittya/learn-golang/tree/code-with-skeleton/go-skeleton/error"
	errwrap "github.com/pkg/errors"

	"gorm.io/gorm"
)

type ITodoListCategoryRepository interface {
	TrxSupportRepo
	Create(ctx context.Context, dbTrx TrxObj, params *entity.TodoListCategory, nonZeroVal bool) error
	GetAll(ctx context.Context) (e []*entity.TodoListCategory, err error)
}

type TodoListCategoryRepository struct {
	GormTrxSupport
}

func NewTodoListCategoryRepository(mysql *config.Mysql) *TodoListCategoryRepository {
	return &TodoListCategoryRepository{GormTrxSupport{db: mysql.DB}}
}

func (r *TodoListCategoryRepository) Create(ctx context.Context, dbTrx TrxObj, params *entity.TodoListCategory, nonZeroVal bool) error {
	funcName := "TodoListCategoryRepository.Create"

	if err := helper.CheckDeadline(ctx); err != nil {
		return errwrap.Wrap(err, funcName)
	}

	cols := helper.NonZeroCols(params, nonZeroVal)
	return r.Trx(dbTrx).Select(cols).Create(&params).Error
}

func (r *TodoListCategoryRepository) GetAll(ctx context.Context) (result []*entity.TodoListCategory, err error) {
	funcName := "TodoListCategoryRepository.GetAll"

	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errwrap.Wrap(err, funcName)
	}

	err = r.db.Raw("SELECT * FROM todo_list_categories ").Scan(&result).Error
	if errwrap.Is(err, gorm.ErrRecordNotFound) {
		return nil, apperr.ErrRecordNotFound()
	}

	return result, err
}
