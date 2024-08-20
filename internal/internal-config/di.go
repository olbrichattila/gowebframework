package internalconfig

import (
	"framework/internal/app/args"
	commandexecutor "framework/internal/app/command"
	"framework/internal/app/cron"
	"framework/internal/app/db"
	"framework/internal/app/env"
	"framework/internal/app/event"
	"framework/internal/app/logger"
	"framework/internal/app/mail"
	"framework/internal/app/queue"
	"framework/internal/app/request"
	"framework/internal/app/session"
	"framework/internal/app/storage"
	"framework/internal/app/view"
	wizard "framework/internal/app/wizards/class"
	commandcreator "framework/internal/app/wizards/command"
	"os"

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

var DiBindings = map[string]interface{}{
	"internal.app.args.CommandArger": args.New(),
	"internal.app.view.Viewer":       view.New(),
	"internal.app.request.Requester": request.New(),
	"internal.app.env.Enver":         env.New(),
	"internal.app.db.DBFactoryer":    db.NewDBFactory(),
	// "internal.app.db.DBer":                   db.New(),
	"internal.app.db.DBer":                        getOpenedDb,
	"internal.app.session.Sessioner":              session.New(storage.NewFileStorage()),
	"internal.app.cron.JobTimer":                  cron.New(),
	"internal.app.queue.Quer":                     queue.New(),
	"internal.app.mail.Mailer":                    mail.New(),
	"internal.app.logger.Logger":                  logger.New(storage.NewFileStorage()),
	"internal.app.wizards.command.CommandCreator": commandcreator.New(),
	"internal.app.command.CommandExecutor":        commandexecutor.New(),
	"internal.app.wizards.class.ClassCreator":     wizard.NewClassCreator(),
	"internal.app.event.Eventer":                  event.NewLocalEvent(),
	"olbrichattila.gosqlbuilder.pkg.Builder":      getSqlBuilder,
}
