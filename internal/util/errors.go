package util

import (
	"fmt"
)

type AppError struct {
	ErrorObject ErrorObject
	Error       error
}

type ErrorObject struct {
	Code    string
	Message string
}

func ErrorMessage(code string, obj ...any) string {
	if len(obj) == 0 || obj[0] == nil {
		return ERROR_MESSAGE[code]
	}
	return fmt.Sprintf(ERROR_MESSAGE[code], obj...)
}

func InfoMessage(code string, obj ...any) string {
	if len(obj) == 0 || obj[0] == nil {
		return INFO_MESSAGE[code]
	}
	return fmt.Sprintf(INFO_MESSAGE[code], obj...)
}

var ERROR_MESSAGE = map[string]string{
	// App errors
	ERROR_CODE[INTERNAL_SERVER]: "Internal server error",

	// User errors
	ERROR_CODE[USER_ALREADY_EXISTS]: "User %s already exists",
	ERROR_CODE[USER_NOT_FOUND]:      "User not found",

	// Fail errors
	ERROR_CODE[FAIL_TO_HASH_PASSWORD]: "Fail to hash password",
	ERROR_CODE[FAIL_TO_CREATE]:        "Fail to create %s",
}

var INFO_MESSAGE = map[string]string{
	INFO_CODE[INFO_INSERTED_USER]:      "Inserted user with id %d",
	INFO_CODE[INFO_INSERTED_ACCOUNT]:   "Inserted account with id %d",
	INFO_CODE[INFO_ACTION_SUCCESFULLY]: "%s occured succesfully",
}

var ERROR_CODE = map[string]string{
	// App errors
	INTERNAL_SERVER: "A1001",

	// User errors
	USER_ALREADY_EXISTS: "DU1001",
	USER_NOT_FOUND:      "DU1002",

	// Fail errors
	FAIL_TO_HASH_PASSWORD: "DF1001",
	FAIL_TO_CREATE:        "DF1002",
}

var INFO_CODE = map[string]string{
	INFO_ACTION_SUCCESFULLY: "IF1001",
	INFO_INSERTED_USER:      "IF1002",
	INFO_INSERTED_ACCOUNT:   "IF1003",
}

const (
	// App errors
	INTERNAL_SERVER = "INTERNAL_SERVER"

	// User errors
	USER_ALREADY_EXISTS = "USER_ALREADY_EXISTS"
	USER_NOT_FOUND      = "USER_NOT_FOUND"

	// Fail errors
	FAIL_TO_HASH_PASSWORD = "FAIL_TO_HASH_PASSWORD"
	FAIL_TO_CREATE        = "FAIL_TO_CREATE"

	// Info
	INFO_ACTION_SUCCESFULLY = "INFO_ACTION_SUCCESFULLY"
	INFO_INSERTED_USER      = "INFO_INSERTED_USER"
	INFO_INSERTED_ACCOUNT   = "INFO_INSERTED_ACCOUNT"
)
