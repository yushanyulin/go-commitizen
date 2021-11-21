package model

/*
	commitizen models
*/
type Question struct {
	Type string
	Name string
	Message string
	Options []Option
	Filter  Filter
}

type Filter func(value string) string

type Option struct {
	Name  string
	Value string
}
