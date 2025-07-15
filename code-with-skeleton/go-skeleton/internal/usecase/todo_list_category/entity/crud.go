package entity

type TodoListCategoryReq struct {
	ID          int64  `json:"id,omitempty" swaggerignore:"true"`
	Name        string `json:"name" validate:"required" name:"Nama"`
	Description string `json:"description" validate:"required" name:"Deskripsi"`
	UserID      int64  `json:"user_id,omitempty" validate:"required"`
}

type TodoListCategoryResponse struct {
	ID          int64  `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedBy   string `json:"created_by"`
}

func (r *TodoListCategoryReq) SetID(ID int64) {
	r.ID = ID
}

func (r *TodoListCategoryReq) SetUserID(UserID int64) {
	r.UserID = UserID
}
