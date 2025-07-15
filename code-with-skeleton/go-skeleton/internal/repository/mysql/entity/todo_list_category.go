package entity

import "time"

type TodoListCategory struct {
	ID          int64     `gorm:"column:id"`
	CreatedBy   int64     `gorm:"column:created_by"`
	Name        string    `gorm:"column:name"`
	Description string    `gorm:"column:description"`
	CreatedAt   time.Time `gorm:"column:created_at"`
}
type TodoListCategorySelect struct {
	ID          int64     `gorm:"column:id"`
	CreatedBy   int64     `gorm:"column:created_by"`
	Name        string    `gorm:"column:name"`
	Description string    `gorm:"column:description"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UserName    string    `gorm:"column:user_name"`
}

func (TodoListCategory) TableName() string {
	return "todo_list_categories"
}
