package notebook

import "errors"

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Note struct {
	Id          int    `json:"id" db:"id"`
	Date        string `json:"date" db:"date"`
	Description string `json:"description" db:"description" binding:"required" `
}

type UsersNote struct {
	Id      int
	UserId  int
	NotesId int
}
type UpdateNoteInput struct {
	Date        *string `json:"date"`
	Description *string `json:"description"`
}

func (i UpdateNoteInput) Validate() error {
	if i.Date == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
