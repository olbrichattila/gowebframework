package internalconfig

import (
	"framework/internal/app/args"
	"framework/internal/app/cache"
	commandexecutor "framework/internal/app/command"
	"framework/internal/app/config"
	"framework/internal/app/cron"
	"framework/internal/app/db"
	"framework/internal/app/env"
	"framework/internal/app/event"
	"framework/internal/app/logger"
	"framework/internal/app/mail"
	"framework/internal/app/queue"
	"framework/internal/app/request"
	"framework/internal/app/router"
	"framework/internal/app/session"
	"framework/internal/app/validator"
	"framework/internal/app/view"
	wizard "framework/internal/app/wizards/class"
	commandcreator "framework/internal/app/wizards/command"
	"os"

	"github.com/olbrichattila/godi"
	gosqlbuilder "github.com/olbrichattila/gosqlbuilder"
	pkg "github.com/olbrichattila/gosqlbuilder/pkg"
)

func getOpenedDb() interface{} {
	return db.New()
}

func getSqlBuilder() interface{} {
	dbConnection := os.Getenv(db.EnvdbConnection)
	builder := gosqlbuilder.New()

	switch dbConnection {
	case db.DbConnectionTypeSqLite:
		builder.SetSQLFlavour(pkg.FlavourSqLite)
	case db.DbConnectionTypeMySQL:
		builder.SetSQLFlavour(pkg.FlavourMySQL)
	case db.DbConnectionTypePgSQL:
		builder.SetSQLFlavour(pkg.FlavourPgSQL)
	case db.DbConnectionTypeFirebird:
		builder.SetSQLFlavour(pkg.FlavourFirebirdSQL)
	}

	return builder
}

var DiBindings = []config.DiCallback{
	func(di godi.Container) (string, interface{}, error) {
		env, err := di.Get(env.New())
		return "internal.app.env.Enver", env, err
	},
	func(di godi.Container) (string, interface{}, error) {
		return "internal.app.args.CommandArger", args.New(), nil
	},
	func(di godi.Container) (string, interface{}, error) {
		return "internal.app.router.Router", router.NewRouter(), nil
	},
	func(di godi.Container) (string, interface{}, error) {
		return "internal.app.view.Viewer", view.New(), nil
	},
	func(di godi.Container) (string, interface{}, error) {
		return "internal.app.request.Requester", request.New(), nil
	},
	func(di godi.Container) (string, interface{}, error) {
		return "internal.app.validator.Validator", validator.New(), nil
	},
	func(di godi.Container) (string, interface{}, error) {
		return "internal.app.db.DBFactoryer", db.NewDBFactory(), nil
	},
	func(di godi.Container) (string, interface{}, error) {
		return "internal.app.db.DBer", getOpenedDb, nil
	},
	func(di godi.Container) (string, interface{}, error) {
		return "olbrichattila.gosqlbuilder.pkg.Builder", getSqlBuilder, nil
	},
	func(di godi.Container) (string, interface{}, error) {
		return "internal.app.logger.LoggerStorageResolver", logger.NewSessionStorageResolver(), nil
	},
	func(di godi.Container) (string, interface{}, error) {
		l, err := di.Get(logger.New())
		return "internal.app.logger.Logger", l, err
	},
	func(di godi.Container) (string, interface{}, error) {
		return "internal.app.session.SessionStorageResolver", session.NewSessionStorageResolver(), nil
	},
	func(di godi.Container) (string, interface{}, error) {
		session, err := di.Get(session.New())
		return "internal.app.session.Sessioner", session, err
	},

	func(di godi.Container) (string, interface{}, error) {
		return "internal.app.cache.CacheStorageResolver", cache.NewCacheStorageResolver(), nil
	},
	func(di godi.Container) (string, interface{}, error) {
		session, err := di.Get(cache.New())
		return "internal.app.cache.Cacher", session, err
	},

	func(di godi.Container) (string, interface{}, error) {
		return "internal.app.cron.JobTimer", cron.New(), nil
	},
	func(di godi.Container) (string, interface{}, error) {
		return "internal.app.queue.Quer", queue.New(), nil
	},
	func(di godi.Container) (string, interface{}, error) {
		return "internal.app.mail.Mailer", mail.New(), nil
	},

	func(di godi.Container) (string, interface{}, error) {
		return "internal.app.wizards.command.CommandCreator", commandcreator.New(), nil
	},
	func(di godi.Container) (string, interface{}, error) {
		return "internal.app.command.CommandExecutor", commandexecutor.New(), nil
	},
	func(di godi.Container) (string, interface{}, error) {
		return "internal.app.wizards.class.ClassCreator", wizard.NewClassCreator(), nil
	},
	func(di godi.Container) (string, interface{}, error) {
		return "internal.app.event.Eventer", event.NewLocalEvent(), nil
	},
}
