package response

const (
	ErrCodeSucces       = 20001 // success
	ErrCodeParamInvalid = 20003 //email is invalid
	ErrInvalidToken     = 30001 // token invalid
	ErrInvalidOTP       = 30002 // OTP invalid
	ErrSendEmailOtp     = 30003 // sent email
	// user authentication
	ErrUnauthorized   = 40001 // user authentication unauthorized
	ErrCodeAuthFailed = 40005
	// Register Code
	ErrCodeUserHasExists = 50001 // user already registered
	// Err Login
	ErrCodeOtpNotExists     = 60009
	ErrCodeUserOtpNotExists = 60010
	//  Two Factor Authentication
	ErrCodeTwoFactorAuthSetupFailed  = 80001
	ErrCodeTwoFactorAuthVerifyFailed = 80002
	// error code Menu
	ErrCodeMenuNotFound  = 90001 // menu not found
	ErrCodeMenuHasExists = 90002 // menu already exists
	ErrCodeMenuErrror    = 90003
	// error code Role
	ErrCodeRoleError    = 10001
	ErrCodeRoleSucces   = 10002
	ErrCodeRoleNotFound = 10003
	// error code RoleMenu
	ErrCodeRoleMenuError    = 11001
	ErrCodeRoleMenuSucces   = 11002
	ErrCodeRoleMenuNotFound = 11003
	// error code RoleAccount
	ErrCodeRoleAccountError    = 12001
	ErrCodeRoleAccountSucces   = 12002
	ErrCodeRoleAccountNotFound = 12003
	ErrCodeRoleAccountValid    = 12004 //
	// error code account
	ErrCodeRoleAccountMaxNumber = 13002 // role account max number
	// error code license
	ErrCodeLicenseValid = 13000
)

// message
var msg = map[int]string{
	ErrCodeSucces:                    "success",
	ErrCodeParamInvalid:              "Email is invalid",
	ErrInvalidToken:                  "Token invalid",
	ErrCodeUserHasExists:             "User already registered",
	ErrInvalidOTP:                    "OTP invalid",
	ErrSendEmailOtp:                  "Failed to send email OTP",
	ErrCodeUserOtpNotExists:          "User OTP not exists",
	ErrCodeAuthFailed:                "Authentication failed",
	ErrCodeTwoFactorAuthSetupFailed:  "Failed to setup Two Factor Authentication",
	ErrUnauthorized:                  "Unauthorized",
	ErrCodeTwoFactorAuthVerifyFailed: "Two Factor Authentication setup failed",
	ErrCodeMenuNotFound:              "menu not found",
	ErrCodeMenuHasExists:             "menu already exists",
	ErrCodeMenuErrror:                "menu error",
	ErrCodeRoleMenuError:             "role menu error",
	ErrCodeRoleMenuSucces:            "role menu succes",
	ErrCodeRoleMenuNotFound:          "role menu not found",
	ErrCodeRoleError:                 "role error",
	ErrCodeRoleSucces:                "role succes",
	ErrCodeRoleNotFound:              "role not found",
	ErrCodeRoleAccountError:          "role account error",
	ErrCodeRoleAccountSucces:         "role account succes",
	ErrCodeRoleAccountNotFound:       "role account not found",
	ErrCodeRoleAccountValid:          "role account valid",
	ErrCodeRoleAccountMaxNumber:      "role account max number",
	ErrCodeLicenseValid:              "License valid",
}
