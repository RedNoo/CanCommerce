package middleware

import (
	"errors"
	"fmt"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/valyala/fasthttp"
)

var mySigningKey = []byte("AllYourBase")

type UserCredential struct {
	UserID   string `json:"userID"`
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func CreateToken(userID string, username string, password string, day int) (string, time.Time) {

	expireAt := time.Now().Add(time.Hour * 24 * time.Duration(day))

	claims := UserCredential{
		UserID:   userID,
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireAt.Unix(), // 30 days
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Println(err)
	}

	return ss, expireAt
}

func ParseToken(requestToken string) (*jwt.Token, *UserCredential, error) {

	token, err := jwt.ParseWithClaims(requestToken, &UserCredential{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if claims, ok := token.Claims.(*UserCredential); ok && token.Valid {
		fmt.Printf("%v %v", claims.Username, claims.StandardClaims.ExpiresAt)
		return token, claims, err
	} else {
		fmt.Println(err)
		return token, claims, err
	}

}

func JwtMiddleware(ctx *fasthttp.RequestCtx) error {
	jwtcookie := ctx.Request.Header.Cookie("jwtcookie")

	if len(jwtcookie) == 0 {
		return errors.New("login required")
	}

	token, _, err := ParseToken(string(jwtcookie))

	if !token.Valid {
		return errors.New("your session is expired, login again please")
	}

	return err
}

// https://godoc.org/github.com/valyala/fasthttp#RequestHandler
// BasicAuth is the basic auth handler
func Middleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {

		if err := JwtMiddleware(ctx); err != nil {
			fmt.Fprint(ctx, err)
			return
		}

		next(ctx)
	})
}
