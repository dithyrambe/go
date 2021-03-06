/*
 * Copyright (c) 2021. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package service

import (
	"lbc/cache"
	"lbc/db"
	rick_morty "lbc/sdk/rick-morty"
)

type Service struct{
	db db.Store
	cache cache.Cache
	sdkRick *rick_morty.Client
}

func New(db db.Store,cache cache.Cache)*Service{
	return &Service{
		db:db,
		cache: cache,
	}
}
