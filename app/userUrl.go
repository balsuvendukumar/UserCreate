package app

import (
	"fmt"
	"net/http"

	"github.com/balsuvendukumar/UserCreate/controller"
)

func Address() {
	fmt.Println("handle start")
	http.HandleFunc("/user", controller.GetUserDetail)
	fmt.Println("handle end")
}
