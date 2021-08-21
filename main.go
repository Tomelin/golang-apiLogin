package main

import (
	"apiLogin/src/router"
	"fmt"
	"net/http"
)

func main() {

	r := router.LoadRouter()
	http.ListenAndServe(fmt.Sprintf(":%d", 5002), r)
}
