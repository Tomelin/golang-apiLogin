package router

import (
	"apiLogin/src/router/routesGenerate"

	"github.com/gorilla/mux"
)

func LoadRouter() *mux.Router {

	r := mux.NewRouter()
	return routesGenerate.RouteConfiguration(r)
}
