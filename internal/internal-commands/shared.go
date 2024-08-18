package internalcommand

import wizard "framework/internal/app/wizards/class"

func getDefaultCreateDiMapping() map[string]wizard.ParameterInfo {
	return map[string]wizard.ParameterInfo{
		"request":    {Name: "request.Requester", Alias: "r", ImportPath: "\"framework/internal/app/request\""},
		"db":         {Name: "db.DBer", Alias: "db", ImportPath: "\"framework/internal/app/db\""},
		"logger":     {Name: "logger.Logger", Alias: "l", ImportPath: "\"framework/internal/app/logger\""},
		"sqlBuilder": {Name: "builder.Builder", Alias: "sqlBuilder", ImportPath: "builder \"github.com/olbrichattila/gosqlbuilder/pkg\""},
		"session":    {Name: "session.Sessioner", Alias: "s", ImportPath: "\"framework/internal/app/session\""},
		"view":       {Name: "view.Viewer", Alias: "v", ImportPath: "\"framework/internal/app/view\""},
		"mail":       {Name: "mail.Mailer", Alias: "m", ImportPath: "\"framework/internal/app/mail\""},
		"config":     {Name: "config.Configer", Alias: "c", ImportPath: "\"framework/internal/app/config\""},
		"response":   {Name: "http.ResponseWriter", Alias: "w", ImportPath: "\"net/http\""},
		"cargs":      {Name: "args.CommandArger", Alias: "a", ImportPath: "\"framework/internal/app/args\""},
		"queue":      {Name: "queue.Quer", Alias: "q", ImportPath: "\"framework/internal/app/queue\""},
	}
}
