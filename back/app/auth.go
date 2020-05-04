package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	u "backend/utils"
)

/*
JWT clarms struct
*/
type Token struct {
	UserId uint
	jwt.StandardClaims
}

func generateAccessToken(sub uint) string {
	token := jwt.New(jwt.SigningMethodHS256)

	expirationTime := time.Now().Add(time.Minute * 15)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = sub
	claims["exp"] = expirationTime.Unix()

	t, err := token.SignedString([]byte(os.Getenv("token_password")))
	if err != nil {
		return ""
	}
	return t

}

func generateRefreshToken(sub uint) string {
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = sub
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	rt, err := refreshToken.SignedString([]byte(os.Getenv("token_password")))
	if err != nil {
		return ""
	}
	return rt

}

func SetAccessCookie(sub uint, response http.ResponseWriter) {
	access_token := generateAccessToken(sub)
	cookie := &http.Cookie{
		Name:   "access",
		Value:  access_token,
		Path:   "/",
		MaxAge: 3600,
	}
	http.SetCookie(response, cookie)
}

func SetRefreshCookie(sub uint, response http.ResponseWriter) {
	refresh_token := generateRefreshToken(sub)
	cookie := &http.Cookie{
		Name:   "refresh",
		Value:  refresh_token,
		Path:   "/",
		MaxAge: 5 * 24 * 15,
	}
	http.SetCookie(response, cookie)
}

func ClearSession(response http.ResponseWriter){
	acookie := &http.Cookie{
		Name:   "access",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	cookie := &http.Cookie{
		Name:   "refresh",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, acookie)
	http.SetCookie(response, cookie)

}

var JwtAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		notAuth := []string{"/api/v1/auth/signup", "/api/v1/auth/token"}
		requestPath := r.URL.Path
		for _, value := range notAuth {

			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		cookie, err:=r.Cookie("access")
		fmt.Println("COOKIE VALUE", cookie.Value)

		response := make(map[string]interface{})
		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" {
			response = u.Message(422, "Missing auth token")
			u.Respond(w, response)
			return
		}

		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			response = u.Message(422, "Invalid/Malformed auth token")
			u.Respond(w, response)
			return
		}

		tokenPart := splitted[1]
		tk := &Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil {
			response = u.Message(422, "Malformed authentication token")
			u.Respond(w, response)
			return
		}

		if !token.Valid {
			response = u.Message(422, "Token is not valid.")
			u.Respond(w, response)
			return
		}

		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
