package domain

import (
	"net/http"

	"github.com/balsuvendukumar/UserCreate/util"
)

var (
	m = map[int64]*UserDetail{
		123: {"suvendu", "bal", 123, 1000, "july"}}
)

func GetUserDGetUserDetail(key int64) (*UserDetail, *util.UserError) {
	UserDetails, ok := m[key]
	if ok {
		return UserDetails, nil
	}
	return nil, &util.UserError{ErrorStatusCode: http.StatusNotFound, ErrorStatusMessage: "User iD is not available on database"}

}
