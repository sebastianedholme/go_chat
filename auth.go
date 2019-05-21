package main

import (
	"log"
	"net/http"
	"strings"
)

type authHandler struct {
	next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("auth")
	if err == http.ErrNoCookie {
		// not authenticated
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}

	if err != nil {
		// some error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// success
	h.next.ServeHTTP(w, r)
}

func MustAuth(handler http.Handler) http.Handler {
	return &authHandler{next: handler}
}

// loginHandler handles the third-party login process.
// format: /auth/{action}/{provider}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	// split segments
	segs := strings.Split(r.URL.Path, "/")

	// secure segement length
	if len(segs) <= 3 {
		w.WriteHeader(http.StatusNotFound)
		log.Fprintf(w, "This is not permitted")
		return
	}

	action := segs[2]
	provider := segs[3]
	switch action {
	case "login":
		log.Println("TODO handle login for", provider)
	default:
		w.WriteHeader(http.StatusNotFound)
		log.Fprintf(w, "Auth actgion %s not supported", action)
	}
}
