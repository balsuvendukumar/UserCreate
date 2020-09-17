package app

import (
	"fmt"
	"net/http"
)

func StartApp() {
	Address()
	fmt.Println("listen 1")

	http.ListenAndServe(":8080", nil)
	//fmt.Println("listen 2")
}
