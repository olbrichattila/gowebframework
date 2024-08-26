package controller

import (
	"framework/internal/app/db"
	"framework/internal/app/queue"
	"framework/internal/app/request"
	"framework/internal/app/session"
	"framework/internal/app/view"

	builder "github.com/olbrichattila/gosqlbuilder/pkg"
	"golang.org/x/crypto/bcrypt"
)

// PasswordControllerAction function can take any parameters defined in the Di config
func PasswordReminderControllerAction(v view.Viewer, s session.Sessioner) string {
	data := map[string]string{
		"lastError": s.Get("lastError"),
	}
	return v.RenderView("password-reminder.html", data)
}

func PasswordReminderPostControllerAction(r request.Requester, s session.Sessioner, q queue.Quer) {
	email := r.AllOne("email", "")
	if email == "" {
		s.Set("lastError", "Email required")
		s.Redirect("/password-reminder")
		return
	}
	s.Delete("lastError")
	q.Dispatch("password-reminder", "user", map[string]interface{}{"email": email})
	s.Redirect("/reminder-sent")
}

func PasswordReminderSentControllerAction(v view.Viewer) string {
	return v.RenderView("password-reminder-sent.html", nil)
}

func PasswordChangeControllerAction(db db.DBer, sqlBuilder builder.Builder, r request.Requester, v view.Viewer, s session.Sessioner) string {
	uuid := r.AllOne("uuid", "")
	for {
		if uuid == "" {
			s.Set("lastError", "Invalid password reminder email")
			break
		}
		sql, err := sqlBuilder.Select("password_reminders").Fields("user_id").Where("uuid", "=", uuid).AsSQL()
		if err != nil {
			s.Set("lastError", err.Error())
			break
		}

		_, err = db.QueryOne(sql, sqlBuilder.GetParams()...)
		if err != nil {
			s.Set("lastError", "Invalid password reminder email")
			break
		}

		break
	}
	data := map[string]string{
		"lastError": s.Get("lastError"),
		"uuid":      uuid,
	}
	s.Delete("lastError")
	return v.RenderView("password-change.html", data)
}

func PasswordChangePostControllerAction(db db.DBer, sqlBuilder builder.Builder, r request.Requester, v view.Viewer, s session.Sessioner) {
	uuid := r.AllOne("uuid", "")
	redirectUrl := "/change_password?uuid=" + uuid
	password := r.AllOne("password", "")
	repeatPassword := r.AllOne("repeat-password", "")
	if len(password) < 6 {
		s.Set("lastError", "Passwords must be longer then 6 character")
		s.Redirect(redirectUrl)
		return
	}

	if password != repeatPassword {
		s.Set("lastError", "The two password does not match")
		s.Redirect(redirectUrl)
		return
	}

	sql, err := sqlBuilder.Select("password_reminders").Fields("user_id").Where("uuid", "=", uuid).AsSQL()
	if err != nil {
		s.Set("lastError", err.Error())
		s.Redirect(redirectUrl)
		return
	}

	reminderData, err := db.QueryOne(sql, sqlBuilder.GetParams()...)
	if err != nil {
		s.Set("lastError", "Invalid password reminder email")
		s.Redirect(redirectUrl)
		return
	}

	userId, ok := reminderData["user_id"]
	if !ok {
		s.Set("lastError", "cannot get user")
		s.Redirect(redirectUrl)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		s.Set("lastError", err.Error())
		s.Redirect("/error")
		return
	}

	sql, err = sqlBuilder.Update("users").Fields("password").Values(hashedPassword).Where("id", "=", userId).AsSQL()
	if err != nil {
		s.Set("lastError", err.Error())
		s.Redirect("/error")
		return
	}

	_, err = db.Execute(sql, sqlBuilder.GetParams()...)
	if err != nil {
		s.Set("lastError", err.Error())
		s.Redirect("/error")
		return
	}

	sql, err = sqlBuilder.Delete("password_reminders").Where("uuid", "=", uuid).AsSQL()
	if err != nil {
		s.Set("lastError", err.Error())
		s.Redirect("/error")
		return
	}

	_, err = db.Execute(sql, sqlBuilder.GetParams()...)
	if err != nil {
		s.Set("lastError", err.Error())
		s.Redirect("/error")
		return
	}

	s.Delete("lastError")
	s.Redirect("/login")
}
