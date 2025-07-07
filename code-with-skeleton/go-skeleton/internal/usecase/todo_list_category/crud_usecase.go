package todo_list_category_usecase

import (
	"context"
	"time"

	"github.com/dikadittya/learn-golang/tree/code-with-skeleton/go-skeleton/internal/repository/mysql"
	myentity "github.com/dikadittya/learn-golang/tree/code-with-skeleton/go-skeleton/internal/repository/mysql/entity"
	"github.com/dikadittya/learn-golang/tree/code-with-skeleton/go-skeleton/internal/usecase/todo_list_category/entity"
)

// TODO
// Create
// GetByID
// Update
// DeleteByID
// Get All

type CrudTodoListCategory struct {
	TodoListCategoryRepo mysql.ITodoListCategoryRepository
}

func NewCrudTodoListCategory(
	TodoListCategoryRepo mysql.ITodoListCategoryRepository,
) *CrudTodoListCategory {
	return &CrudTodoListCategory{TodoListCategoryRepo}
}

type ICrudTodoListCategory interface {
	Create(ctx context.Context, req entity.TodoListCategoryReq) error
	GetAll(ctx context.Context) (res []*entity.TodoListCategoryResponse, err error)
}

func (t *CrudTodoListCategory) GetAll(ctx context.Context) (res []*entity.TodoListCategoryResponse, err error) {
	// funcName := "CrudTodoListCategory.GetAll"

	result, err := t.TodoListCategoryRepo.GetAll(ctx)
	if err != nil {
		// helper.LogError("TodoListCategoryRepo.GetByAll", funcName, err, captureFieldError, "")

		return nil, err
	}

	for _, v := range result {
		res = append(res, &entity.TodoListCategoryResponse{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
		})
	}

	return res, nil
}
func (u *CrudTodoListCategory) Create(ctx context.Context, req entity.TodoListCategoryReq) error {
	// TODO
	// Ambil request data
	data := &myentity.TodoListCategory{
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   time.Now(),
	}

	// Insert ke Table
	err := u.TodoListCategoryRepo.Create(ctx, nil, data, false)
	if err != nil {
		return err
	}

	return nil
}
