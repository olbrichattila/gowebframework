package job

import (
	"fmt"
	"framework/internal/app/db"
	"framework/internal/app/logger"
	"framework/internal/app/mail"
	"framework/internal/app/queue"
	"framework/internal/app/view"
	"time"

	builder "github.com/olbrichattila/gosqlbuilder/pkg"
)

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

func ExpireEmailConfJob(db db.DBer, bd builder.Builder, l logger.Logger) {
	defer db.Close()
	expireAt := time.Now().Add(-1 * time.Minute).Format("2006-01-02 15:04:05")
	sql, err := bd.Select("reg_confirmations").Fields("id", "user_id").Where("created_at", "<", expireAt).AsSQL()
	if err != nil {
		l.Error(err.Error())
		return
	}

	rows := db.QueryAll(sql, bd.GetParams()...)
	confRow := make([]map[string]interface{}, 0)
	for row := range rows {
		confRow = append(confRow, row)
	}
	for _, row := range confRow {
		id := row["id"]
		userId := row["user_id"]

		delUserSql, err := bd.Delete("users").Where("id", "=", userId).AsSQL()
		if err != nil {
			l.Error(err.Error())
			return
		}
		pars := bd.GetParams()
		_, err = db.Execute(delUserSql, pars...)
		if err != nil {
			l.Error(err.Error())
			return
		}

		delConfSql, err := bd.Delete("reg_confirmations").Where("id", "=", id).AsSQL()
		if err != nil {
			l.Error(err.Error())
			return
		}
		delPars := bd.GetParams()
		_, err = db.Execute(delConfSql, delPars...)
		if err != nil {
			l.Error(err.Error())
			return
		}
		l.Info(fmt.Sprintf("User %d ID expired", userId))
	}

	if db.GetLastError() != nil {
		l.Error(db.GetLastError().Error())
		return
	}
}
