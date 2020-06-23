package Auth

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"

	"github.com/GhvstCode/Notes-API/models"
	"github.com/GhvstCode/Notes-API/utils"
)

 var JWT = func(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		openRoutes := []string{"/api/user/new", "/api/user/login"}
		requestPath := r.URL.Path

		for _, value := range openRoutes {

			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		res := make(map[string]interface{})
		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			res = utils.Message(false, "Missing auth token", http.StatusUnauthorized)
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.RespondJson(w, res)
			return
		}

		tArray := strings.Split(tokenHeader, " ")
		if len(tArray) != 2 {
			res = utils.Message(false, "Invalid/Malformed auth token", http.StatusBadRequest)
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.RespondJson(w, res)
			return
		}

		tokenPart := tArray[1] //Grab the token part, what we are truly interested in
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil { //Malformed token, returns with http code 403 as usual
			res = utils.Message(false, "Malformed authentication token", http.StatusBadRequest)
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.RespondJson(w, res)
			return
		}

		if !token.Valid { //Token is invalid, maybe not signed on this server
			res = utils.Message(false, "Token is not valid.", http.StatusBadRequest)
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.RespondJson(w, res)
			return
		}

		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}