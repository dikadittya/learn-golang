package handler

import (
	"net/http"

	"github.com/dikadittya/learn-golang/tree/code-with-skeleton/go-skeleton/internal/parser"
	"github.com/dikadittya/learn-golang/tree/code-with-skeleton/go-skeleton/internal/presenter/json"
	todo_list_category_usecase "github.com/dikadittya/learn-golang/tree/code-with-skeleton/go-skeleton/internal/usecase/todo_list_category"
	"github.com/dikadittya/learn-golang/tree/code-with-skeleton/go-skeleton/internal/usecase/todo_list_category/entity"

	fiber "github.com/gofiber/fiber/v2"
)

type TodoListCategoryHandler struct {
	parser                      parser.Parser
	presenter                   json.JsonPresenter
	CrudTodoListCategoryUsecase todo_list_category_usecase.ICrudTodoListCategory
}

func NewTodoListCategoryHandler(
	parser parser.Parser,
	presenter json.JsonPresenter,
	CrudTodoListCategoryUsecase todo_list_category_usecase.ICrudTodoListCategory,
) *TodoListCategoryHandler {
	return &TodoListCategoryHandler{parser, presenter, CrudTodoListCategoryUsecase}
}

func (w *TodoListCategoryHandler) Register(app fiber.Router) {
	app.Post("/todo-lists-category", w.Create)
	app.Get("/todo-lists-category", w.GetAll)
	app.Get("/todo-lists-category/:id", w.GetByID)
	app.Put("/todo-lists-category/:id", w.Update)
	app.Delete("/todo-lists-category/:id", w.Delete)
}
func (w *TodoListCategoryHandler) Create(c *fiber.Ctx) error {

	var req entity.TodoListCategoryReq
	err := w.parser.ParserBodyRequest(c, &req)

	if err != nil {
		return w.presenter.BuildError(c, err)
	}
	err = w.CrudTodoListCategoryUsecase.Create(c.Context(), req)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	return w.presenter.BuildSuccess(c, nil, "Success", http.StatusOK)
}

func (w *TodoListCategoryHandler) GetAll(c *fiber.Ctx) error {

	data, err := w.CrudTodoListCategoryUsecase.GetAll(c.Context())
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	return w.presenter.BuildSuccess(c, data, "Success", http.StatusOK)
}
func (w *TodoListCategoryHandler) GetByID(c *fiber.Ctx) error {
	id, err := w.parser.ParserIntIDFromPathParams(c)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	data, err := w.CrudTodoListCategoryUsecase.GetByID(c.Context(), id)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	return w.presenter.BuildSuccess(c, data, "Success", http.StatusOK)
}
func (w *TodoListCategoryHandler) Update(c *fiber.Ctx) error {
	var req entity.TodoListCategoryReq
	err := w.parser.ParserBodyWithIntIDPathParams(c, &req)

	if err != nil {
		return w.presenter.BuildError(c, err)
	}
	err = w.CrudTodoListCategoryUsecase.UpdateByID(c.Context(), req)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	return w.presenter.BuildSuccess(c, nil, "Success", http.StatusOK)
}
func (w *TodoListCategoryHandler) Delete(c *fiber.Ctx) error {
	id, err := w.parser.ParserIntIDFromPathParams(c)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	err = w.CrudTodoListCategoryUsecase.DeleteByID(c.Context(), id)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	return w.presenter.BuildSuccess(c, nil, "Success", http.StatusOK)
}
