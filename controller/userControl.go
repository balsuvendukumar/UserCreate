package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/balsuvendukumar/UserCreate/service"
	"github.com/balsuvendukumar/UserCreate/util"
)

func GetUserDetail(w http.ResponseWriter, r *http.Request) {
	fmt.Print("key befor", r)
	key := r.URL.Query().Get("user_id")
	fmt.Print("key befor", key)
	keyIsInt, error := strconv.ParseInt(key, 10, 64)
	fmt.Println(keyIsInt)
	if error != nil {
		err := &util.UserError{http.StatusBadRequest, "key is not found to be numerical"}
		json.NewEncoder(w).Encode(err)
		return
	}
	userDetail, resError := service.FindUserDetail(keyIsInt)
	if resError != nil {
		json.NewEncoder(w).Encode(resError)
		return
	}
	json.NewEncoder(w).Encode(userDetail)
}
