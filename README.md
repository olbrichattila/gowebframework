# Golang Web Framework 

> This framework is work in progress, lot's to come, see todo.txt. including full code cleanup and testing

### Available main features.
- Controllers
- Middlewares
- Commands
- Jobs
- Events
- Mailer
- Views

- Migrations

## Other features:
- Automatic dependency injection in controllers, middlewares, jobs, commands, event consumers
- Queue
- Logger
- SQL Builder
- Config
- Request
- Router
- Session (file only, others, redis, db... to come soon)
- Storage (file only, others, redis, db... to come soon)
- Built in artisan commands

## Artisan commands:
```
go run ./cmd/ artisan
```

```
- create:command
- create:controller
     possible flags: (-api, -rest -in= -out=). try -help for more details
- create:custom-validator-rule
- create:event-consumer
- create:job
     possible flags: (-in= -out=). try -help for more details
- create:middleware
     possible flags: (-in= -out=). try -help for more details
- create:view-function
- list-commands
```
### create:command
Create a new blank command into app/commands folder
Usage:
```
go run ./cmd/ artisan create:command data-list


Please register your new command in:
  app/config/commands.go
```


It will create:
```
package command

import (
	"framework/internal/app/args"
)

// DataListCommand function can take any parameters defined in the Di config
func DataListCommand(a args.CommandArger) {
}
```

Mapping:
Add to config.Commands
```
var ConsoleCommands = map[string]commandexecutor.CommandItem{
	"command-name": {Fn: command.DataListCommand, Desc: "add usage info here"},
}
```

### Create controller
Usage:

help:
```
go run ./cmd/ artisan create:controller -help
```

```go run ./cmd/ artisan create:controller <controller-name> <optional-parameters>```

Template variations if set, otherwise it will be the default:
- api
- crud

Optional -in parameter values:
- cargs: a args.CommandArger
- config: c config.Configer
- db: db db.DBer
- logger: l logger.Logger
- mail: m mail.Mailer
- queue: q queue.Quer
- request: r request.Requester
- response: w http.ResponseWriter
- session: s session.Sessioner
- sqlBuilder: sqlBuilder builder.Builder
- view: v view.Viewer

Optional -out parameter values:
- bool: bool
- error: error
- string: string
```
go run ./cmd/ artisan create:controller products -in=config,db,logger,mail -out=string,error

This will create a controller with return parameters string and error, input parameters on actions resolving config,db,logger,mail
```
```
package controller

import (
     "framework/internal/app/config"
     "framework/internal/app/db"
     "framework/internal/app/logger"
     "framework/internal/app/mail"
)

// ProductsAction function can take any parameters defined in the Di config
func ProductsAction(c config.Configer, db db.DBer, l logger.Logger, m mail.Mailer) (string, error) {
     return "", nil
}
```
### Pre defined controller for API and CRUD
```
go run ./cmd/ artisan create:controller products-api -api -in=config,db,logger,mail -out=string,error
```

Generates:
```
package controller

import (
     "framework/internal/app/config"
     "framework/internal/app/db"
     "framework/internal/app/logger"
     "framework/internal/app/mail"
)

// IndexProductsApiAction function can take any parameters defined in the Di config
func IndexProductsApiAction(c config.Configer, db db.DBer, l logger.Logger, m mail.Mailer) (string, error) {
     return "", nil
}

// StoreProductsApiAction function can take any parameters defined in the Di config
func StoreProductsApiAction(c config.Configer, db db.DBer, l logger.Logger, m mail.Mailer) (string, error) {
     return "", nil
}

// ShowProductsApiAction function can take any parameters defined in the Di config
func ShowProductsApiAction(c config.Configer, db db.DBer, l logger.Logger, m mail.Mailer) (string, error) {
     return "", nil
}

// UpdateProductsApiAction function can take any parameters defined in the Di config
func UpdateProductsApiAction(c config.Configer, db db.DBer, l logger.Logger, m mail.Mailer) (string, error) {
     return "", nil
}

// DestroyProductsApiAction function can take any parameters defined in the Di config
func DestroyProductsApiAction(c config.Configer, db db.DBer, l logger.Logger, m mail.Mailer) (string, error) {
     return "", nil
}
```

### Generating CRUD controller:
```
go run ./cmd/ artisan create:controller products-crud -crud -in=config,db,logger,mail -out=string,error
```
Generates:
```
package controller

import (
     "framework/internal/app/config"
     "framework/internal/app/db"
     "framework/internal/app/logger"
     "framework/internal/app/mail"
)

// IndexProductsCrudAction function can take any parameters defined in the Di config
func IndexProductsCrudAction(c config.Configer, db db.DBer, l logger.Logger, m mail.Mailer) (string, error) {
     return "", nil
}

// CreateProductsCrudAction function can take any parameters defined in the Di config
func CreateProductsCrudAction(c config.Configer, db db.DBer, l logger.Logger, m mail.Mailer) (string, error) {
     return "", nil
}

// StoreProductsCrudAction function can take any parameters defined in the Di config
func StoreProductsCrudAction(c config.Configer, db db.DBer, l logger.Logger, m mail.Mailer) (string, error) {
     return "", nil
}

// EditProductsCrudAction function can take any parameters defined in the Di config
func EditProductsCrudAction(c config.Configer, db db.DBer, l logger.Logger, m mail.Mailer) (string, error) {
     return "", nil
}

// UpdateProductsCrudAction function can take any parameters defined in the Di config
func UpdateProductsCrudAction(c config.Configer, db db.DBer, l logger.Logger, m mail.Mailer) (string, error) {
     return "", nil
}

// DestroyProductsCrudAction function can take any parameters defined in the Di config
func DestroyProductsCrudAction(c config.Configer, db db.DBer, l logger.Logger, m mail.Mailer) (string, error) {
     return "", nil
}
```
### Mapping controllers to routes
Add to: ```app/config/routes.go```

```
{
    Path:        "/vehicles",
    RequestType: http.MethodGet,
    Fn:          controller.DisplayVehicles,
    Middlewares: AuthMiddleware,
},
```
### Generating middlewares

it is exactly the same process as generating controllers, except the command called create:middleware. Please see the description above.
Note: Middleware can have only bool or no return parameter
- If no return parameter, the next middleware will process until we get to controller
- If the return parameter is false, the middleware stops processing further. Here is a good place to redirect after failed authorization for example.

### Register your middleware
```app/config/middlewares.go```

```
// Add middlewares here to execute at every load
var Middlewares = []interface{}{
	middleware.SessionMiddleware,
}

// Add middlewares here, or create similar middleware groups and use them in your route
var AuthMiddleware = []interface{}{
	middleware.AuthorizeMiddleware,
}
```
### Create new job
It is very similar to controllers and middlewares but there is no return (-out)

Usage:
```
go run ./cmd/ artisan create:job clear-payments -in=queue
 
Please register your new job in:
  app/config/job.go
```

Generated:
```
package job

import (
     "framework/internal/app/queue"
)

// ClearPaymentsJob function can take any parameters defined in the Di config
func ClearPaymentsJob(q queue.Quer) {
}
```
### Schedule a Job

Provide in seconds how frequently it should be triggered
```
// app/config/job.go

var Jobs = []cron.Job{
	{Seconds: 5, Fn: job.SendRegistrationEmail},
	{Seconds: 30, Fn: job.ExpireEmailConfJob},
}
```
### Create events
example:
```
go run ./cmd/ artisan create:event-consumer my-first-event
```

Generated file:
```
package eventconsumer

// MyFirstEventEventConsumer function takes first parameter as string, which is the event payload
// after can take any parameters defined in the Di config
func MyFirstEventEventConsumer(payload string) {
}
```

Dispatch event which will be consumed to events subscribed to it.

```
// Bootstrap is a good place to subscribe to events globally, but can in any type of function, like controllers, 
// you can usubscribe if you need
func Bootstrap(l logger.Logger, e event.Eventer) {
	// Example event subscriber, consumer
	e.Subscribe("topic", "example1", eventconsumer.ExampleConsumer)
	e.Subscribe("topic", "example2", eventconsumer.ExampleConsumer2)
}

// Unsubscribe events:
e.UnSubscribe("topic", "example1")
e.UnSubscribe("topic", "example2")


// Dispatch anywhere, all events consumers will pick it up with it's payload asynchronously 
func MyAction(e event.Eventer) string {
	// Payload should be a string, good way to use json marshal, and unmarshal in consumer if you pass complex data
	payload := "{\"event\": \"dispatcher\"}"
	e.Dispatch("topic", payload)

	return "Event dispatched"
}
```
## Create view functions
View functions are specific user defined functions which you can add to your views.
Use them:

```{{ myFunc . }}```

Command to create:
```
go run ./cmd/ artisan create:view-function ucfirst
```
Creates something like:
```
package viewfunction

// UcfirstViewFunction function can take and return any primitives string, int, int64, bool and more...
func UcfirstViewFunction(s string) string {
    return "TODO"
}
```

Where you can implement your logic, for example Capitalize first character of the string and use like ```{{ ucFirst . }}```
after registering the function name in ```app/config/view.go```

```
// don't forget to import viewfunction package at the top
var ViewFuncConfig = template.FuncMap{
	"ucFirst": viewfunction.UcfirstViewFunction,
}
```

Please register your new command in:
  app/config/view.go

## Custom validators:

You can add custom validators to your validator, It can be used in route validator and as a dependency injected validator as well, Pease see the validator section:

Example:
```go run ./cmd/ artisan create:custom-validator-rule my-custom```
It generates: ```app/validator-configs/my-custom-rule.go```
```
package customrule

import "framework/internal/app/validator"

// MyCustomRule is a custom validator rule,
// val is the value to validate,
// pars is the elements in the rule signature, like myrule:1,2,3 will be 1, 2 and 3
// returns error message and bool if validation is OK
func MyCustomRule(val string, pars ...string) (validator.ValidationErrors, bool) {
	return nil, true
}

```

Map your rule to a new rule name/parameter pair.
Example: ```app/config/validators.go```
```
var ValidatorRules = map[string]validator.RuleFunc{
	"myRule": customrule.MyCustomRule,
}
```

Usage:
```
rules := map[string]string{
		"fieldName":  "min:5|max:80",
		"fieldName2": "max:55",
		"fieldName5": `myRule:10,50`,
	}
	ok, messages, validated := val.Validate(values, rules)
```

alternatively add it to your route validation rule config: ```app/config/route-validation-rules.go```
```
var RouteValidationRules = map[string]ValidationRule{
	"register": {
		Redirect: "/gohere",
		Rules: map[string]string{
			"password": "minSize:6|maxSize:255",
			"name":     "minSize:6|maxSize:255",
			"email":    "myRule:50",
		},
	},
}
```

## Map view partials to auto load:
If a view use a partial it should be auto-loaded.
We have two categories, View and Mail.
If you use the ```v.RenderMail``` it will load the mail partials, if you use the ```v.RenderView``` it will load the configured view partials:
Using simple ```v.Render``` will not load any view, but you can add any custom view for that specific render with: ```LoadTemplateParts```
Example:

```
templateFiles := []string{
		"register.html",
		"templates/mypartial.html",
	}

v.LoadTemplateParts()
```

The configuration for view partial auto-load:
```
// app/config/view-autoload.go
var TemplateAutoLoad = map[string][]string{
	view.ViewTypeHTML: {
		"template/head.html",
		"template/header.html",
		"template/footer.html",
	},
	view.ViewTypeEmail: {},
}
```

## Migrations
Install migrator
```
go install github.com/olbrichattila/godbmigrator_cmd/cmd/migrator@latest
```
Please see: https://github.com/olbrichattila/godbmigrator_cmd how to configure your DB, this configuration will be used with the framework as well.

then you can:
```
migrator migrate
migrator rollback
migrator migrate 2
migrator rollback 2
migrator report
migrator add <optional suffix>
```

### Sessions:
Type hint in your controller, job, middleware or other functions like: ```s session.Sessioner```
Usage:

In your session middleware if you not commented out from middleware mapping, the session ID automatically initialized
```
func SessionMiddleware(w http.ResponseWriter, r request.Requester, s session.Sessioner) {
	s.Init(w, r.GetRequest())
}
```
## Use your session:

```
// Set new session value
s.Set("sessionkeyname", email)

// remove session key
s.Delete("sessionkeyname")

// Redirect (HTTP)
s.Redirect("/login")

// Flush session data
s.Close()

// Remove session cookie
s.RemoveSession()
```

## Rendering HTML views
Example in controller:
```
func RegisterAction(v view.Viewer) string {
    // Data passing to the template
	data := map[string]string{
		"regUserEmail": s.Get("regUserEmail"),
		"regUserName":  s.Get("regUserName"),
		"lastError":    s.Get("lastError"),
	}

    //  v.Render will return the rendered template which is in your app/views folder
	return v.RenderView("register.html", data)
}
```

## Built in functions
- urlEscape (exapmle {{ urlEscape . }}) 
- further to come

## Adding custom functions for a particular view

```
func yourFunc(str string) string {
	return "it is your " + str
}

-------
funcMap := template.FuncMap{
		"yourFunc": yourFunc,
	}
v.Funcs(funcMap)

```

### Adding custom function in config for all views

```
// app/config/view
var ViewFuncConfig = template.FuncMap{
	"myFuncName" = customfunctions.CustomFunction,
}
```

## Return any text:
```
func YourAction() string {
    return "Hello World"
}
```

## Return json, method 1: using struct
```
type Response struct {
	Name  string `json:"name"`
}

func TestAction() Response {
	resp := Response{Name: "Vehicle makes"}

	return resp
}
```

## Return json, returning map
```
func TestAction(r request.Requester) map[string][]string {
    // Get the request as map[string][]string and return it
	return r.All()
}
```

## Return any struct
```
type Response = struct {
	Name string `json:"name"`
}

func TestAction(r request.Requester) Request {
	resp := &Response{}

    // Where JSONToStruct will marshall the request into the struct, it is a simple way to load request payload to structs
	r.JSONToStruct(resp)

	return *resp
}
```

## DB module
Example:
```
type Response struct {
	Name  string                   `json:"name"`
	Data  []map[string]interface{} `json:"data"`
	Error error                    `json:"error"`
}

func TestAction(r request.Requester, db db.DBer) (Response, error) {
    // database comes open, it is a new instance, we can safely close
    defer db.Close()
	resp := Response{Name: "Vehicle makes"}
	
	data := db.QueryAll("select make, count(*) as cnt from data group by make")
	for d := range data {
		resp.Data = append(resp.Data, d)
	}
	lastError := db.GetLastError()
	if lastError != nil {
		return nil, lastError
	}

	resp.Error = lastError

	return resp
}
```

### Database functions:
```
Open()
Close()
QueryAll(string, ...any) <-chan map[string]interface{}
QueryOne(string, ...any) (map[string]interface{}, error)
Execute(string, ...any) (int64, error)
GetLastError() error
```

### Query Builder:

Example:
```
// Type hint in function to be resolved as: ```sqlBuilder builder.Builder```

sqlBuilder.Select("users").Fields("id", "password").Where("email", "=", email).IsNotNull("activated_at")
sql, err := sqlBuilder.AsSQL()
if err != nil {
    return
}

params := sqlBuilder.GetParams()
res, err := db.QueryOne(sql, params...)
if err != nil {
    return
}
```

To see all features look at the package documentation at: https://github.com/olbrichattila/gosqlbuilder

## Logger:
Example:
```
func Bootstrap(l logger.Logger) {
	l.Info("bootstrap.called")
}
```

Possible log levels:
- Info(string)
- Warning(string)
- Error(string)
- Critical(string)

The log can be find in ./log/app.log

## Queue
it is currently internal queue, using it's own db, new versions to come

Example:
````
func TestQueJob(q queue.Quer, m mail.Mailer, v view.Viewer, l logger.Logger) {
	res, err := q.Pull("register")
	if err != nil {
		return
	}
    // ... work with RES which is a map[string]interface{}, which can be converted to json
}
````
Dispatch to queue
````
func TestQueJob(q queue.Quer) {
    email := "email@email.com"
    name := "member"
    
	q.Dispatch("register", "register-user", map[string]interface{}{"email": email, "name": name})
}
````

## Local event publisher, subscriber
You can dispatch events, and who ever subscribed to the event will get the event payload which is a string

Example event class:
```
package eventconsumer

import (
	"fmt"
	"framework/internal/app/logger"
)

func ExampleConsumer(payload string, l logger.Logger) {
	l.Info(fmt.Sprintf("Event %s consumed", payload))
}
```

Dispatch, consume:
```
func Bootstrap(l logger.Logger, e event.Eventer) {
	// Example event subscriber, consumer
	e.Subscribe("topic", "eventname", eventconsumer.ExampleConsumer)
	e.Dispatch("topic", "event1")
	l.Info("bootstrap.called")
}
```

## Event features
```
Subscribe(topic, name string, interface{})
UnSubscribe(topic, name string)
Flush()
Dispatch(topic, payload string)
```

## Cache:
The cache module caches string results, can be used with closure or using methods on cache instance.
Example:
```

func youFunc(...., c cache.Cacher) {

	cached := c.Cache("models", func(_ ...interface{}) string {
		return v.RenderView("make.html", report)
	}), nil
}
```

## Mailer
Example send registration main from consuming a queue event published by a user registration
```
func SendRegistrationEmail(q queue.Quer, m mail.Mailer, v view.Viewer, l logger.Logger) {
	res, err := q.Pull("register")
	if err != nil {
		return
	}

	email, ok := res["email"]
	if !ok {
		l.Error("Missing email from the message")
		return
	}

	rendered := v.RenderMail([]string{"regconfirm.html"}, res)
	err = m.Send("attila@osoft.hu", email.(string), "Please confirm your email address", rendered)
	if err != nil {
		l.Error(err.Error())
		return
	}

	l.Info(fmt.Sprintf("Registration mail sent to %s", email))
}
```

## Validator:

Example using of the validator, (new rules are coming soon, custom rules as well, will be able to automatically attach to request soon)
```
package controller

import (
	"fmt"
	"framework/internal/app/session"
	"framework/internal/app/validator"
)

funcMyControllerAction(s session.Sessioner, val validator.Validator) {
	values := map[string]string{
		"fieldName":  "33",
		"fieldName2": "54",
		"fieldName3": "hello",
		"fieldName4": "hello,| world",
		"fieldName5": "60",
	}
	rules := map[string]string{
		"fieldName":  "min:5|max:80",
		"fieldName2": "max:55",
		"fieldName3": "in:a,bc,de,hello,bukk",
		"fieldName4": `regex:^hello,\\|.*world$`,
		"fieldName5": `between:10,50`,
	}
	ok, messages, validated := val.Validate(values, rules)
	fmt.Printf("%v\n%v\n%v\n\n", ok, messages, validated)
}
```

Current validation rules:
- required
- min
- max
- in
- regex
- between
- size
- email
- url
- uuid
- numeric
- integer
- date
- dateTime
- boolean
- json

## Route validation

You can set route level validation rules. 

Example:
```ValidationRules: "rule-name",```
This is not mandatory to set on the route, if not set, no validation, but you are free to use validator in your controller individually
```
// in your route configuration:
	{
		Path:            "/doregister",
		RequestType:     http.MethodPost,
		Fn:              controller.PostRegister,
		ValidationRules: "register",
	},
```

In your validator config
- Redirect (not mandatory, if you set on failure it will redirect here, if not set then it will continue calling your controller action)
- Rules (non mandatory, if not set, no validation takes place using validator class)
- CustomRule (non mandatory. This is a function ```func(fields map[string]string) (validator.ValidationErrors, bool) {}```. If  you set then you receive the parameters (post, get, route as well) in a map, you return a validation error message, and bool for OK or not )


```
// app/config/route-validation-rules.go
var RouteValidationRules = map[string]ValidationRule{
	"register": {
		Redirect: "/register",
		Rules: map[string]string{
			"password": "minSize:6|maxSize:255",
			"name":     "minSize:6|maxSize:255",
			"email":    "email",
		},
		CustomRule: func(fields map[string]string) (validator.ValidationErrors, bool) {
			return validator.ValidationErrors{"name": []string{"error1", "error2"}}, false
		},
	},
}
```

If you are using this route validation, errors will be stored in session ```lastError``` and can be displayed on the page, or json as your choice from getting it from the session.


## Bootstrapping the application
```
// app/bootstrap.go

// Bootstrap is always called after application started, put whatever want to initiate here
func Bootstrap(l logger.Logger, e event.Eventer) {
	// Example event subscriber, consumer
	e.Subscribe("topic", "e4", eventconsumer.ExampleConsumer)
	e.Dispatch("topic", "event1")
	l.Info("bootstrap.called")
}
```

## Running the web server:
```
go run ./cmd
```

.. To be continued

## Custom dependencies
you can provide your own dependencies.
Map them in ```app/config/di.go```
Example:
```
var DiBindings = []config.DiCallback{
	func(di godi.Container) (string, interface{}, error) {
		env, err := di.Get(env.New())
		return "internal.app.env.Enver", env, err
	},
	...
	...
```
## .env variables
```
## redis
REDIS_SERVER_HOST=localhost

## memcace
MEMCACHE_HOST=localhost

## mail
SMTP_USER_NAME=mailtrap
SMTP_PASSWORD=mailtrap
SMTP_HOST=localhost
SMTP_PORT=1025

## Session storage
SESSION_STORAGE=file
SESSION_STORAGE=redis
SESSION_STORAGE=db
SESSION_STORAGE=memcached

## Logger storage
LOGGER_STORAGE=file
LOGGER_STORAGE=redis
LOGGER_STORAGE=db
LOGGER_STORAGE=memcached
```
