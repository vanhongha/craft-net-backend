package middleware

import (
	"context"
	"craftnet/internal/app/services"
	"craftnet/internal/util"
	"net/http"
	"time"

	"github.com/samber/lo"
)

type authString string

// AuthMiddleware validates JWT tokens and adds user context
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")

		if auth == "" {
			next.ServeHTTP(w, r)
			return
		}

		bearer := "Bearer "
		auth = auth[len(bearer):]

		validate, err := services.ValidateJWT(auth)
		if !lo.IsNil(err) || !validate.Valid {
			errMsg := util.ErrorMessage(util.ERROR_CODE[util.CANNOT_VALIDATE_TOKEN])
			util.GetLogger().LogErrorWithMsgAndError(errMsg, err, false)
			http.Error(w, "Invalid token", http.StatusForbidden)
			return
		}

		customClaim, _ := validate.Claims.(*services.JwtCustomClaim)

		if customClaim.ExpiresAt != 0 && customClaim.ExpiresAt < time.Now().Unix() {
			errMsg := util.ErrorMessage(util.ERROR_CODE[util.CANNOT_VALIDATE_TOKEN])
			util.GetLogger().LogErrorWithMsgAndError(errMsg, err, false)
			http.Error(w, "Token is expired", http.StatusForbidden)
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")          // Allowed methods
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type") // Include Authorization
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Vary", "Origin, Access-Control-Request-Method, Access-Control-Request-Headers")

		// Handle preflight request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		ctx := context.WithValue(r.Context(), authString("auth"), customClaim)

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)

		infoMsg := util.InfoMessage(util.INFO_CODE[util.INFO_TOKEN_VALIDATED_SUCCESSFULLY])
		util.GetLogger().LogInfo(infoMsg)
	})
}

func CtxValue(ctx context.Context) *services.JwtCustomClaim {
	raw, _ := ctx.Value(authString("auth")).(*services.JwtCustomClaim)
	return raw
}
