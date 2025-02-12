package middleware

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	ssopb "github.com/LavaJover/storage-sso-service/sso-service/proto/gen"
)

func AuthMiddleware(ssoClient ssopb.AuthServiceClient, next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Extract JWT from header
		authHeader := r.Header.Get("Authorization")
		if authHeader == ""{
			http.Error(w, "Authorization token was not found", http.StatusBadRequest)
			return
		}
		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer"{
			http.Error(w, "Wrong token scheme", http.StatusBadRequest)
			return
		}
		tokenString := authHeaderParts[1]

		// Validate JWT by requesting SSO
		response, err := ssoClient.ValidateToken(context.Background(), &ssopb.ValidateTokenRequest{
			AccessToken: tokenString,
		})

		// Process the response
		if err != nil{
			http.Error(w, "Token is on valid", http.StatusUnauthorized)
			return
		}
		userID := response.GetUserId()
		r.Header.Set("user_id", strconv.FormatUint(userID, 10))
		next.ServeHTTP(w, r)
	})
}