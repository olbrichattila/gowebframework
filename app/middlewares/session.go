package middleware

import (
	"framework/internal/app/request"
	"framework/internal/app/session"
	"net/http"
)

func SessionMiddleware(w http.ResponseWriter, r request.Requester, s session.Sessioner) {
	s.Init(w, r.GetRequest())
}
