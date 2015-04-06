package model

type File struct {
	IsLocked bool   `json:"is_locked"`
	Content  string `json:"content"`
}
