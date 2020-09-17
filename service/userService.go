package service

import (
	"github.com/balsuvendukumar/UserCreate/domain"
	"github.com/balsuvendukumar/UserCreate/util"
)

func FindUserDetail(key int64) (*domain.UserDetail, *util.UserError) {
	return domain.GetUserDGetUserDetail(key)
}
