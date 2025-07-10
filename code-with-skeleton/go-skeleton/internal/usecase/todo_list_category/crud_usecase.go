package todo_list_category_usecase

import (
	"context"
	"fmt"
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
	GetByID(ctx context.Context, todoListID int64) (*entity.TodoListCategoryResponse, error)
	GetAll(ctx context.Context) (res []*entity.TodoListCategoryResponse, err error)
	UpdateByID(ctx context.Context, todoListReq entity.TodoListCategoryReq) error
	DeleteByID(ctx context.Context, todoListID int64) error
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

func (t *CrudTodoListCategory) GetByID(ctx context.Context, todoListID int64) (*entity.TodoListCategoryResponse, error) {

	data, err := t.TodoListCategoryRepo.GetByID(ctx, todoListID)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, nil
	}

	return &entity.TodoListCategoryResponse{
		ID:          data.ID,
		Name:        data.Name,
		Description: data.Description,
	}, nil
}

func (t *CrudTodoListCategory) UpdateByID(ctx context.Context, req entity.TodoListCategoryReq) error {
	// funcName := "CrudTodoListUsecase.UpdateByID"
	todoListID := req.ID

	if err := mysql.DBTransaction(t.TodoListCategoryRepo, func(trx mysql.TrxObj) error {
		lockedData, err := t.TodoListCategoryRepo.LockByID(ctx, trx, todoListID)
		if err != nil {
			return err
		}
		if lockedData == nil {
			return fmt.Errorf("DATA IS NOT EXIST")
		}

		if err := t.TodoListCategoryRepo.Update(ctx, trx, lockedData, &myentity.TodoListCategory{
			Name:        req.Name,
			Description: req.Description,
		}); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (t *CrudTodoListCategory) DeleteByID(ctx context.Context, todoListID int64) error {
	// funcName := "CrudTodoListUsecase.DeleteByID"
	// captureFieldError := generalEntity.CaptureFields{
	// 	"todo_list_id": helper.ToString(todoListID),
	// }

	err := t.TodoListCategoryRepo.DeleteByID(ctx, nil, todoListID)
	if err != nil {
		// helper.LogError("todoListRepo.DeleteByID", funcName, err, captureFieldError, "")

		return err
	}

	return nil
}
