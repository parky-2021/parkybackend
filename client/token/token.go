package token

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	ApiSecret = "parky2021"
)

type AuthDetails struct {
	UserID string
}

func CreateToken(authD AuthDetails) (string, error) {
	expirationTime := time.Now().Add(60 * 24 * 60 * time.Minute)
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["auth_user_id"] = authD.UserID
	claims["ExpiresAt"] = int64(expirationTime.Unix())
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(ApiSecret))
}

func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}
func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//does this token conform to "SigningMethodHMAC" ?
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(ApiSecret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if ok && token.Valid {
		exptime, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["ExpiresAt"]), 10, 64)
		if err != nil {
			return nil, err
		}

		t := time.Unix(exptime, 0)
		remainder := t.Sub(time.Now())
		if remainder <= 0 {
			return nil, http.ErrHandlerTimeout
		}

		_, ok := claims["auth_user_id"].(string) //convert the interface to string
		if !ok {
			return nil, http.ErrHandlerTimeout
		}

	}
	return token, nil
}

//get the token from the request body
func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func ExtractTokenAuth(r *http.Request) (*AuthDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if ok && token.Valid {
		authUuid, ok := claims["auth_user_id"].(string) //convert the interface to string
		if !ok {
			return nil, err
		}

		return &AuthDetails{
			UserID: authUuid,
		}, nil
	}
	return nil, err
}
