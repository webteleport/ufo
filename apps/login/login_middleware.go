package login

import (
	"fmt"
	"net/http"
	"strings"
)

type LoginMiddleware struct {
	Password string
}

func (lm *LoginMiddleware) RedirectToLogin(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, "/login/") {
		http.Redirect(w, r, "/login/", 302)
		return
	}
	login := http.StripPrefix("/login", http.FileServer(http.Dir("apps/login/www")))
	login.ServeHTTP(w, r)
}

func (lm *LoginMiddleware) SetCookiesAndRedirect(w http.ResponseWriter, r *http.Request) {
	cookies := fmt.Sprintf(`WebTeleportSecretCode="%s"; Path=/; Max-Age=2592000; HttpOnly; Domain=%s`, lm.Password, r.Host)
	w.Header().Set("Set-Cookie", cookies)
	http.Redirect(w, r, "/", 302)
}

func (lm *LoginMiddleware) ValidateRequest(r *http.Request) bool {
	if r.Method == http.MethodPost {
		pw := r.PostFormValue("password")
		if pw == lm.Password {
			return true
		}
	}
	return r.URL.Path == fmt.Sprintf("/login/%s", lm.Password)
}

// PrecheckAccessToken returns a bool that indicates whether the caller should continue
func (lm *LoginMiddleware) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hostonly, _, _ := strings.Cut(r.URL.Host, ":")
		if lm.Password == "" || strings.HasSuffix(hostonly, "localhost") {
			next.ServeHTTP(w, r)
			return
		}
		if lm.ValidateRequest(r) {
			lm.SetCookiesAndRedirect(w, r)
			return
		}
		wtat, err := r.Cookie("WebTeleportSecretCode")
		if err != nil {
			lm.RedirectToLogin(w, r)
			return
		}
		if wtat.Value != lm.Password {
			lm.RedirectToLogin(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}
