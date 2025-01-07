package handlers

import (
	"craftnet/graph/model"
	"craftnet/internal/app/services"
	"craftnet/internal/util"
	"errors"

	"github.com/samber/lo"
)

func RegisterAccountHandler(username string, password string) (*model.RegisterResponse, error) {
	account, err := services.RegisterAccount(username, password)
	if !lo.IsNil(err) {
		return nil, errors.New(err.Message)
	}

	if lo.IsNil(account) {
		errMsg := util.ErrorMessage(util.ERROR_CODE[util.FAIL_TO_CREATE], "account")
		util.GetLogger().LogErrorWithMsg(errMsg, false)
		return nil, errors.New(errMsg)
	}

	response := &model.RegisterResponse{
		UserID:    account.User.ID,
		AccountID: account.ID,
		Username:  account.Username,
		Message:   util.NotiMessage(util.NOTI_CODE[util.NOTI_CREATE_ACCOUNT_SUCCESSFULLY]),
	}

	return response, nil
}

func Login(username string, password string) (*model.LoginResponse, error) {
	authPlayload, err := services.Login(username, password)
	if !lo.IsNil(err) {
		return nil, errors.New(err.Message)
	}

	return authPlayload, nil
}
