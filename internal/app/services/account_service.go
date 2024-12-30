package services

import (
	"craftnet/graph/model"
	"craftnet/internal/app/sql"
	craftnet_model "craftnet/internal/model"
	"craftnet/internal/util"

	"github.com/samber/lo"
	"golang.org/x/crypto/bcrypt"
)

func RegisterAccount(username string, password string) (*model.Account, *craftnet_model.Error) {
	existingAccount, craftnetError := sql.GetAccountByUsername(username)
	if !lo.IsNil(craftnetError) {
		return nil, craftnetError
	}

	if !lo.IsNil(existingAccount) {
		errMsg := util.ErrorMessage(util.ERROR_CODE[util.USER_ALREADY_EXISTS], username)
		util.GetLogger().LogErrorWithMsg(errMsg, false)
		return nil, &craftnet_model.Error{
			Message: errMsg,
			Error:   nil,
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if !lo.IsNil(err) {
		errMsg := util.ErrorMessage(util.ERROR_CODE[util.FAIL_TO_HASH_PASSWORD])
		util.GetLogger().LogErrorWithMsgAndError(errMsg, err, false)
		return nil, &craftnet_model.Error{
			Message: errMsg,
			Error:   nil,
		}
	}

	newAccount, craftnetError := sql.CreateAccount(username, string(hashedPassword))
	if !lo.IsNil(craftnetError) {
		return nil, craftnetError
	}

	return newAccount, nil
}
