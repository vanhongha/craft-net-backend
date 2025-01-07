package services

import (
	"craftnet/graph/model"
	"craftnet/internal/app/sql"
	"craftnet/internal/app/tools"
	craftnet_model "craftnet/internal/model"
	"craftnet/internal/util"
	"time"

	"github.com/samber/lo"
)

func Login(username string, password string) (*model.LoginResponse, *craftnet_model.Error) {
	// Get user information
	account, craftnetError := sql.GetAccountByUsername(username)
	if !lo.IsNil(craftnetError) {
		return nil, craftnetError
	}

	// If there is no user account
	if lo.IsNil(account) {
		errMsg := util.ErrorMessage(util.ERROR_CODE[util.USER_INVALID_USERNAME_PASSWORD])
		util.GetLogger().LogErrorWithMsg(errMsg, false)
		return nil, &craftnet_model.Error{
			Message: errMsg,
			Error:   nil,
		}
	}

	// If password is not correct
	if err := tools.ComparePassword(account.PasswordHash, password); !lo.IsNil(err) {
		errMsg := util.ErrorMessage(util.ERROR_CODE[util.USER_INVALID_USERNAME_PASSWORD])
		util.GetLogger().LogErrorWithMsg(errMsg, false)
		return nil, &craftnet_model.Error{
			Message: errMsg,
			Error:   nil,
		}
	}

	// Generate jwt token
	accessToken, err := GenerateJWT(account.Username, time.Hour*24)
	if !lo.IsNil(err) {
		errMsg := util.ErrorMessage(util.ERROR_CODE[util.FAIL_TO_CREATE], "access token")
		util.GetLogger().LogErrorWithMsg(errMsg, false)
		return nil, &craftnet_model.Error{
			Message: errMsg,
			Error:   nil,
		}
	}

	refreshToken, err := GenerateJWT(account.Username, time.Hour*24*30)
	if !lo.IsNil(err) {
		errMsg := util.ErrorMessage(util.ERROR_CODE[util.FAIL_TO_CREATE], "reload token")
		util.GetLogger().LogErrorWithMsg(errMsg, false)
		return nil, &craftnet_model.Error{
			Message: errMsg,
			Error:   nil,
		}
	}

	authPayload := &model.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       account.User.ID,
	}

	return authPayload, nil
}
