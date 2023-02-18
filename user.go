package notebook

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name"`
	UserName string `json:"username" binding:"required"`
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

/*
CREATE TABLE users
(
id            serial       not null unique,
name          varchar(255) not null,
username      varchar(255) not null unique,
password_hash varchar(255) not null
);

CREATE TABLE notes
(
id          serial       not null unique,
date  varchar(255) not null,
description varchar(255)
);

CREATE TABLE users_notes
(
id      serial                                           not null unique,
user_id int references users (id) on delete cascade      not null,
note_id int references todo_lists (id) on delete cascade not null
);

*/
