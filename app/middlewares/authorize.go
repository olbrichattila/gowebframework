package middleware

import (
	"framework/internal/app/request"
	"framework/internal/app/session"
	"net/http"
)

func AuthorizeMiddleware(w http.ResponseWriter, r request.Requester, s session.Sessioner) bool {
	userId := s.Get("userId")
	if userId == "" {
		http.Redirect(w, r.GetRequest(), "/login", http.StatusSeeOther)
		return false
	}

	return true
}
