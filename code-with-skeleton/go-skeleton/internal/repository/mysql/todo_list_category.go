package mysql

import (
	"context"

	"github.com/dikadittya/learn-golang/tree/code-with-skeleton/go-skeleton/config"
	"github.com/dikadittya/learn-golang/tree/code-with-skeleton/go-skeleton/internal/helper"
	"github.com/dikadittya/learn-golang/tree/code-with-skeleton/go-skeleton/internal/repository/mysql/entity"
)

type ITodoListCategoryRepository interface {
	TrxSupportRepo
	Create(ctx context.Context, dbTrx TrxObj, params *entity.TodoListCategory, nonZeroVal bool) error
}

type TodoListCategoryRepository struct {
	GormTrxSupport
}

func NewTodoListCategoryRepository(mysql *config.Mysql) *TodoListCategoryRepository {
	return &TodoListCategoryRepository{GormTrxSupport{db: mysql.DB}}
}

func (r *TodoListCategoryRepository) Create(ctx context.Context, dbTrx TrxObj, params *entity.TodoListCategory, nonZeroVal bool) error {
	// funcName := "TodoListCategoryRepository.Create"

	// fmt.Println(params)
	// fmt.Println("dataXXXXXXXXXXXXXXXXX")
	// if err := helper.CheckDeadline(ctx); err != nil {
	// 	return errwrap.Wrap(err, funcName)
	// }

	cols := helper.NonZeroCols(params, nonZeroVal)
	return r.Trx(dbTrx).Select(cols).Create(&params).Error
}
