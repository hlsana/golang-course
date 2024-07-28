package models

type Student struct {
	ID     string
	Name   string
	Scores map[string]int
}

type Teacher struct {
	Username    string
	Name        string
	Password    string
	IsAuthrized bool
}

var EnglishTeacher = Teacher{
	Username: "enteach",
	Password: "iloveborsch",
}

type Class struct {
	Name     string
	Students map[string]Student
}
