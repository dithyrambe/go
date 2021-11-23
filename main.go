
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
	s := service.New(moke.New())
	r.POST("/users", s.CreateUser)
	r.GET("/users/:id", s.GetUser)
	r.DELETE("/users/:id", s.DeleteUser)
	r.Run()
}
