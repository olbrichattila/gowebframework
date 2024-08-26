package controller

import (
	"framework/internal/app/db"
	"framework/internal/app/queue"
	"framework/internal/app/request"
	"framework/internal/app/session"
	"framework/internal/app/view"
	"strings"
	"time"

	"github.com/google/uuid"
	builder "github.com/olbrichattila/gosqlbuilder/pkg"
	"golang.org/x/crypto/bcrypt"
)

func Register(v view.Viewer, s session.Sessioner) string {
	data := map[string]string{
		"regUserEmail": s.Get("regUserEmail"),
		"regUserName":  s.Get("regUserName"),
		"lastError":    s.Get("lastError"),
	}

	s.Delete("lastError")

	return v.Render("register.html", data)
}

func PostRegister(r request.Requester, db db.DBer, sqlBuilder builder.Builder, s session.Sessioner, q queue.Quer) {
	defer db.Close()

	validation := make([]string, 0)
	name := r.AllOne("name", "")
	email := r.AllOne("email", "")
	password := r.AllOne("password", "")
	repeatPassword := r.AllOne("repeat-password", "")

	s.Set("regUserEmail", email)
	s.Set("regUserName", name)

	if name == "" {
		validation = append(validation, "Name is required")
	}

	if email == "" {
		validation = append(validation, "Email is required")
	}

	if len(password) < 6 {
		validation = append(validation, "Password should be minimum 6 character")
	}

	if password != repeatPassword {
		validation = append(validation, "Two password does not match")
	}

	if len(validation) > 0 {
		s.Set("lastError", strings.Join(validation, "<br >"))
		s.Redirect("/register")
		return
	}

	sqlBuilder.Select("users").Fields("id").Where("email", "=", email)
	sql, err := sqlBuilder.AsSQL()
	if err != nil {
		s.Set("lastError", err.Error())
		s.Redirect("/error")
		return
	}

	params := sqlBuilder.GetParams()
	_, err = db.QueryOne(sql, params...)
	if err == nil {
		s.Set("lastError", "User already exists")
		s.Redirect("/register")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		s.Set("lastError", err.Error())
		s.Redirect("/error")
		return
	}

	sqlBuilder.Insert("users").Fields("name", "email", "password").Values(name, email, hashedPassword)
	sql, err = sqlBuilder.AsSQL()
	if err != nil {
		s.Set("lastError", err.Error())
		s.Redirect("/error")
		return
	}

	params = sqlBuilder.GetParams()

	userId, err := db.Execute(sql, params...)
	if err != nil {
		s.Set("lastError", err.Error())
		s.Redirect("/error")
		return
	}

	uuid := uuid.New().String()
	createdAt := time.Now().Format("2006-01-02 15:04:05")
	sqlBuilder.Insert("reg_confirmations").Fields("user_id", "uuid", "created_at").Values(userId, uuid, createdAt)
	sql, err = sqlBuilder.AsSQL()
	if err != nil {
		s.Set("lastError", err.Error())
		s.Redirect("/error")
		return
	}

	params = sqlBuilder.GetParams()
	regId, err := db.Execute(sql, params...)
	if err != nil {
		s.Set("lastError", err.Error())
		s.Redirect("/error")
		return
	}

	q.Dispatch("register", "register-user", map[string]interface{}{"email": email, "name": name, "uuid": uuid, "id": regId})
	s.Delete("lastError")
	s.Redirect("/login")
}

func ActivateAction(r request.Requester, db db.DBer, sqlBuilder builder.Builder, s session.Sessioner) {
	defer db.Close()
	uuid := r.GetOne("uuid", "")
	if uuid == "" {
		s.Set("lastError", "Invalid confirmation link")
		s.Redirect("/error")
		return
	}

	sql, err := sqlBuilder.Select("reg_confirmations").Where("uuid", "=", uuid).AsSQL()
	if err != nil {
		s.Set("lastError", err.Error())
		s.Redirect("/error")
		return
	}

	params := sqlBuilder.GetParams()
	result, err := db.QueryOne(sql, params...)
	if err != nil {
		s.Set("lastError", "Invalid or expired confirmation link "+err.Error())
		s.Redirect("/error")
		return
	}

	if userId, ok := result["user_id"]; ok {
		updateSql, err := sqlBuilder.Update("users").
			Fields("activated_at").
			Values(time.Now().Format("2006-01-02 15:04:05")).
			Where("id", "=", userId).AsSQL()
		if err != nil {
			s.Set("lastError", err.Error())
			s.Redirect("/error")
			return
		}

		params = sqlBuilder.GetParams()
		_, err = db.Execute(updateSql, params...)
		if err != nil {
			s.Set("lastError", err.Error())
			s.Redirect("/error")
			return
		}

		deleteSql, err := sqlBuilder.Delete("reg_confirmations").Where("uuid", "=", uuid).AsSQL()
		if err != nil {
			s.Set("lastError", err.Error())
			s.Redirect("/error")
			return
		}

		params = sqlBuilder.GetParams()
		_, err = db.Execute(deleteSql, params...)
		if err != nil {
			s.Set("lastError", err.Error())
			s.Redirect("/error")
			return
		}

		s.Delete("lastError")
		s.Redirect("/login")
		return
	}

	s.Set("lastError", "Unknown error")
	s.Redirect("/error")
}
