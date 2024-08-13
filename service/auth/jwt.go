package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/itsmoirob/ecom-auth/config"
	"github.com/itsmoirob/ecom-auth/types"
	"github.com/itsmoirob/ecom-auth/utils"
)

type contextKey string

const userIDKey contextKey = "userID"

func CreateJWT(secret []byte, userID int) (string, error) {
	expiration := time.Second * time.Duration(config.Envs.JWTExpirationInSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(userID),
		"expiredAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func WithJWTAuth(handlerFunc http.HandlerFunc, store types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// does token exist?
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("missing Authorization header"))
			return
		}

		// is token valid?
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Envs.JWTSecret), nil
		})
		if err != nil {
			log.Printf("failed to validate token: %v", err)
			permisionDenied(w)
			return
		}

		// get user id from the token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			log.Printf("failed to parse token: %v", err)
			permisionDenied(w)
			return
		}

		userID, err := strconv.Atoi(claims["userID"].(string))
		if err != nil {
			log.Printf("failed to get userId: %v", err)
			permisionDenied(w)
			return
		}
		user, err := store.GetUserById(userID)
		if err != nil {
			log.Printf("user not in database: %v", err)
			permisionDenied(w)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), userIDKey, user.ID))
		handlerFunc(w, r)
	}
}

func permisionDenied(w http.ResponseWriter) {
	utils.WriteError(w, http.StatusForbidden, fmt.Errorf("permission denied"))
}

func GetUserIdFromContext(ctx context.Context) int {
	userID := ctx.Value(userIDKey)
	if userID == nil {
		return 0
	}
	return userID.(int)
}
