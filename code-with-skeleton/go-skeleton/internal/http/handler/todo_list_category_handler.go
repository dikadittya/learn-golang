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
	app.Post("/todo-list-category", w.Create)
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
