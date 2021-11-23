
package main

import (
	"github.com/gin-gonic/gin"

	"lbc/db/moke"
	"lbc/service"
)

func init(){
	//TODO get config.
}

func main(){
	r := gin.Default()
	db := moke.New()
	db.SponeUser()
	s := service.New(db)
	r.GET("/users", s.GetUsers)
	r.GET("/users/:id", s.GetUser)
	r.POST("/users", s.CreateUser)
	r.DELETE("/users/:id", s.DeleteUser)
	r.Run()
}
