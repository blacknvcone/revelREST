package controllers

import (
	"github.com/revel/revel"
)

func init() {

	// Register startup functions with OnAppStart
	revel.InterceptFunc(Auth, revel.BEFORE, &UserController{})

}
