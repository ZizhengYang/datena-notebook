package models

type NoteBase struct {
	Id       int
	NoteName string
	Desc     string
}

type Tag struct {
	Id        int
	TagString string
	Num       int
}

type Issue struct {
	Id    int
	Title string
	Open  bool
}

type Question struct {
	Id    int
	Title string
}

type Note struct {
	Id       int
	NoteName string
}


