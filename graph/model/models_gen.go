// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type LoginOutput struct {
	Token string `json:"token"`
}

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
