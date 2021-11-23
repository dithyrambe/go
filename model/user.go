/*
 * Copyright (c) 2021. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package model

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

type User struct{
	ID string `json:"uuid"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Password Pass `json:"password,omitempty"`
}

type Login struct{
	Email string `json:"email"`
	Password Pass `json:"password,omitempty"`
}


type Pass string

func (p *Pass) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	h := sha256.New()
	h.Write([]byte(s))
	*p = Pass(fmt.Sprintf("%x", h.Sum(nil)))
	return nil
}

func (p Pass) MarshalJSON() ([]byte, error) {
	return json.Marshal(nil)
}


func (u *User)Hello()string{
	return fmt.Sprintf("%v",u.FirstName)
}
