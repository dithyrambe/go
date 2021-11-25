package service

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"

	rick_morty "lbc/sdk/rick-morty"
)

func (s *Service) GetCharacters(ctx *gin.Context) {
	data, err := s.cache.Get(ctx.Request.RequestURI)
	if err != nil && err != redis.Nil {
		log.Println("real error from cache")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if data != nil {
		log.Println("get from cash")
		ctx.JSON(http.StatusOK, data)
		return
	}

	caracters, err := rick_morty.New().GetCaracters()
	data, err = json.Marshal(caracters)
	if err != nil {
		log.Println("err marshal",err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	log.Println("set new cash")
	err = s.cache.Set(ctx.Request.RequestURI, data)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(200,caracters)
}
