package entity

type Category struct {
	ID			int		`json:"id"`
	Title		string	`json:"title" binding:"required"`
	ParentID	int		`json:"parent_id"`
}

