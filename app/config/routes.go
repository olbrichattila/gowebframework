package appconfig

import (
	controller "framework/app/controllers"
	"framework/internal/app/router"
	"net/http"
)

var Routes = []router.ControllerAction{
	{
		Path:        "/login",
		RequestType: http.MethodGet,
		Fn:          controller.Login,
	},
	{
		Path:        "/logout",
		RequestType: http.MethodGet,
		Fn:          controller.Logout,
		Middlewares: AuthMiddleware,
	},
	{
		Path:        "/dologin",
		RequestType: http.MethodPost,
		Fn:          controller.LoginPost,
	},
	{
		Path:        "/register",
		RequestType: http.MethodGet,
		Fn:          controller.Register,
	},
	{
		Path:        "/doregister",
		RequestType: http.MethodPost,
		Fn:          controller.PostRegister,
	},
	{
		Path:        "/confirm",
		RequestType: http.MethodGet,
		Fn:          controller.ActivateAction,
	},
	{
		Path:        "/error",
		RequestType: http.MethodGet,
		Fn:          controller.DisplayError,
	},
	{
		Path:        "/",
		RequestType: http.MethodGet,
		Fn:          controller.DisplayAllMakes,
		Middlewares: AuthMiddleware,
	},
	{
		Path:        "/model",
		RequestType: http.MethodGet,
		Fn:          controller.DisplayAllModels,
		Middlewares: AuthMiddleware,
	},
	{
		Path:        "/basemodel",
		RequestType: http.MethodGet,
		Fn:          controller.DisplayAllSubModels,
		Middlewares: AuthMiddleware,
	},
	{
		Path:        "/fuel_type",
		RequestType: http.MethodGet,
		Fn:          controller.DisplayAllFuelType,
		Middlewares: AuthMiddleware,
	},
	{
		Path:        "/year",
		RequestType: http.MethodGet,
		Fn:          controller.DisplayAllYear,
		Middlewares: AuthMiddleware,
	},
	{
		Path:        "/vehicles",
		RequestType: http.MethodGet,
		Fn:          controller.DisplayVehicles,
		Middlewares: AuthMiddleware,
	},
}
