package controller

import (
	"fmt"
	"framework/internal/app/db"
	"framework/internal/app/logger"
	"framework/internal/app/request"
	"framework/internal/app/session"
	"framework/internal/app/view"

	builder "github.com/olbrichattila/gosqlbuilder/pkg"
	"golang.org/x/crypto/bcrypt"
)

func Login(v view.Viewer, s session.Sessioner) string {
	data := map[string]string{
		"email":     s.Get("loginUserId"),
		"lastError": s.Get("lastError"),
	}
	s.Delete("lastError")

	return v.RenderView("login.html", data)
}

func LoginPost(r request.Requester, db db.DBer, sqlBuilder builder.Builder, s session.Sessioner, l logger.Logger) {
	defer db.Close()

	email := r.AllOne("email", "")
	password := r.AllOne("password", "")

	s.Set("loginUserId", email)
	s.Set("userId", email)

	sqlBuilder.Select("users").Fields("id", "password").Where("email", "=", email).IsNotNull("activated_at")
	sql, err := sqlBuilder.AsSQL()
	if err != nil {
		s.Set("lastError", err.Error())
		s.Redirect("/error")
		return
	}

	params := sqlBuilder.GetParams()
	res, err := db.QueryOne(sql, params...)
	if err != nil {
		s.Set("lastError", "User or password incorrect")
		s.Redirect("/login")
		return
	}

	if dbHashedPassword, ok := res["password"]; ok {
		err := bcrypt.CompareHashAndPassword([]byte(dbHashedPassword.(string)), []byte(password))
		if err != nil {
			l.Info(fmt.Sprintf("User %s provides incorrect password", email))
			s.Set("lastError", "User or password incorrect")
			s.Redirect("/login")
			return
		}
	}

	l.Info(fmt.Sprintf("User %s logged in", email))

	s.Delete("lastError")
	s.Redirect("/")
}

func Logout(s session.Sessioner) {
	s.Close()
	s.RemoveSession()
	s.Redirect("/login")
}
