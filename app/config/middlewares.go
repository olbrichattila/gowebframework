package appconfig

import middleware "framework/app/middlewares"

// Add middlewares here to execute at every load
var Middlewares = []interface{}{
	middleware.SessionMiddleware,
}

// Add middlewares here, or create similar middleware groups and use them in your route
var AuthMiddleware = []interface{}{
	middleware.AuthorizeMiddleware,
}
