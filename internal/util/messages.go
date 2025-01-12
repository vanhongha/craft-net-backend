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

func NotiMessage(code string, obj ...any) string {
	if len(obj) == 0 || obj[0] == nil {
		return NOTI_MESSAGE[code]
	}
	return fmt.Sprintf(NOTI_MESSAGE[code], obj...)
}

var ERROR_MESSAGE = map[string]string{
	// App errors
	ERROR_CODE[INTERNAL_SERVER]:       "Internal server error",
	ERROR_CODE[CANNOT_VALIDATE_TOKEN]: "Invalid token",

	// User errors
	ERROR_CODE[USER_ALREADY_EXISTS]:            "User %s already exists",
	ERROR_CODE[USER_NOT_FOUND]:                 "User not found",
	ERROR_CODE[USER_INVALID_USERNAME_PASSWORD]: "Invalid username or password",

	// Fail errors
	ERROR_CODE[FAIL_TO_HASH_PASSWORD]: "Fail to hash password",
	ERROR_CODE[FAIL_TO_CREATE]:        "Fail to create %s",
	ERROR_CODE[FAIL_TO_VALIDATE]:      "Fail to validate %s",
	ERROR_CODE[FAIL_TO_GENERATE]:      "Fail to generate %s",
	ERROR_CODE[FAIL_TO_GET]:           "Fail to get %s",
}

var INFO_MESSAGE = map[string]string{
	INFO_CODE[INFO_INSERTED_USER]:                "Inserted user with id %d",
	INFO_CODE[INFO_INSERTED_ACCOUNT]:             "Inserted account with id %d",
	INFO_CODE[INFO_TOKEN_VALIDATED_SUCCESSFULLY]: "Token validated successfully",
	INFO_CODE[INFO_ACTION_SUCCESFULLY]:           "%s occured succesfully",
}

var NOTI_MESSAGE = map[string]string{
	NOTI_CODE[NOTI_CREATE_ACCOUNT_SUCCESSFULLY]: "Create account succesfully",
}

var ERROR_CODE = map[string]string{
	// App errors
	INTERNAL_SERVER:       "A1001",
	CANNOT_VALIDATE_TOKEN: "A1002",

	// User errors
	USER_ALREADY_EXISTS:            "DU1001",
	USER_NOT_FOUND:                 "DU1002",
	USER_INVALID_USERNAME_PASSWORD: "DU1003",

	// Fail errors
	FAIL_TO_HASH_PASSWORD: "DF1001",
	FAIL_TO_CREATE:        "DF1002",
	FAIL_TO_VALIDATE:      "DF1003",
	FAIL_TO_GENERATE:      "DF1004",
	FAIL_TO_GET:           "DF1005",
}

var INFO_CODE = map[string]string{
	INFO_ACTION_SUCCESFULLY:           "IF1001",
	INFO_INSERTED_USER:                "IF1002",
	INFO_INSERTED_ACCOUNT:             "IF1003",
	INFO_TOKEN_VALIDATED_SUCCESSFULLY: "IF1004",
}

var NOTI_CODE = map[string]string{
	NOTI_CREATE_ACCOUNT_SUCCESSFULLY: "NT1001",
}

const (
	// App errors
	INTERNAL_SERVER       = "INTERNAL_SERVER"
	CANNOT_VALIDATE_TOKEN = "CANNOT_VALIDATE_TOKEN"

	// User errors
	USER_ALREADY_EXISTS            = "USER_ALREADY_EXISTS"
	USER_NOT_FOUND                 = "USER_NOT_FOUND"
	USER_INVALID_USERNAME_PASSWORD = "USER_INVALID_USERNAME_PASSWORD"

	// Fail errors
	FAIL_TO_HASH_PASSWORD = "FAIL_TO_HASH_PASSWORD"
	FAIL_TO_CREATE        = "FAIL_TO_CREATE"
	FAIL_TO_VALIDATE      = "FAIL_TO_VALIDATE"
	FAIL_TO_GENERATE      = "FAIL_TO_GENERATE"
	FAIL_TO_GET           = "FAIL_TO_GET"

	// Noti
	NOTI_CREATE_ACCOUNT_SUCCESSFULLY = "NOTI_CREATE_ACCOUNT_SUCCESSFULLY"

	// Info
	INFO_ACTION_SUCCESFULLY           = "INFO_ACTION_SUCCESFULLY"
	INFO_INSERTED_USER                = "INFO_INSERTED_USER"
	INFO_INSERTED_ACCOUNT             = "INFO_INSERTED_ACCOUNT"
	INFO_TOKEN_VALIDATED_SUCCESSFULLY = "INFO_TOKEN_VALIDATED_SUCCESSFULLY"
)
