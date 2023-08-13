package structs

import (
	"gorm.io/gorm"
)

var (
	JsLang int8 = 1
	Pylang int8 = 2
	GoLang int8 = 3
)

type Bot struct {
	gorm.Model

	Name        string
	BotId       string
	AutoRestart bool
	Language    int8
}

type Response struct {
	Success bool
	Message string
	Data    AnyData
}

type AnyData map[string]interface{}

type User struct {
	gorm.Model

	Username string
	Password string
	Token    string
}
