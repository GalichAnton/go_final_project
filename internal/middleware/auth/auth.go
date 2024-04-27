package auth

import (
	"net/http"

	"github.com/GalichAnton/go_final_project/internal/utils"
)

func Auth(next http.HandlerFunc, pass string) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if len(pass) > 0 {
				cookie, err := r.Cookie("token")
				if err != nil {
					http.Error(w, "Authentification required", http.StatusUnauthorized)
					return
				}

				hash := utils.CreateHash(pass)

				if cookie.Value != hash {
					http.Error(w, "Authentification required", http.StatusUnauthorized)
					return
				}
			}

			next(w, r)
		},
	)
}
