package entity

type TodoListCategoryReq struct {
	Name        string `json:"name" validate:"required" name:"Nama"`
	Description string `json:"description" validate:"required" name:"Deskripsi"`
}

type TodoListCategoryResponse struct {
	ID          int64  `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
