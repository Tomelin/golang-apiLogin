package routesGenerate

import (
	"apiLogin/src/middlewares"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type RouteStruct struct {
	URI            string
	Method         string
	Function       func(http.ResponseWriter, *http.Request)
	Authentication bool
}

func RouteConfiguration(r *mux.Router) *mux.Router {

	loadRoutes := routesRegister

	for _, routeLoop := range loadRoutes {
		r.HandleFunc(routeLoop.URI, routeLoop.Function)

		if routeLoop.Authentication {
			fmt.Printf("route authentication if")
			r.HandleFunc(routeLoop.URI, middlewares.Authentication(routeLoop.Function))
		}
	}

	return r
}
