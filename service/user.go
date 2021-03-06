/*
 * Copyright (c) 2021. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package service

import (
	"github.com/gin-gonic/gin"
	"lbc/model"
	"lbc/util"
	"log"
	"net/http"
)

func (s *Service) CreateUser(ctx *gin.Context) {
	var payload model.User
	err := ctx.BindJSON(&payload)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err = payload.Valid()
	if err != nil {
		log.Printf("email not valid %v",err)
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	log.Printf("payload create user %#v", payload)

	err = s.db.AddUser(&payload)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(200, payload)
}

func (s *Service) DeleteUser(ctx *gin.Context) {
	err := s.db.DeleteUser(ctx.Param("id"))
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(200, nil)
}

func (s *Service) GetUser(ctx *gin.Context) {
	u, err := s.db.GetUserByID(ctx.Param("id"))
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(200, u)
}

func (s *Service) GetUsers(ctx *gin.Context) {
	users, err := s.db.GetUsers()
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (s *Service) Login(ctx *gin.Context) {
	var payload model.Login
	ctx.BindJSON(&payload)
	user, err := s.db.GetUserByEmail(payload.Email)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	if user.Password != payload.Password {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claim := util.NewClaim(user.ID,user.FirstName,user.AccessLevel)
	jwe,err := util.CreateJWE(claim)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"jwt": jwe,
	})

}
