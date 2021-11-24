
package main

import (
	"github.com/gin-gonic/gin"
	"lbc/db/moke"
	"lbc/service"
	"lbc/util"
)

func init(){
	//TODO get config.
}

func main(){
	r := gin.Default()
	db := moke.New()
	db.SponeUser()
	s := service.New(db)
	r.GET("/users", util.MiddlewareVerifyJWE(1) ,s.GetUsers)
	r.GET("/users/:id", util.MiddlewareVerifyJWE(0), s.GetUser)
	r.POST("/users",  s.CreateUser)
	r.DELETE("/users/:id", util.MiddlewareVerifyJWE(1), s.DeleteUser)
	r.POST("/login", s.Login)
	r.Run()
}
