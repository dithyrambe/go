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

var _ db.Store = &Moke{}

type Moke struct {
	listUser map[string]*model.User
}

func New() *Moke {
	return &Moke{
		listUser: make(map[string]*model.User),
	}
}

func (m *Moke) SponeUser() {

	m.listUser = map[string]*model.User{
		"a0d09b91-0dcb-4aae-a0de-3a797891666c": &model.User{
			ID:        "a0d09b91-0dcb-4aae-a0de-3a797891666c",
			FirstName: "Bob",
			LastName:  "Picke",
			Email: "bob.picke@domain.fr",
			Password: "47625ed74cab8fbc0a8348f3df1feb07f87601e34d62bd12eb0d51616566fab5", // 123password
			AccessLevel: 0,
		},
		"2c3759ce-1c08-47f4-9dcc-bc5fb61f4d66": &model.User{
			ID:        "2c3759ce-1c08-47f4-9dcc-bc5fb61f4d66",
			FirstName: "René",
			LastName:  "Leblanc",
			Email: "r@google.com",
			Password: "51d59a88b0e1f1ad672dad1fb56626c2169a737ac5d6cd544355c53bc495e769", // monmdp
			AccessLevel: 1,
		},
		"05eca63d-c6e5-4a38-ba4f-65ff100c17bc": &model.User{
			ID:        "05eca63d-c6e5-4a38-ba4f-65ff100c17bc",
			FirstName: "Dennis",
			LastName:  "Richie",
			Email: "dennis.richie@domain.fr",
			Password: "47625ed74cab8fbc0a8340f3df1feb07f87601e34d62bd12eb0d51616566fab5", // ???
			AccessLevel: 1,
		},
	}
}

func (m *Moke) GetUserByID(id string) (*model.User, error) {
	u := m.listUser[id]
	fmt.Printf("user value:%v type:%T\n", u, u)
	if u == nil {
		return nil, db.NewNotFound("db/moke")
	}
	return u, nil
}

func (m *Moke) GetUserByEmail(email string) (*model.User, error) {
	for k := range m.listUser {
		if m.listUser[k].Email == email {
			return m.listUser[k], nil
		}
	}
	return nil, db.NewNotFound("db/moke")
}

func (m *Moke) DeleteUser(id string) error {
	_, err := m.GetUserByID(id)
	if err != nil {
		return db.NewNotFound("db/moke", err)
	}
	delete(m.listUser, id)
	return nil
}

func (m *Moke) AddUser(u *model.User) error {
	u.ID = uuid.NewString()
	m.listUser[u.ID] = u
	return nil
}

func (m *Moke) GetUsers() ([]*model.User, error) {
	users := make([]*model.User, 0, len(m.listUser))
	for k := range m.listUser {
		users = append(users, m.listUser[k])
	}
	return users, nil
}
