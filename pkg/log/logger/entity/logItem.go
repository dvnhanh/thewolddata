package entity

type LogItem struct {
	Message string
	Data    interface{}
	Error   error
}
