package http

import (
	"context"
	"net/http"
	"strings"
)

type key int

const (
	keySession key = iota
)

func (h *Handler) authorizationHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("authorization")
		s := strings.Split(authorizationHeader, " ")
		var sessionID string
		if len(s) == 2 {
			sessionID = s[1]
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		session, err := h.SessionService.GetSession(sessionID)

		if err != nil || session == nil { // if no session or there is an error
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if session.Status != "activated" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), keySession, session)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
