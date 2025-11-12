package main

type User struct {
	Name   string `json:"name"`
	Age    int16  `json:"age"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	Gender Gender `json:"gender"`
}

type Gender string

const (
	Male   Gender = "Male"
	Female Gender = "Female"
)
