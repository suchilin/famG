package app

import (
	u "backend/utils"
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

/*
JWT clarms struct
*/
type Token struct {
	UserId uint
	jwt.StandardClaims
}

func generateAccessToken(sub string, userId uint) string {
	token := jwt.New(jwt.SigningMethodHS256)

	expirationTime := time.Now().Add(time.Minute * 15)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = sub
	claims["UserId"] = userId
	claims["exp"] = expirationTime.Unix()

	t, err := token.SignedString([]byte(os.Getenv("token_password")))
	if err != nil {
		return ""
	}
	return t

}

func generateRefreshToken(sub string, userId uint) string {
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = sub
	rtClaims["UserId"] = userId
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	rt, err := refreshToken.SignedString([]byte(os.Getenv("token_password")))
	if err != nil {
		return ""
	}
	return rt

}

func SetAccessCookie(sub string, userId uint, response http.ResponseWriter) {
	access_token := generateAccessToken(sub, userId)
	cookie := &http.Cookie{
		Name:   "access",
		Value:  access_token,
		Path:   "/",
		MaxAge: 3600,
	}
	http.SetCookie(response, cookie)
}

func SetRefreshCookie(sub string, userId uint, response http.ResponseWriter) {
	refresh_token := generateRefreshToken(sub, userId)
	cookie := &http.Cookie{
		Name:   "refresh",
		Value:  refresh_token,
		Path:   "/",
		MaxAge: 5 * 24 * 15,
	}
	http.SetCookie(response, cookie)
}

func ClearSession(response http.ResponseWriter) {
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

		notAuth := []string{"/api/v1/auth/signup", "/api/v1/auth/login"}
		requestPath := r.URL.Path
		for _, value := range notAuth {

			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		cookie, _ := r.Cookie("access")

		// response := make(map[string]interface{})
		// tokenHeader := r.Header.Get("Authorization")
		// if tokenHeader == "" {
		//         response = u.Message(422, "Missing auth token")
		//         u.Respond(w, response)
		//         return
		// }
		//
		// splitted := strings.Split(tokenHeader, " ")
		// if len(splitted) != 2 {
		//         response = u.Message(422, "Invalid/Malformed auth token")
		//         u.Respond(w, response)
		//         return
		// }
		//
		tk := &Token{}

		token, err := jwt.ParseWithClaims(cookie.Value, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if token.Valid {
			ctx := context.WithValue(r.Context(), "user", tk.UserId)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				response := u.Message(http.StatusBadRequest, "That's not even a token")
				u.Respond(w, response)
				return
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				fmt.Println("REFRESHING TOKEN")
				// Token is either expired or not active yet
				rcookie, _ := r.Cookie("refresh")
				rtoken, _ := jwt.ParseWithClaims(rcookie.Value, tk, func(token *jwt.Token) (interface{}, error) {
					return []byte(os.Getenv("token_password")), nil
				})
				if rtoken.Valid {
					SetAccessCookie(tk.Subject, tk.UserId, w)
					ctx := context.WithValue(r.Context(), "user", tk.UserId)
					r = r.WithContext(ctx)
					next.ServeHTTP(w, r)
					return
				} else {
					response := u.Message(http.StatusUnauthorized, "Token Expired")
					u.Respond(w, response)
					return
				}

			} else {
				response := u.Message(http.StatusBadRequest, "Couldn't handle this token")
				u.Respond(w, response)
				return
			}
		} else {
			response := u.Message(http.StatusBadRequest, "Couldn't handle this token")
			u.Respond(w, response)
			return
		}

	})
}
