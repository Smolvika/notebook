package notebook

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Notes struct {
	Id          int    `json:"id"`
	Date        string `json:"date_completion"`
	Description string `json:"description" `
}

type UserNotes struct {
	Id      int
	UserId  int
	NotesId int
}
