package service

import (
	"fmt"
	"golang-courses-exercise-1/entity"
)

type UserServiceIface interface {
	Register(user *entity.User)
}

type UserSvc struct{}

func NewUserSvc() UserServiceIface {
	return &UserSvc{}
}

func (u *UserSvc) Register(user *entity.User) {
	fmt.Println(user)
}
