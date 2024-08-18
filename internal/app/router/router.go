package router

type ControllerAction struct {
	Path        string
	RequestType string
	Fn          interface{}
	Middlewares []any
}

// TODO add route phaser
