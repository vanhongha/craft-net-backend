package sql

import (
	"craftnet/graph/model"
	"craftnet/internal/db"
	craftnet_model "craftnet/internal/model"
	"craftnet/internal/util"
	"database/sql"
	"errors"

	"github.com/samber/lo"
)

// GetAccountByUsername retrieves account details for a given username.
// It takes a context for request scoping and the username as input parameters.
// Returns the account details or an error if the account is not found.
func GetAccountByUsername(username string) (*model.Account, *craftnet_model.Error) {
	var account = &model.Account{
		User: &model.User{},
	}
	if err := db.Instance.QueryRow("SELECT id, user_id, username FROM accounts WHERE username = ?", username).Scan(&account.ID, &account.User.ID, &account.Username); !lo.IsNil(err) {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		errMsg := util.ErrorMessage(util.ERROR_CODE[util.INTERNAL_SERVER], nil)
		util.GetLogger().LogErrorWithMsgAndError(errMsg, err, false)
		return nil, &craftnet_model.Error{
			Message: errMsg,
			Error:   err,
		}
	}

	return account, nil
}

// CreateAccount creates a new account with the specified username and password hash.
// It takes a context for request scoping, the username, and a hashed password as input parameters.
// Returns the newly created account model on success or an error if the creation fails.
func CreateAccount(username string, passwordHash string) (*model.Account, *craftnet_model.Error) {
	// Begin a new database transaction using the provided context.
	// 'tx' represents the transaction object, and 'err' captures any error encountered during initialization.
	tx, err := db.Instance.Begin()
	if !lo.IsNil(err) {
		errMsg := util.ErrorMessage(util.ERROR_CODE[util.INTERNAL_SERVER], nil)
		util.GetLogger().LogErrorWithMsgAndError(errMsg, err, false)
		return nil, &craftnet_model.Error{
			Message: errMsg,
			Error:   err,
		}
	}
	defer tx.Rollback()

	/*
	 **************************************************
	 *               CREATE USER                      *
	 **************************************************
	 */
	query := `
	    INSERT INTO users (last_name, first_name, date_of_birth, email, phone_number, created_at, updated_at, status)
	    VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
	var user = &model.User{
		ID:          0,
		FirstName:   "Temp first name",
		LastName:    "Temp last name",
		DateOfBirth: "20000101",
		Email:       "example@mail.com",
		PhoneNumber: "0000000000",
		Status:      util.STATUS_IDLE,
	}

	res, err := tx.Exec(
		query,
		user.LastName,
		user.FirstName,
		user.DateOfBirth,
		user.Email,
		user.PhoneNumber,
		util.TimeNowJSTZone(),
		util.TimeNowJSTZone(),
		user.Status)
	if !lo.IsNil(err) {
		errMsg := util.ErrorMessage(util.ERROR_CODE[util.FAIL_TO_CREATE], "user")
		util.GetLogger().LogErrorWithMsgAndError(errMsg, err, false)
		return nil, &craftnet_model.Error{
			Message: errMsg,
			Error:   err,
		}
	}

	// Get last inserted user id
	userId, err := res.LastInsertId()
	user.ID = int(userId)
	if !lo.IsNil(err) {
		errMsg := util.ErrorMessage(util.ERROR_CODE[util.FAIL_TO_CREATE], "user")
		util.GetLogger().LogErrorWithMsgAndError(errMsg, err, false)
		return nil, &craftnet_model.Error{
			Message: errMsg,
			Error:   err,
		}
	}
	infoMsg := util.InfoMessage(util.INFO_CODE[util.INFO_INSERTED_USER], userId)
	util.GetLogger().LogInfo(infoMsg)

	/*
	 **************************************************
	 *               CREATE ACCOUNT                   *
	 **************************************************
	 */
	query = `
	    INSERT INTO accounts (user_id, username, password_hash, created_at, updated_at)
	    VALUES (?, ?, ?, ?, ?)
	`

	res, err = tx.Exec(
		query,
		user.ID,
		username,
		passwordHash,
		util.TimeNowJSTZone(),
		util.TimeNowJSTZone())
	if !lo.IsNil(err) {
		errMsg := util.ErrorMessage(util.ERROR_CODE[util.FAIL_TO_CREATE], "account")
		util.GetLogger().LogErrorWithMsgAndError(errMsg, err, false)
		return nil, &craftnet_model.Error{
			Message: errMsg,
			Error:   err,
		}
	}

	// Get last inserted account id
	accountId, err := res.LastInsertId()
	if !lo.IsNil(err) {
		errMsg := util.ErrorMessage(util.ERROR_CODE[util.FAIL_TO_CREATE], "account")
		util.GetLogger().LogErrorWithMsgAndError(errMsg, err, false)
		return nil, &craftnet_model.Error{
			Message: errMsg,
			Error:   err,
		}
	}
	infoMsg = util.InfoMessage(util.INFO_CODE[util.INFO_INSERTED_ACCOUNT], accountId)
	util.GetLogger().LogInfo(infoMsg)

	// Commit the transaction.
	if err := tx.Commit(); !lo.IsNil(err) {
		errMsg := util.ErrorMessage(util.ERROR_CODE[util.FAIL_TO_CREATE], "account")
		util.GetLogger().LogErrorWithMsgAndError(errMsg, err, false)
		return nil, &craftnet_model.Error{
			Message: errMsg,
			Error:   err,
		}
	}

	infoMsg = util.InfoMessage(util.INFO_CODE[util.INFO_ACTION_SUCCESFULLY], "CreateAccount")
	util.GetLogger().LogInfo(infoMsg)

	account := &model.Account{
		ID:           int(accountId),
		Username:     username,
		PasswordHash: passwordHash,
		User:         user,
	}
	return account, nil
}
