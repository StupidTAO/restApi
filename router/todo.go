package router

import "time"

type Todo struct {
	id int		`json:"id"`
	name string	`json:"name"`
	completed bool	`json:"completed"`
	due time.Time	`json:"due"`
}

type Todos []Todo
