package main

import (
	"golang-courses-exercise-1/entity"
	"golang-courses-exercise-1/service"
)

func main() {
	userSvc := service.NewUserSvc()
	userSvc.Register(&entity.User{
		Username: "budi123",
		Email:    "budi123@gmail.com",
		Password: "password123",
		Age:      9,
	})
}
