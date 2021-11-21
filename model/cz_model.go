package model

// Question commitizen协议中的模型
type Question struct {
	Type    string
	Name    string
	Message string
	Options []Option
	Filter  Filter
}

// Filter commitizen协议中的模型
type Filter func(value string) string

// Option commitizen协议中的模型
type Option struct {
	Name  string
	Value string
}
