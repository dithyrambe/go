/*
 * Copyright (c) 2021. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package moke

import (
	"fmt"

	"github.com/google/uuid"

	"lbc/db"
	"lbc/model"
)

type Moke struct{
	listUser map[string]*model.User
}

func New()*Moke{
	return &Moke{
		listUser: make(map[string]*model.User),
	}
}

func (m *Moke)SponeUser(){


	m.listUser = map[string]*model.User{
		"a0d09b91-0dcb-4aae-a0de-3a797891666c" : &model.User{
			ID:"a0d09b91-0dcb-4aae-a0de-3a797891666c",
			FirstName: "Bob",
			LastName: "Picke",
		},
		"05eca63d-c6e5-4a38-ba4f-65ff100c17bc" : &model.User{
			ID:"05eca63d-c6e5-4a38-ba4f-65ff100c17bc",
			FirstName: "Dennis",
			LastName: "Richie",
		},
	}
}


func (m *Moke)GetUserByID(id string)(*model.User,error){
	u := m.listUser[id]
	fmt.Printf("user value:%v type:%T\n",u,u)
	if u == nil{
		return nil,db.NewNotFound("db/moke")
	}
	return u,nil
}

func(m *Moke) DeleteUser(id string)error{
	_, err := m.GetUserByID(id)
	if err != nil {
		return db.NewNotFound("db/moke",err)
	}
	delete(m.listUser, id)
	return nil
}

func (m *Moke)AddUser(u *model.User)error{
	u.ID = uuid.NewString()
    m.listUser[u.ID] = u
	return nil
}
