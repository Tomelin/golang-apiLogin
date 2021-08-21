package routesGenerate

import (
	"apiLogin/src/controllers"
	"net/http"
)

var routesRegister = []RouteStruct{
	{
		URI:            "/login",
		Method:         http.MethodPost,
		Function:       controllers.Login,
		Authentication: false,
	},
	{
		URI:            "/logout",
		Method:         http.MethodPost,
		Function:       controllers.Logout,
		Authentication: false,
	},
	{
		URI:            "/logout",
		Method:         http.MethodGet,
		Function:       controllers.Logout,
		Authentication: false,
	},
}
