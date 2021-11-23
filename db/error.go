/*
 * Copyright (c) 2021. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package db

import (
	"fmt"
	"net/http"
)

type ErrorDB struct{
	Code int
	packageName string
	Err []error
}

func (e ErrorDB)Error() string{
	return fmt.Sprintf("%v: code: %, error: %v",e.packageName,e.Code,e.Err)
}

func NewNotFound(pkgName string,err ...error) error {
	return &ErrorDB{
		Code: http.StatusNotFound,
		packageName: pkgName,
		Err:err,
	}
}

