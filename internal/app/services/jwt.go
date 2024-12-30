package services

import (
	"craftnet/config"
	"craftnet/internal/util"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/samber/lo"
)

type JwtCustomClaim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var jwtSecret = []byte(getJwtSecret())

// GenerateJWT generates a new JWT token for a given user ID
func GenerateJWT(username string, aliveHour time.Duration) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &JwtCustomClaim{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(aliveHour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	token, err := t.SignedString(jwtSecret)
	if !lo.IsNil(err) {
		errMsg := util.ErrorMessage(util.ERROR_CODE[util.FAIL_TO_GENERATE], "token")
		util.GetLogger().LogErrorWithMsgAndError(errMsg, err, false)
		return "", err
	}

	return token, nil
}

// ValidateJWT validates the given JWT token and returns the claims if valid
func ValidateJWT(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &JwtCustomClaim{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			errMsg := util.ErrorMessage(util.ERROR_CODE[util.FAIL_TO_VALIDATE], "token")
			util.GetLogger().LogErrorWithMsg(errMsg, false)
			return nil, errors.New(errMsg)
		}
		return jwtSecret, nil
	})
}

func getJwtSecret() string {
	return config.GetAppConfig().Jwt.Secret
}
