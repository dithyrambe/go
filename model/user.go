/*
 * Copyright (c) 2021. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package model

import (
	"encoding/json"
	"fmt"
)

type User struct{
	ID string `json:"uuid"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Password string `json:"password,omitempty"`
}

func (u User) MarshalJSON() ([]byte, error) {
	u.Password = ""
	return json.Marshal(u)
}


func (u *User)Hello()string{
	return fmt.Sprintf("%v",u.FirstName)
}
