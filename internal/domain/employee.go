package domain

type Employee struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}
