package job

import (
	"fmt"
	"framework/internal/app/db"
	"framework/internal/app/logger"
	"framework/internal/app/mail"
	"framework/internal/app/queue"
	"framework/internal/app/view"
	"time"

	"github.com/google/uuid"
	builder "github.com/olbrichattila/gosqlbuilder/pkg"
)

// PasswordReminderJob function can take any parameters defined in the Di config
func PasswordReminderJob(db db.DBer, sqlBuilder builder.Builder, q queue.Quer, m mail.Mailer, v view.Viewer, l logger.Logger) {
	defer db.Close()
	res, err := q.Pull("password-reminder")
	if err != nil {
		// That's ok, queue is empty
		return
	}

	email, ok := res["email"]
	if !ok {
		l.Error("Missing email from the message")
		return
	}

	sql, err := sqlBuilder.Select("users").Fields("id", "name").Where("email", "=", email).IsNotNull("activated_at").AsSQL()
	if err != nil {
		l.Info(err.Error())
	}

	data, err := db.QueryOne(sql, sqlBuilder.GetParams()...)
	if err != nil {
		l.Info(err.Error())
	}

	userId, ok := data["id"]
	if !ok {
		l.Error("cannot resolve use ID sending password reminder")
		return

	}

	uuid := uuid.New().String()
	createdAt := time.Now().Format("2006-01-02 15:04:05")
	sqlBuilder.Insert("password_reminders").Fields("user_id", "uuid", "created_at").Values(userId, uuid, createdAt)
	sql, err = sqlBuilder.AsSQL()
	if err != nil {
		l.Error(err.Error())
		return
	}

	_, err = db.Execute(sql, sqlBuilder.GetParams()...)
	if err != nil {
		l.Error(err.Error())
		return
	}

	emailParams := map[string]interface{}{
		"name": data["name"],
		"uuid": uuid,
	}

	rendered := v.RenderMail("password-reminder.html", emailParams)
	err = m.Send("attila@osoft.hu", email.(string), "Re: Forgotten password", rendered)
	if err != nil {
		l.Error(err.Error())
		return
	}

	l.Info(fmt.Sprintf("Password reminder mail sent to %s", email))
}

func ExpirePasswordReminderJob(db db.DBer, sqlBuilder builder.Builder, l logger.Logger) {
	defer db.Close()
	expireAt := time.Now().Add(-1 * time.Minute).Format("2006-01-02 15:04:05")

	sql, err := sqlBuilder.Delete("password_reminders").Where("created_at", "<", expireAt).AsSQL()
	if err != nil {
		l.Error(err.Error())
		return
	}

	_, err = db.Execute(sql, sqlBuilder.GetParams()...)
	if err != nil {
		l.Error(err.Error())
		return
	}
}
