
package main

import (
	"github.com/gin-gonic/gin"
	"lbc/cache/redis"

	"lbc/db/sqlite"
	"lbc/service"
	"lbc/util"
)

func init(){
	//TODO get config.
}

func main(){
	r := gin.Default()
	db := sqlite.New("my.db")
	db.SponeUser()
	cache := redis.New("6379")
	s := service.New(db,cache)
	r.GET("/users", util.MiddlewareVerifyJWE(1) ,s.GetUsers)
	r.GET("/users/:id", util.MiddlewareVerifyJWE(0), s.GetUser)
	r.POST("/users",  s.CreateUser)
	r.DELETE("/users/:id", util.MiddlewareVerifyJWE(1), s.DeleteUser)
	r.POST("/login", s.Login)
	// Rik services.
	r.GET("/rick", s.GetCharacters)

	r.Run()
}

