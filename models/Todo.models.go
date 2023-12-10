package models

type TodoModel struct {
	Id          int    `gorm:"primaryKey" column:"id" json:"id"`
	Title       string `column:"title" json:"title"`
	Description string `column:"description" json:"description"`
	IsCompleted int    `column:"is_completed" json:"is_completed"`
}

func (TodoModel) TableName() string {
	return "todos"
}
